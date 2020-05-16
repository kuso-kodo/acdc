package client

import "github.com/gin-gonic/gin"

func BindAPIRouters(router *gin.RouterGroup) {
	router.POST("/register", Register)
	router.POST("/heartbeat", HeartBeat)
}
