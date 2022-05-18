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
	//channel, err := os.Open("channel.json")
	//if err != nil {
	//	log.Panicln("读取频道列表失败", err)
	//}
	//errLog, err = os.OpenFile("error.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	//if err != nil {
	//	panic("创建错误日志失败")
	//}
	//var channels []User
	//decoder := json.NewDecoder(channel)
	//err = decoder.Decode(&channels)
	//if err != nil {
	//	log.Panicln("解析频道列表失败", err)
	//}
	//chs := make(Tasks, len(channels))
	//for i := range channels {
	//
	//	chs[i] = Task(channels[i])
	//	//binary.LittleEndian.PutUint32(chs[i*4:], uint32(channels[i].RoomID))
	//}
	//os.WriteFile("tasks.db", chs.ToBytes(), 0644)
	data, err := os.ReadFile("tasks.db")
	if err != nil {
		panic(err)
	}
	tasks := TasksFromBytes(data)
	for i := range tasks {
		v := &User{id: uint32(tasks[i])}
		v.Start(i)
	}

	//select {}
}
