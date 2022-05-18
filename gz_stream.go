package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"poDanmaku/rgzip"
	"strconv"
	"unsafe"
)

type gzStreamFile struct {
	// sync.RWMutex
	f        *os.File
	lastHead int64
	*rgzip.Writer
}

func NewGzStreamFile(name string) (*gzStreamFile, error) {
	f, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	n, _ := f.Seek(0, io.SeekEnd)
	return &gzStreamFile{
		f:        f,
		lastHead: n,
		Writer:   rgzip.NewWriter(f),
	}, nil
}

func (g *gzStreamFile) Close() {
	g.Writer.Close()
	g.f.Close()
}

func (g *gzStreamFile) LastHead() (offset int64) {
	return g.lastHead
}

func (g *gzStreamFile) ReOpen() (offset int64) {
	g.Writer.Close()
	g.lastHead, _ = g.f.Seek(0, io.SeekEnd)
	g.Writer.Reset(g.f)
	return g.lastHead
}

// WriteChunk 用于响应整个gz数据中的一段
func (g *gzStreamFile) WriteChunk(eTag string, wr http.ResponseWriter, start, end int64) {
	if end < 8 {
		http.Error(wr, "bad gz range", http.StatusInternalServerError)
		return
	}
	var crc [4]byte
	n, err := g.f.ReadAt(crc[:], end-8)
	if err != nil {
		http.Error(wr, fmt.Sprint("read crc fail", err), http.StatusGone)
		return
	}
	if n != 4 {
		http.Error(wr, "short read", http.StatusGone)
		return
	}
	eT := encodeETag(*(*uint32)(unsafe.Pointer(&crc)))
	if eTag == eT {
		wr.WriteHeader(http.StatusNotModified)
		return
	}

	contentLength := end - start
	wr.Header().Set("Content-Length", strconv.FormatInt(contentLength, 10))
	wr.Header().Set("Content-Encoding", "gzip")
	wr.Header().Set("ETag", eT)
	err = g.sendRange(start, end, wr)
	if err != nil {
		log.Print(err)
	}
}

func (g *gzStreamFile) sendRange(start, end int64, wr http.ResponseWriter) error {
	fGz, err := os.Open(g.f.Name())
	if err != nil {
		http.Error(wr, fmt.Sprint("fail to open file", err), http.StatusServiceUnavailable)
		return err
	}
	defer fGz.Close()
	fGz.Seek(start, io.SeekStart)

	_, err = io.Copy(wr, io.LimitReader(fGz, end-start))
	return err
}

// WriteLatest 但是先得确定这个wr对应的
func (g *gzStreamFile) WriteLatest(eTag string, wr http.ResponseWriter) {
	h := g.Writer.Hash()
	eT := encodeETag(h)
	if eTag == eT {
		wr.WriteHeader(http.StatusNotModified)
		return
	}
	// 如果在以上期间写入了新数据，则hash不等，也无所谓...
	offsetGzMid, _ := g.f.Seek(0, io.SeekCurrent)
	gzTail := g.Writer.Tail()

	contentLength := (offsetGzMid - g.lastHead) + int64(len(gzTail))
	wr.Header().Set("Content-Length", strconv.FormatInt(contentLength, 10))
	wr.Header().Set("Content-Encoding", "gzip")
	wr.Header().Set("ETag", eT)
	err := g.sendRange(g.lastHead, offsetGzMid, wr)
	if err != nil {
		return
	}
	_, err = wr.Write(gzTail)
	if err != nil {
		log.Print(err)
	}
}

func encodeETag(crc uint32) string {
	eT := make([]byte, 2+base64.RawURLEncoding.EncodedLen(4))
	eT[0] = '"'
	eT[len(eT)-1] = '"'
	base64.RawURLEncoding.Encode(eT[1:], (*(*[4]byte)(unsafe.Pointer(&crc)))[:])
	return string(eT)
}
