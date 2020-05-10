package api

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/name1e5s/acdc/api/v1"
)

func BindAPIRouters(router *gin.RouterGroup) {
	apiV1 := router.Group("/v1")
	apiV1.POST("/login", v1.Login)
}
