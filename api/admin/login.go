package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/name1e5s/acdc/middleware"
)

// @Summary Perform login.
// @Accept  json
// @Produce  json
// @Param userRequest body schema.AuthLoginRequest true "User request"
// @Success 200 {object} schema.AuthLoginResponse
// @Failure 401 {object} schema.CommonFailureSchema
// @Router /admin/login [post]
func Login(c *gin.Context) {
	middleware.JWTBaseAuthenticator().LoginHandler(c)
}
