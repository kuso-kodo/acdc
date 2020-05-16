package model

import "time"

const (
	ReasonShutDown = iota
	ReasonFanSpeedChanged
	ReasonPriorityChanged
)

type Ticket struct {
	TicketID     uint      `gorm:"primary_key" json:"ticket_id"`
	Room         Room      `json:"-"`
	User         User      `json:"-"`
	StartAt      time.Time `json:"start_at"`
	EndAt        time.Time `json:"end_at"`
	ServiceCount uint      `json:"service_count"`
	FanSpeed     uint      `json:"fan_speed"`
	TotalFee     float32   `json:"total_fee"`
	RoomRefer    uint      `json:"-"`
	UserRefer    uint      `json:"-"`
	Paid         bool      `json:"paid"`
	Reason       int       `json:"-"`
}
