package client

import (
	"github.com/gin-gonic/gin"
	"github.com/name1e5s/acdc/air"
	"github.com/name1e5s/acdc/db"
	"github.com/name1e5s/acdc/model"
	"github.com/name1e5s/acdc/schema"
	"github.com/name1e5s/acdc/service"
	"net/http"
	"time"
)

// @Summary Handle heartbeat.
// @Accept  json
// @Param clientRequest body schema.ClientHeartBeatRequest true "Client request"
// @Produce json
// @Success 200 {object} schema.ClientHeartBeatResponse
// @Router /heartbeat [post]
func HeartBeat(c *gin.Context) {
	var heartBeatRequest schema.ClientHeartBeatRequest
	if err := c.BindJSON(&heartBeatRequest); err != nil {
		c.JSON(http.StatusOK, schema.ClientHeartBeatResponse{
			Status:  schema.ERR,
			Serve:   false,
			Fee:     0,
			Message: err.Error(),
		})
		return
	}
	prevRoom := &model.Room{RoomID: 0}
	db.GetDataBase().Where("room_name = ?", heartBeatRequest.Room).First(&prevRoom)
	if prevRoom.IsPowerOn == false && heartBeatRequest.Power == true {
		userID, _ := service.GetCheckInCheckOutMap().FindUserByRoom(prevRoom.RoomID)
		user := &model.User{UserID: 0}
		db.GetDataBase().Where("user_id = ?", userID).First(&user)
		var task = &air.Task{
			RoomID:              prevRoom.RoomID,
			LastModifiedTime:    time.Now(),
			Priority:            user.Priority,
			DefaultPriority:     user.Priority,
			CurrentTemperature:  heartBeatRequest.CurrentTemperature,
			TargetTemperature:   heartBeatRequest.TargetTemperature,
			FanSpeed:            heartBeatRequest.FanSpeed,
			CurrentServiceCount: 0,
		}
		air.GetAir().AddTask(task)
		prevRoom.LastOnTime = time.Now()
	} else if prevRoom.IsPowerOn == true && heartBeatRequest.Power == false {
		task := air.GetAir().GetTaskByRoomID(prevRoom.RoomID)
		air.GetAir().RemoveTaskByRoomID(prevRoom.RoomID)
		air.GenerateShutdownTicket(&task)
	} else if prevRoom.IsPowerOn == true && heartBeatRequest.Power == true {
		air.GetAir().UpdateTaskByRoomID(prevRoom.RoomID, heartBeatRequest.CurrentTemperature, heartBeatRequest.TargetTemperature, heartBeatRequest.FanSpeed)
	}
	prevRoom.FanSpeed = heartBeatRequest.FanSpeed
	prevRoom.IsPowerOn = heartBeatRequest.Power
	prevRoom.TargetTemperature = heartBeatRequest.TargetTemperature
	prevRoom.CurrentTemperature = heartBeatRequest.CurrentTemperature
	prevRoom.IsServicing = air.GetAir().GetServeOptionByRoomID(prevRoom.RoomID)
	db.GetDataBase().Save(&prevRoom)
	c.JSON(http.StatusOK, schema.ClientHeartBeatResponse{
		Status:  schema.ACK,
		Serve:   prevRoom.IsServicing,
		Fee:     service.GetTotalFeeByRoomID(prevRoom.RoomID),
		Message: "Done.",
	})
}
