package rgzip

import (
	"bytes"
	"compress/gzip"
	"io"
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestGzip(t *testing.T) {
	const size = 1024 * 1024
	data := make([]byte, size)
	rand.Read(data)
	buf := new(bytes.Buffer)
	writer := NewWriter(buf)
	n, err := writer.Write(data)
	if err != nil || n != size {
		t.Fatal(err)
	}
	err = writer.Close()
	if err != nil {
		t.Fatal(err)
	}
	decompressed := make([]byte, size)
	decompress(buf.Bytes(), decompressed)
	if !bytes.Equal(data, decompressed) {
		t.Fatal("fail")
	}
}

func TestTempCloseEqualOutput(t *testing.T) {
	const size = 1 << 20
	data := make([]byte, size)
	rand.Read(data)
	buf1 := new(bytes.Buffer)
	writer := NewWriter(buf1)
	_, err := writer.Write(data)
	if err != nil {
		t.Fatal(err)
	}
	remain := writer.TempClose()
	bufResumed := make([]byte, buf1.Len()+len(remain))
	copy(bufResumed, buf1.Bytes())
	copy(bufResumed[buf1.Len():], remain)
	writer.Close()
	if !bytes.Equal(buf1.Bytes(), bufResumed) {
		t.Fatal("output not equal")
	}
	decompressed := make([]byte, size)
	decompress(buf1.Bytes(), decompressed)
	if !bytes.Equal(data, decompressed) {
		t.Fatal("Close() decompressed not equal")
	}
	decompress(bufResumed, decompressed)
	if !bytes.Equal(data, decompressed) {
		t.Fatal("TempClose() decompressed not equal")
	}

}

func TestResume(t *testing.T) {
	const size = 2 << 20
	data := make([]byte, size)
	rand.Read(data)
	buf1 := new(bytes.Buffer)
	writer := NewWriter(buf1)
	_, err := writer.Write(data[:size/2])
	if err != nil {
		t.Fatal(err)
	}
	remain := writer.TempClose()
	bufTemp := make([]byte, buf1.Len()+len(remain))
	copy(bufTemp, buf1.Bytes())
	copy(bufTemp[buf1.Len():], remain)
	output := make([]byte, size)
	decompress(bufTemp, output[:size/2])
	if !bytes.Equal(data[:size/2], output[:size/2]) {
		t.Fatal("fail1")
	}

	_, err = writer.Write(data[size/2:])
	if err != nil {
		panic(err)
	}
	err = writer.Close()
	if err != nil {
		panic(err)
	}
	decompress(buf1.Bytes(), output)
	if !bytes.Equal(data, output) {
		t.Fatal("fail2")
	}
}

func decompress(buf, dst []byte) {
	reader, err := gzip.NewReader(bytes.NewReader(buf))
	if err != nil {
		panic(err)
	}
	_, err = io.ReadFull(reader, dst)
	if err != nil {
		panic(err)
	}
	reader.Close()
}
