package user

import (
	"github.com/gin-gonic/gin"
)

func BindAPIRouters(router *gin.RouterGroup) {
	apiUser := router.Group("/user")
	BindUserRouters(apiUser)
	BindTicketRouters(apiUser)
}
