package maintainer

import (
	"github.com/gin-gonic/gin"
	"github.com/name1e5s/acdc/db"
	"github.com/name1e5s/acdc/middleware"
	"github.com/name1e5s/acdc/model"
	"net/http"
)

// @Summary List all rooms.
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {array} model.Room
// @Failure 401 {object} schema.CommonStatusSchema
// @Failure 403 {object} schema.CommonStatusSchema
// @Router /maintainer/room [get]
func GetAllRoom(c *gin.Context) {
	var rooms []model.Room
	db.GetDataBase().Find(&rooms)
	c.JSON(http.StatusOK, rooms)
}

func BindRoomRouters(router *gin.RouterGroup) {
	router.GET("/room", middleware.JWMaintainerAuthenticator().MiddlewareFunc(), GetAllRoom)
}
