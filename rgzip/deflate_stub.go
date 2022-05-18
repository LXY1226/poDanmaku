package rgzip

import (
	"bytes"
	"io"
	"math"
	"reflect"
	"sort"
	"unsafe"
)

const (
	logWindowSize = 15
	windowSize    = 1 << logWindowSize
	windowMask    = windowSize - 1

	// The LZ77 step produces a sequence of literal tokens and <length, offset>
	// pair tokens. The offset is also known as distance. The underlying wire
	// format limits the range of lengths and offsets. For example, there are
	// 256 legitimate lengths: those in the range [3, 258]. This package's
	// compressor uses a higher minimum match length, enabling optimizations
	// such as finding matches via 32-bit loads and compares.
	baseMatchLength = 3       // The smallest match length per the RFC section 3.2.5
	minMatchLength  = 4       // The smallest match length that the compressor actually emits
	maxMatchLength  = 258     // The largest match length
	baseMatchOffset = 1       // The smallest match offset
	maxMatchOffset  = 1 << 15 // The largest match offset

	// The maximum number of tokens we put into a single flate block, just to
	// stop things from getting too large.
	maxFlateBlockTokens = 1 << 14
	maxStoreBlockSize   = 65535
	hashBits            = 17 // After 17 performance degrades
	hashSize            = 1 << hashBits
	hashMask            = (1 << hashBits) - 1
	maxHashOffset       = 1 << 24

	skipNever = math.MaxInt32
)

const (
	// The largest offset code.
	offsetCodeCount = 30

	// The special code used to mark the end of a block.
	endBlockMarker = 256

	// The first length code.
	lengthCodesStart = 257

	// The number of codegen codes.
	codegenCodeCount = 19
	badCode          = 255

	// bufferFlushSize indicates the buffer size
	// after which bytes are flushed to the writer.
	// Should preferably be a multiple of 6, since
	// we accumulate 6 bytes between writes to the buffer.
	bufferFlushSize = 240

	// bufferSize is the actual output byte buffer size.
	// It must have additional headroom for a flush
	// which can contain up to 8 bytes.
	bufferSize = bufferFlushSize + 8
)

// just an empty copy of deflate.go

type compressionLevel struct {
	level, good, lazy, nice, chain, fastSkipHashing int
}

// hcode is a huffman code with a bit code and bit length.
type hcode struct {
	code, len uint16
}

type literalNode struct {
	literal uint16
	freq    int32
}

type byLiteral []literalNode

func (s *byLiteral) sort(a []literalNode) {
	*s = byLiteral(a)
	sort.Sort(s)
}

func (s byLiteral) Len() int { return len(s) }

func (s byLiteral) Less(i, j int) bool {
	return s[i].literal < s[j].literal
}

func (s byLiteral) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

type byFreq []literalNode

func (s *byFreq) sort(a []literalNode) {
	*s = byFreq(a)
	sort.Sort(s)
}

func (s byFreq) Len() int { return len(s) }

func (s byFreq) Less(i, j int) bool {
	if s[i].freq == s[j].freq {
		return s[i].literal < s[j].literal
	}
	return s[i].freq < s[j].freq
}

func (s byFreq) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

type huffmanEncoder struct {
	codes     []hcode
	freqcache []literalNode
	bitCount  [17]int32
	lns       byLiteral // stored to avoid repeated allocation in generate
	lfs       byFreq    // stored to avoid repeated allocation in generate
}

func (h *huffmanEncoder) clone() *huffmanEncoder {
	h2 := cloneT(h)
	h2.codes = cloneArrT(h.codes)
	h2.freqcache = nil
	h2.lns = cloneArrT(h.lns)
	h2.lfs = cloneArrT(h.lfs)
	return h2
}

type huffmanBitWriter struct {
	// writer is the underlying writer.
	// Do not use it directly; use the write method, which ensures
	// that Write errors are sticky.
	writer io.Writer

	// Data waiting to be written is bytes[0:nbytes]
	// and then the low nbits of bits.  Data is always written
	// sequentially into the bytes array.
	bits            uint64
	nbits           uint
	bytes           [bufferSize]byte
	codegenFreq     [codegenCodeCount]int32
	nbytes          int
	literalFreq     []int32
	offsetFreq      []int32
	codegen         []uint8
	literalEncoding *huffmanEncoder
	offsetEncoding  *huffmanEncoder
	codegenEncoding *huffmanEncoder
	err             error
}

func (w *huffmanBitWriter) clone() *huffmanBitWriter {
	if w == nil {
		return nil
	}
	w2 := cloneT(w)
	w2.codegen = cloneArrT(w.codegen)
	w2.literalFreq = cloneArrT(w.literalFreq)
	w2.offsetFreq = cloneArrT(w.offsetFreq)
	w2.literalEncoding = w.literalEncoding.clone()
	w2.offsetEncoding = w.offsetEncoding.clone()
	w2.codegenEncoding = w.codegenEncoding.clone()
	return w2
}

func (w *huffmanBitWriter) setOutput(b *bytes.Buffer) {
	w.writer = b
}

type token uint32

type compressor struct {
	compressionLevel

	w          *huffmanBitWriter
	bulkHasher func([]byte, []uint32)

	// compression algorithm
	fill      func(*compressor, []byte) int // copy data to window
	step      func(*compressor)             // process window
	sync      bool                          // requesting flush
	bestSpeed uintptr                       // Encoder for BestSpeed, never used pointer *deflateFast

	// Input hash chains
	// hashHead[hashValue] contains the largest inputIndex with the specified hash value
	// If hashHead[hashValue] is within the current window, then
	// hashPrev[hashHead[hashValue] & windowMask] contains the previous index
	// with the same hash value.
	chainHead  int
	hashHead   [hashSize]uint32
	hashPrev   [windowSize]uint32
	hashOffset int

	// input window: unprocessed data is window[index:windowEnd]
	index         int
	window        []byte
	windowEnd     int
	blockStart    int  // window index where current tokens start
	byteAvailable bool // if true, still need to process window[index-1].

	// queued output tokens
	tokens []token

	// deflate state
	length         int
	offset         int
	hash           uint32
	maxInsertIndex int
	err            error

	// hashMatch must be able to contain hashes for the maximum match length.
	hashMatch [maxMatchLength - 1]uint32
}

func (c *compressor) clone(c2 *compressor) {
	copyT(c, c2)
	c.w = c2.w.clone()
	c.tokens = cloneArrT(c2.tokens)
	c.window = cloneArrT(c2.window)
}

func (c *compressor) setOutput(b *bytes.Buffer) {
	c.w.setOutput(b)
}

// A Writer takes data written to it and writes the compressed
// form of that data to an underlying writer (see NewWriter).
type flateWriter struct {
	d    compressor
	dict []byte
}

func (w *flateWriter) Write(data []byte) (n int, err error) {
	return flateWrite(w, data)
}

func (w *flateWriter) Close() error {
	return flateClose(w)
}

func (w *flateWriter) Reset(dst io.Writer) {
	flateReset(w, dst)
}

//go:linkname flateWrite compress/flate.(*Writer).Write
func flateWrite(w *flateWriter, data []byte) (n int, err error)

//go:linkname flateClose compress/flate.(*Writer).Close
func flateClose(w *flateWriter) error

//go:linkname flateReset compress/flate.(*Writer).Reset
func flateReset(w *flateWriter, dst io.Writer)

//go:linkname flateNewWriter compress/flate.NewWriter
func flateNewWriter(w io.Writer, level int) (*flateWriter, error)

func (w *flateWriter) setOutput(b *bytes.Buffer) {
	w.d.setOutput(b)
}

func (w *flateWriter) clone() *flateWriter {
	if w == nil {
		return nil
	}
	w2 := new(flateWriter)
	w2.d.clone(&w.d)
	// w.dict is nil
	return w2
}

func anyAsBytes(v unsafe.Pointer, size int) (b []byte) {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	bh.Data = uintptr(v)
	bh.Len = size
	bh.Cap = size
	return
}
func cloneT[T any](src *T) *T {
	dst := new(T)
	copyT(dst, src)
	return dst
}

func copyT[T any](dst, src *T) {
	size := unsafe.Sizeof(*src)
	b1 := anyAsBytes(unsafe.Pointer(src), int(size))
	b2 := anyAsBytes(unsafe.Pointer(dst), int(size))
	copy(b2, b1)
}

func cloneArrT[T any](src []T) []T {
	dst := make([]T, len(src))
	copy(dst, src)
	return dst
}
