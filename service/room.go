package service

import (
	"github.com/name1e5s/acdc/db"
	"github.com/name1e5s/acdc/model"
)

func GetRoomByName(roomName string) (model.Room, bool) {
	room := &model.Room{RoomID: 0}
	db.GetDataBase().Where("room_name = ?", roomName).First(&room)
	return *room, room.RoomID != 0
}
