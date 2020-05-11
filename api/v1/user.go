package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/name1e5s/acdc/db"
	"github.com/name1e5s/acdc/middleware"
	"github.com/name1e5s/acdc/model"
	"net/http"
)

// @Summary List all users.
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {array} model.User
// @Router /user/all [get]
func GetAllUsers(c *gin.Context) {
	var users []model.User
	db.GetDataBase().Find(&users)
	c.JSON(http.StatusOK, users)
}

func BindUserRouters(router *gin.RouterGroup) {
	user := router.Group("/user")
	user.Use(middleware.JWTSuperUserAuthenticator().MiddlewareFunc())
	user.GET("/all", GetAllUsers)
}
