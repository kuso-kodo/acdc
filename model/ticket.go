package model

import "time"

type Ticket struct {
	TicketID          uint      `gorm:"primary_key" json:"ticket_id"`
	Room              Room      `json:"-"`
	StartAt           time.Time `json:"start_at"`
	EndAt             time.Time `json:"end_at"`
	ServiceCount      uint      `json:"service_count"`
	TargetTemperature float32   `json:"target_temperature"`
	FanSpeed          uint      `json:"fan_speed"`
	TotalFee          float32   `json:"total_fee"`
	RoomRefer         uint      `json:"-"`
}
