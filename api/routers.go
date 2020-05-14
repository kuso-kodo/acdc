package api

import (
	"github.com/gin-gonic/gin"
	"github.com/name1e5s/acdc/api/admin"
	"github.com/name1e5s/acdc/api/user"
)

func BindAPIRouters(router *gin.RouterGroup) {
	admin.BindAPIRouters(router)
	user.BindAPIRouters(router)
}
