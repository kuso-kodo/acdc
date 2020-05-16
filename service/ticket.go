package service

import (
	"github.com/gin-gonic/gin"
	"github.com/name1e5s/acdc/db"
	"github.com/name1e5s/acdc/model"
	"github.com/name1e5s/acdc/schema"
	"net/http"
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
