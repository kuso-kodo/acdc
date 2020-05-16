package receptionist

import (
	"github.com/gin-gonic/gin"
)

func BindAPIRouters(router *gin.RouterGroup) {
	apiReceptionist := router.Group("/receptionist")
	BindTicketRouters(apiReceptionist)
	BindCheckInOutRouters(apiReceptionist)
}
