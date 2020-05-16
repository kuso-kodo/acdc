package receptionist

import (
	"github.com/gin-gonic/gin"
	"github.com/name1e5s/acdc/db"
	"github.com/name1e5s/acdc/middleware"
	"github.com/name1e5s/acdc/model"
	"github.com/name1e5s/acdc/schema"
	"github.com/name1e5s/acdc/service"
	"net/http"
)

// @Summary Check in user.
// @Security ApiKeyAuth
// @Param phone query string false "User phone number"
// @Param room query string false "Room name"
// @Produce json
// @Success 200 {object} schema.CommonStatusSchema
// @Failure 401 {object} schema.CommonStatusSchema
// @Failure 403 {object} schema.CommonStatusSchema
// @Router /receptionist/checkin [POST]
func CheckIn(c *gin.Context) {
	service.ReceptionistHandlerWrapper(c, func(c *gin.Context, userID uint) {
		roomName, err := service.GetQueryRoomName(c)
		if err != nil {
			schema.NewCommonStatusSchema(c, http.StatusForbidden, err.Error())
			return
		}
		room, ok := service.GetRoomByName(roomName)
		if ok != true {
			schema.NewCommonStatusSchema(c, http.StatusForbidden, "room not found.")
			return
		}
		err = service.GetCheckInCheckOutMap().CheckIn(userID, room.RoomID)
		if err != nil {
			schema.NewCommonStatusSchema(c, http.StatusForbidden, err.Error())
			return
		}
		schema.NewCommonStatusSchema(c, http.StatusOK, "Done.")
	})
}

// @Summary Check out user.
// @Security ApiKeyAuth
// @Param phone query string false "User phone number"
// @Produce json
// @Success 200 {object} schema.CommonStatusSchema
// @Failure 401 {object} schema.CommonStatusSchema
// @Failure 403 {object} schema.CommonStatusSchema
// @Router /receptionist/checkout [POST]
func CheckOut(c *gin.Context) {
	service.ReceptionistHandlerWrapper(c, func(c *gin.Context, userID uint) {
		service.GetCheckInCheckOutMap().CheckOut(userID)
		err := db.GetDataBase().Model(model.Ticket{}).Where("user_refer = ?", userID).Where("paid = ?", false).Update("user_refer", 0).Error
		if err != nil {
			schema.NewCommonStatusSchema(c, http.StatusForbidden, err.Error())
			return
		}
		schema.NewCommonStatusSchema(c, http.StatusOK, "Done.")
	})
}

func BindCheckInOutRouters(router *gin.RouterGroup) {
	router.POST("/checkin", middleware.JWTReceptionistAuthenticator().MiddlewareFunc(), CheckIn)
	router.POST("/checkout", middleware.JWTReceptionistAuthenticator().MiddlewareFunc(), CheckOut)
}
