package main

import (
	"fmt"
	"hash/crc32"
	"os"
	"runtime"
	"testing"
)

func TestBytesGC(t *testing.T) {
	tasks := getTasks()
	runtime.GC()
	fmt.Printf("0x%x", crc32.Checksum(tasks.ToBytes(), crc32.IEEETable))
}

func getTasks() Tasks {
	data, err := os.ReadFile("tasks.db")
	if err != nil {
		panic(err)
	}
	return TasksFromBytes(data)
}
