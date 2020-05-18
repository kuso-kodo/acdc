package api

import (
	"github.com/gin-gonic/gin"
	"github.com/name1e5s/acdc/api/accounting"
	"github.com/name1e5s/acdc/api/admin"
	"github.com/name1e5s/acdc/api/client"
	"github.com/name1e5s/acdc/api/maintainer"
	"github.com/name1e5s/acdc/api/receptionist"
	"github.com/name1e5s/acdc/api/superuser"
	"github.com/name1e5s/acdc/api/user"
)

func BindAPIRouters(router *gin.RouterGroup) {
	admin.BindAPIRouters(router)
	user.BindAPIRouters(router)
	superuser.BindSuperUserRouters(router)
	receptionist.BindAPIRouters(router)
	maintainer.BindAPIRouters(router)
	client.BindAPIRouters(router)
	accounting.BindRouters(router)
}
