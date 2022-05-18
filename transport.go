package main

import (
	"io"
	"sync"
)

var bufPool = sync.Pool{}

func getBuf(size uint32) (buf []byte) {
	for i := 0; i < 5; i++ {
		bAny := bufPool.Get()
		if bAny == nil {
			break
		}
		buf = bAny.([]byte)
		if uint32(cap(buf)) >= size {
			buf = buf[:size]
			return
		}
		bufPool.Put(buf)
	}

	if size < 512 {
		buf = make([]byte, size, 512)
	} else {
		buf = make([]byte, size)
	}
	return
}

func putBuf(buf []byte) {
	if buf == nil {
		return
	}
	if cap(buf) > 4096 {
		return // let large buffer gc
	}
	bufPool.Put(buf)
}

func (r *RecordSession) read(size uint32) (buf []byte, err error) {
	if uint32(len(r.rdBuf)) >= size {
		buf = r.rdBuf[:size]
		r.rdBuf = r.rdBuf[size:]
		return
	}
	putBuf(r.rawBuf)
	buf = getBuf(size)
	r.rawBuf = buf
	_, err = io.ReadFull(r.conn, buf)
	return buf, err
}
