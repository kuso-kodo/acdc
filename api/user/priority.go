package user

import (
	"github.com/gin-gonic/gin"
	"github.com/name1e5s/acdc/air"
	"github.com/name1e5s/acdc/db"
	"github.com/name1e5s/acdc/middleware"
	"github.com/name1e5s/acdc/model"
	"github.com/name1e5s/acdc/schema"
	"github.com/name1e5s/acdc/service"
	"net/http"
)

// @Summary Set user priority.
// @Security ApiKeyAuth
// @Param level query string false "Priority Level"
// @Produce json
// @Success 200 {object} schema.CommonStatusSchema
// @Failure 401 {object} schema.CommonStatusSchema
// @Failure 403 {object} schema.CommonStatusSchema
// @Router /user/priority [post]
func SetPriority(c *gin.Context) {
	service.UserHandlerWrapper(c, func(c *gin.Context, user model.User) {
		level, err := service.GetQueryPriority(c)
		if err != nil {
			schema.NewCommonStatusSchema(c, http.StatusForbidden, err.Error())
			return
		}
		err = db.GetDataBase().Model(model.User{}).Where("user_id = ?", user.UserID).Update("priority", level).Error
		if err != nil {
			schema.NewCommonStatusSchema(c, http.StatusForbidden, err.Error())
			return
		}
		roomID, ok := service.GetCheckInCheckOutMap().FindRoomByUser(user.UserID)
		if ok {
			air.GetAir().UpdatePriorityByRoomID(roomID, level)
		}
		schema.NewCommonStatusSchema(c, http.StatusOK, "Done.")
	})
}

func BindPriorityRouters(router *gin.RouterGroup) {
	router.POST("/priority", middleware.JWTUserAuthenticator().MiddlewareFunc(), SetPriority)
}
