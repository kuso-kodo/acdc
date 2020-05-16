package service

import (
	"errors"
	"github.com/name1e5s/acdc/utility"
	"sync"
)

type CheckInCheckOutMap struct {
	value *utility.Map
}

var (
	checkInCheckOutMap *CheckInCheckOutMap
	once               sync.Once
)

func GetCheckInCheckOutMap() *CheckInCheckOutMap {
	once.Do(func() {
		checkInCheckOutMap = &CheckInCheckOutMap{value: utility.NewMap()}
	})
	return checkInCheckOutMap
}

func (m *CheckInCheckOutMap) CheckIn(userID, roomID uint) error {
	_, ok := m.value.Load(roomID)
	if ok {
		return errors.New("room has already check in")
	}
	m.value.Store(roomID, userID)
	return nil
}

func (m *CheckInCheckOutMap) CheckOut(userID uint) {
	key, ok := m.value.FindKeyByValue(userID)
	if ok {
		m.value.Delete(key)
	}
}

func (m *CheckInCheckOutMap) FindRoomByUser(userID uint) (uint, bool) {
	return m.value.FindKeyByValue(userID)
}

func (m *CheckInCheckOutMap) FindUserByRoom(roomID uint) (uint, bool) {
	return m.value.Load(roomID)
}
