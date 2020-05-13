package api

import (
	"github.com/gin-gonic/gin"
	"github.com/name1e5s/acdc/api/admin"
)

func BindAPIRouters(router *gin.RouterGroup) {
	admin.BindAPIRouters(router)
}
