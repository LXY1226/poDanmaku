package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
	"unsafe"
)

const dataDir = "data/"

var isStopping bool

func init() {
	os.MkdirAll(dataDir, 0644)
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-ch
		log.Println("Stopping... Press once more to force exit")
		isStopping = true
		go func() {
			for _, f := range activeFile {
				f.Close()
			}
			os.Exit(1)
		}()
		<-ch
		log.Println("forcing exit")
		os.Exit(-1)
	}()
}

var date *string

func init() {
	updateDate()
}

var activeFile []*DanmakuFile

func updateDate() {
	dat := time.Now().Format("06-01-02")
	date = &dat
	for _, f := range activeFile {
		f.Close()
	}
	d := time.Duration(time.Now().UnixNano()) + (8 * time.Hour) // CST 00:00
	d %= 24 * time.Hour
	time.AfterFunc(24*time.Hour-d, updateDate)
}

type DanmakuFile struct {
	room   *User
	prefix string
	tags   *os.File
	data   *gzStreamFile
	count  int64
}

type TagType uint64

type Tag struct {
	time  uint64
	stat  uint64
	extra uint64

	coverLen uint32
	titleLen uint32

	cover, title string
}

func (t *Tag) FillWithRoom(room *User) {

}

const (
	TagMask TagType = 0xF << (64 - 4)
	TagTime TagType = iota
	TagNum
	TagStatus

	TagCustom = 0xF
)

func (f *DanmakuFile) TagByTime(time uint64) {
	f.Tag(TagTime, time)
}

type StatType uint64

const (
	StatLive StatType = iota
	StatOffline
)

func (f *DanmakuFile) TagByStatus(status uint64) {
	f.Tag(TagStatus, status)
}

func (f *DanmakuFile) Tag(typ TagType, tag uint64) {
	if tag&uint64(TagMask) != 0 {
		panic("tag's first 4bit should be zero")
	}
	tag |= uint64(typ) << (64 - 4)
	offset := f.data.ReOpen()
	f.tags.Write((*(*[8]byte)(unsafe.Pointer(&tag)))[:])
	f.tags.Write((*(*[8]byte)(unsafe.Pointer(&offset)))[:])
}

func (f *DanmakuFile) WriteDanmaku(data []byte) {
	if f.data == nil {
		if isStopping {
			return
		}
		f.NextFile()
	}
	f.data.Write(data)
}

func (f *DanmakuFile) HttpHandler(w http.ResponseWriter, req *http.Request) {
	if !strings.Contains(req.Header.Get("Accept-Encoding"), "gzip") {
		http.Error(w, "why not support gzip?", http.StatusNotAcceptable)
		return
	}

	// return latest
	f.data.WriteLatest(req.Header.Get("If-None-Match"), w)
}

// Close 关闭文件并刷新缓冲区，下次Write将会重新打开文件
func (f *DanmakuFile) Close() {
	f.data.Close()
	f.data = nil
}

func NewDanmaku(prefix string) (f *DanmakuFile) {
	os.MkdirAll(dataDir+prefix, 0644)
	f = &DanmakuFile{prefix: prefix + "/"}
	activeFile = append(activeFile, f)
	return f
}

func (f *DanmakuFile) NextFile() {
	if f.data != nil {
		f.data.Close()
	}
	if f.tags != nil {
		f.tags.Close()
	}
	tag, err := os.OpenFile(dataDir+f.prefix+".tag", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	f.tags = tag
	f.data, err = NewGzStreamFile(dataDir + f.prefix + *date + ".json.gz")
	if err != nil {
		panic(err)
	}
}

// GzIndex raw uint32 array
type GzIndex struct {
	f *os.File
}
