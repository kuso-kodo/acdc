package model

import "time"

type Room struct {
	RoomID             uint      `gorm:"primary_key" json:"room_id"`
	RoomName           string    `gorm:"unique" json:"room_name"`
	IsPowerOn          bool      `json:"power_on"`
	IsServicing        bool      `json:"servicing"`
	CurrentTemperature float32   `json:"current_temperature"`
	TargetTemperature  float32   `json:"target_temperature"`
	FanSpeed           uint      `json:"fan_speed"`
	LastOnTime         time.Time `json:"last_on"`
	Tickets            []Ticket  `gorm:"foreignkey:RoomRefer;association_foreignkey:RoomID"`
}
