package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/name1e5s/acdc/middleware"
)

func Login(c *gin.Context) {
	middleware.JWTBaseAuthenticator().LoginHandler(c)
}
