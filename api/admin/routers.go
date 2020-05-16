package admin

import (
	"github.com/gin-gonic/gin"
)

func BindAPIRouters(router *gin.RouterGroup) {
	apiAdmin := router.Group("/admin")
	BindLoginRouters(apiAdmin)
}
