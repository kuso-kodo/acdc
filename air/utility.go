package air

import (
	"github.com/name1e5s/acdc/db"
	"github.com/name1e5s/acdc/model"
	"github.com/name1e5s/acdc/service"
	"time"
)

func calculateFee(task *Task) float32 {
	var base float32
	switch task.FanSpeed {
	case 0:
		base = GetAir().config.LowFanSpeedFeeRate
	case 1:
		base = GetAir().config.MediumFanSpeedFeeRate
	case 2:
		base = GetAir().config.HighFanSpeedFeeRate
	}
	var factor float32
	switch task.DefaultPriority {
	case LowPriority:
		factor = GetAir().config.LowPriorityFactor
	case MediumPriority:
		factor = GetAir().config.MediumFanSpeedFeeRate
	case HighPriority:
		factor = GetAir().config.HighFanSpeedFeeRate
	}
	return float32(task.CurrentServiceCount) * base * factor
}

func generateTicket(task *Task, reason int) {
	var shutdown uint = 0
	var priorityChanged uint = 0
	var fanSpeedChanged uint = 0
	if reason == model.ReasonShutDown {
		shutdown = 1
	} else if reason == model.ReasonFanSpeedChanged {
		fanSpeedChanged = 1
	} else if reason == model.ReasonPriorityChanged {
		priorityChanged = 1
	}

	userID, _ := service.GetCheckInCheckOutMap().FindUserByRoom(task.RoomID)
	ticket := model.Ticket{
		StartAt:         task.LastModifiedTime,
		EndAt:           time.Now(),
		ServiceCount:    task.CurrentServiceCount,
		FanSpeed:        task.FanSpeed,
		TotalFee:        calculateFee(task),
		RoomRefer:       task.RoomID,
		UserRefer:       userID,
		Paid:            false,
		Shutdown:        shutdown,
		PriorityChanged: priorityChanged,
		FanSpeedChanged: fanSpeedChanged,
	}
	db.GetDataBase().Create(&ticket)
}

func GenerateShutdownTicket(task *Task) {
	generateTicket(task, model.ReasonShutDown)
}

func GenerateFanSpeedChangedTicket(task *Task) {
	generateTicket(task, model.ReasonFanSpeedChanged)
}

func GeneratePriorityChangedTicket(task *Task) {
	generateTicket(task, model.ReasonPriorityChanged)
}
