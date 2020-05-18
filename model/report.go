package model

type Report struct {
	RoomName  string  `json:"room"`
	TotalTime string  `json:"total_time"`
	Service   uint    `json:"service"`
	Ticket    uint    `json:"ticket"`
	Fan       uint    `json:"fan"`
	Power     uint    `json:"power"`
	Priority  uint    `json:"priority"`
	Fee       float32 `json:"fee"`
}
