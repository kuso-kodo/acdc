package api

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/name1e5s/acdc/api/admin"
)

func BindAPIRouters(router *gin.RouterGroup) {
	apiAdmin := router.Group("/admin")
	apiAdmin.POST("/login", v1.Login)
	v1.BindUserRouters(apiAdmin)
}
