package main

import (
	"os"
)

type channel struct {
	Uid    uint32 `json:"bilibili_uid"`
	RoomID uint32 `json:"bilibili_live_room"`
	Name   string `json:"name"`
	Face   string `json:"face"`
	IsLive bool   `json:"is_live"`
	//LastDanmu  int    `json:"last_danmu"`
	//TotalDanmu int    `json:"total_danmu"`
	//LastLive   int64  `json:"last_live"`
}

func main() {
	data, err := os.ReadFile("tasks.db")
	if err != nil {
		panic(err)
	}
	tasks := TasksFromBytes(data)
	for i := range tasks {
		v := &User{id: uint32(tasks[i])}
		v.Start(i)
	}
}
