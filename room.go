package main

import "poDanmaku/bili"

// User 某用户的当前的信息
// 对应UserInfo，首页展示列表
type User struct {
	id       uint32
	RoomInfo *bili.RoomInfo
	UserInfo *bili.UserInfo
}

func (r *User) Start(index int) {
	r.updateUserInfo()
	r.updateRoomInfo()
	go NewRecord(r).Start(index)
}

func (r *User) updateUserInfo() {
	r.UserInfo = bili.GetUserInfo(r.id)
}

func (r *User) updateRoomInfo() {
	r.RoomInfo = bili.GetRoomInfo(r.id)

}

//func (t *Task)

func (r *User) Refresh() {

}
