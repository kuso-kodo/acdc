package admin

import (
	"encoding/json"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/name1e5s/acdc/middleware"
	"github.com/name1e5s/acdc/model"
	"github.com/name1e5s/acdc/schema"
	"net/http"
)

// @Summary Perform login.
// @Accept  json
// @Produce  json
// @Param userRequest body schema.AuthLoginRequest true "User request"
// @Success 200 {object} schema.AuthLoginResponse
// @Failure 401 {object} schema.CommonStatusSchema
// @Router /admin/login [post]
func Login(c *gin.Context) {
	middleware.JWTBaseAuthenticator().LoginHandler(c)
}

// @Summary Get current user info.
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} model.Admin
// @Failure 401 {object} schema.CommonStatusSchema
// @Router /admin/me [get]
func GetCurrentAdmin(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user := model.Admin{}
	err := json.Unmarshal([]byte(claims["payload"].(string)), &user)
	if err != nil {
		schema.NewCommonStatusSchema(c, http.StatusUnauthorized, "Invalid user token.")
		return
	}
	c.JSON(http.StatusOK, user)
}

func BindLoginRouters(router *gin.RouterGroup) {
	router.POST("/login", Login)
	router.GET("/me", middleware.JWTBaseAuthenticator().MiddlewareFunc(), GetCurrentAdmin)
}
