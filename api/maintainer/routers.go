package maintainer

import "github.com/gin-gonic/gin"

func BindAPIRouters(router *gin.RouterGroup) {
	apiMaintainer := router.Group("/maintainer")
	BindAirRouters(apiMaintainer)
	BindRoomRouters(apiMaintainer)
}
