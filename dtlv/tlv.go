package dtlv

import (
	"encoding/binary"
	"math"
	"reflect"
	"unsafe"
)

type Tag byte
type Type byte
type Length uint16

const (
	TypeI8 Type = iota
	TypeI16
	TypeI32
	TypeI64
	TypeFloat
	TypeDouble
	TypeString1 // i8(len)+data
	TypeString4 // i32(len)+data
	TypeMap
	TypeList
	TypeStructStart
	TypeStructEnd
	TypeZero
	TypeBytes
	TypeVarInt // custom
)

type Head struct {
	Type // 前四位为tag，若Tag == 15，Tag为后面的byte
	Tag
	data dat
}

type dat struct {
	maybe []byte // for string
	data  []byte
}

func (h *Head) AsBArr() *[2]byte {
	if h.Tag < 15 {
		h.Type |= h.Tag << 4
	} else {
		h.Type |= 0xF0
	}
	return (*[2]byte)(unsafe.Pointer(h))
}

type Encoder []Head

func (enc *Encoder) AppendI8(tag Tag, v uint8) {
	if v == 0 {
		enc.AppendZero(tag)
		return
	}
	var b [1]byte
	bs := b[:]
	b[0] = v
	*enc = append(*enc, Head{TypeI8, tag, dat{nil, bs}})
}

func (enc *Encoder) AppendI16(tag Tag, v uint16) {
	if v == 0 {
		enc.AppendZero(tag)
		return
	}
	var b [2]byte
	bs := b[:]
	binary.LittleEndian.PutUint16(bs, v)
	*enc = append(*enc, Head{TypeI16, tag, dat{nil, bs}})
}

func (enc *Encoder) AppendI32(tag Tag, v uint32) {
	if v == 0 {
		enc.AppendZero(tag)
		return
	}
	var b [4]byte
	bs := b[:]
	binary.LittleEndian.PutUint32(bs, v)
	*enc = append(*enc, Head{TypeI32, tag, dat{nil, bs}})
}

func (enc *Encoder) AppendI64(tag Tag, v uint64) {
	if v == 0 {
		enc.AppendZero(tag)
		return
	}
	var b [8]byte
	bs := b[:]
	binary.LittleEndian.PutUint64(bs, v)
	*enc = append(*enc, Head{TypeI64, tag, dat{nil, bs}})
}

func (enc *Encoder) AppendUVarInt(tag Tag, v uint32) {
	if v == 0 {
		enc.AppendZero(tag)
		return
	}
	var b [binary.MaxVarintLen32]byte
	i := 0
	for v >= 0x80 {
		b[i] = byte(v) | 0x80
		v >>= 7
		i++
	}
	b[i] = byte(v)
	*enc = append(*enc, Head{TypeVarInt, tag, dat{nil, b[:i+1]}})
}

func (enc *Encoder) AppendFloat32(tag Tag, v float32) {
	enc.AppendI32(tag, math.Float32bits(v))
}

func (enc *Encoder) AppendFloat64(tag Tag, v float64) {
	enc.AppendI64(tag, math.Float64bits(v))
}

func (enc *Encoder) AppendString(tag Tag, v string) {
	if len(v) <= 0xFF {
		var b [1]byte
		b[0] = byte(len(v))
		*enc = append(*enc, Head{TypeString1, tag, dat{b[:], S2B(v)}})
		return
	}
	var b [4]byte
	bs := b[:]
	binary.LittleEndian.PutUint32(bs, uint32(len(v))) // FIXME int to uint32 overflow on 64b machine
	*enc = append(*enc, Head{TypeString4, tag, dat{bs, S2B(v)}})
}

func (enc *Encoder) AppendZero(tag Tag) {
	*enc = append(*enc, Head{TypeZero, tag, dat{}})
}

func (enc *Encoder) EncodeAppend(buf []byte) []byte {
	l := 0
	for _, v := range *enc {
		l++ // Type
		if v.Tag >= 15 {
			l++ // Tag
		}
		l += len(v.data.maybe)
		l += len(v.data.data)
	}
	buf = GrowSlice(buf, l)
	for _, v := range *enc {
		if v.Tag < 15 {
			buf = append(buf, v.Type|v.Tag<<4)

		} else {
			buf = append(buf, v.Type|0xF0)
			buf = append(buf, byte(v.Tag))
		}
		buf = append(buf, v.data.maybe...)
		buf = append(buf, v.data.data...)
	}
	return buf
}

func GrowSlice[T any](v []T, l int) []T {
	if c := cap(v) - len(v); c < l {
		if c < 2*len(v) {
			c = 2 * len(v)
		}
		buf := make([]T, len(v), c)
		copy(buf, v)
		v = buf
	}
	return v
}

func S2B(s string) (b []byte) {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	bh.Data = sh.Data
	bh.Cap = sh.Len
	bh.Len = sh.Len
	return
}
