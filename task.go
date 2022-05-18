package main

import (
	"reflect"
	"unsafe"
)

type Task uint32

type Tasks []Task

func TasksFromBytes(v []byte) (tasks Tasks) {
	if len(v)%4 != 0 {
		panic("无效数据")
	}
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&v))
	th := (*reflect.SliceHeader)(unsafe.Pointer(&tasks))
	th.Data = bh.Data
	th.Len = bh.Len / 4
	th.Cap = bh.Cap / 4
	return
}

func (t *Tasks) ToBytes() (v []byte) {
	th := (*reflect.SliceHeader)(unsafe.Pointer(t))
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&v))
	bh.Data = th.Data
	bh.Len = th.Len * 4
	bh.Cap = th.Cap * 4
	return
}

//func (t *Task) Room() *Room {
//	var info bili.RoomInfo
//
//	return
//}
//
//func getT
