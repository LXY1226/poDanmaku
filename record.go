package main

import (
	"bytes"
	"compress/zlib"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/andybalholm/brotli"
	"io"
	"log"
	"net"
	"os"
	"poDanmaku/bili"
	"runtime/debug"
	"strconv"
	"time"
)

type RecordSession struct {
	conn    *net.TCPConn
	rdBuf   []byte
	rawBuf  []byte
	logTag  string
	curDate *string
	log.Logger
	User
	storage DanmakuFile
}

const (
	ProtocolPlain = iota
	ProtocolHeartbeat
	ProtocolDeflate
	ProtocolBrotli
)

// NextPacket data在下次调用Read前有效
func (r *RecordSession) NextPacket() (proto uint16, act uint32, data []byte, err error) {
reParse:
	data, err = r.read(4)
	if err != nil {
		return
	}
	packLen := binary.BigEndian.Uint32(data)
	if packLen > 1<<20 {
		err = errors.New("数据包过大:" + strconv.Itoa(int(packLen)))
		return
	}
	data, err = r.read(packLen - 4)
	if err != nil {
		return
	}
	if binary.BigEndian.Uint16(data) != 16 {
		err = errors.New("无效包头长度:" + strconv.Itoa(int(binary.BigEndian.Uint16(data))))
		return
	}
	proto = binary.BigEndian.Uint16(data[2:])
	switch proto {
	case ProtocolDeflate:
		var zr io.ReadCloser
		zr, err = zlib.NewReader(bytes.NewReader(data[12:]))
		r.rdBuf, err = io.ReadAll(zr)
		zr.Close()
		if err != nil {
			return 0, 0, nil, err
		}
		goto reParse
	case ProtocolBrotli:
		zr := brotli.NewReader(bytes.NewReader(data[12:]))
		r.rdBuf, err = io.ReadAll(zr)
		if err != nil {
			return 0, 0, nil, err
		}
		goto reParse
	}
	act = binary.BigEndian.Uint32(data[4:])
	data = data[12:]
	return
}

var dataHeartBeat = []byte{0x00, 0x00, 0x00, 0x10,
	0x00, 0x10,
	0x00, 0x01,
	0x00, 0x00, 0x00, 0x02,
	0x00, 0x00, 0x00, 0x01}

func (r *RecordSession) Heartbeat() {
	time.AfterFunc(45*time.Second, r.Heartbeat)
	if r.conn != nil {
		_, err := r.conn.Write(dataHeartBeat)
		if err != nil {
			r.conn.Close()
			return
		}
	}
}

func NewRecord(user *User) *RecordSession {
	return &RecordSession{User: *user, storage: *NewDanmaku(strconv.Itoa(int(user.UserInfo.UID)))}
}

var errLog *os.File

func (r *RecordSession) GetHandshake() {

}

func (r *RecordSession) UpdateUserInfo() {

}

func (r *RecordSession) Start(i int) {
	r.Logger = *log.New(os.Stdout, fmt.Sprintf("#%05d %d %s ", i, r.RoomInfo.RoomId, r.UserInfo.Uname), log.Ldate|log.Ltime|log.Lmsgprefix)
	defer func() {
		if er := recover(); er != nil {
			r.SetOutput(io.MultiWriter(os.Stderr, errLog))
			r.Printf("发生致命错误 %v", er)
			debug.PrintStack()
		}
	}()
	roomID := strconv.Itoa(int(r.RoomInfo.RoomId))
	var sleep time.Duration
	for {
		if sleep != 0 {
			time.Sleep(sleep)
		} else {
			sleep = 1 * time.Second
		}
		info := bili.GetDanmuInfo(r.RoomInfo.RoomId)

		buf := getBuf(16) // 将会返回一个len=16, cap=512的buf

		buf[5] = 0x10
		buf[7] = 0x01
		buf[11] = 0x07
		buf = append(buf, `{"uid":0,"protover":3,"platform":"web","clientver":"2.6.25","type":2,"key":`...)
		buf = append(buf, info.Token...)
		buf = append(buf, `,"roomid":`...)
		buf = append(buf, roomID...)
		buf = append(buf, '}')

		binary.BigEndian.PutUint32(buf, uint32(len(buf)))

		for _, host := range info.HostList {
			if r.conn != nil {
				r.conn.Close()
				r.conn = nil
			}
			addrs, err := net.DefaultResolver.LookupIPAddr(context.Background(), host.Host)
			if err != nil {
				r.Println(err)
				continue
			}
			for _, addr := range addrs {
				r.conn, err = net.DialTCP("tcp", nil, &net.TCPAddr{IP: addr.IP, Port: 2243})
				if err != nil {
					r.Println("连接弹幕姬失败", err)
					continue
				}
				r.conn.Write(buf)
				putBuf(buf) // release here
				proto, _, _, err := r.NextPacket()
				if err != nil {
					r.Println("与弹幕姬握手失败", err)
					break
				}
				const (
					ActHeartBeat     = 2
					ActHeartBeatResp = 3
					ActMessage       = 5
					ActUserAuth      = 7
					ActConnected     = 8
				)
				if proto != ProtocolHeartbeat {
					r.Println("回应首包：期待心跳包，但是得到", proto)
					continue
				}
				r.Heartbeat()
				r.Println("连接成功")
				sleep = 0
				for {
					_, act, data, err := r.NextPacket()
					if err != nil {
						r.Println("失去连接", err)
						break
					}
					if act == ActHeartBeatResp {
						continue
					}
					if act == ActMessage {
						r.storage.WriteDanmaku(data)
						continue
					}
					r.Println(act, data)
				}
			}
		}
		r.Println("无可用弹幕姬")
	}
}
