package service

import (
	"github.com/gin-gonic/gin"
	"github.com/name1e5s/acdc/db"
	"github.com/name1e5s/acdc/model"
	"github.com/name1e5s/acdc/schema"
	"net/http"
	"time"
)

func GetAllTicketByUserID(c *gin.Context, userID uint, pageSize int, offset int) {
	var tickets []model.Ticket
	err := db.GetDataBase().Where("user_refer = ?", userID).Limit(pageSize).Offset((offset - 1) * pageSize).Find(&tickets).Error
	if err != nil {
		schema.NewCommonStatusSchema(c, http.StatusForbidden, err.Error())
		return
	}
	c.JSON(http.StatusOK, tickets)
}

func GetUnpaidTicketByUserID(c *gin.Context, userID uint, pageSize int, offset int) {
	var tickets []model.Ticket
	err := db.GetDataBase().Where("user_refer = ?", userID).Where("paid = ?", false).Limit(pageSize).Offset((offset - 1) * pageSize).Find(&tickets).Error
	if err != nil {
		schema.NewCommonStatusSchema(c, http.StatusForbidden, err.Error())
		return
	}
	c.JSON(http.StatusOK, tickets)
}

func GetTotalFeeByUserID(c *gin.Context, userID uint) {
	var tickets []model.Ticket
	err := db.GetDataBase().Where("user_refer = ?", userID).Where("paid = ?", false).Find(&tickets).Error
	if err != nil {
		schema.NewCommonStatusSchema(c, http.StatusForbidden, err.Error())
		return
	}
	var result float32 = 0
	for _, value := range tickets {
		result += value.TotalFee
	}
	c.JSON(http.StatusOK, schema.UserTotalFeeResponse{Fee: result})
}

func GetTotalFeeByRoomID(roomID uint) float32 {
	var tickets []model.Ticket
	db.GetDataBase().Where("room_refer = ?", roomID).Where("paid = ?", false).Find(&tickets)
	var result float32 = 0
	for _, value := range tickets {
		result += value.TotalFee
	}
	return result
}

func GetResultsByTime(startTime, endTime time.Time) ([]model.Report, error) {
	var result []model.Report
	err := db.GetDataBase().Table("tickets").Select("(select room_name from rooms as T where T.room_id = room_refer) as room_name, sum(end_at - start_at) as total_time, sum(service_count) as service, count(*) as ticket, sum(fan_speed_changed) as fan, sum(shutdown) as power, sum(priority_changed) as priority, sum(total_fee) as fee").Where("room_refer <> ?", 1).Where("end_at between ? and ?", startTime, endTime).Group("room_refer").Scan(&result).Error
	if err != nil {
		return nil, err
	}
	return result, err
}
