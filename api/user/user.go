package user

import (
	"github.com/gin-gonic/gin"
	"github.com/name1e5s/acdc/db"
	"github.com/name1e5s/acdc/middleware"
	"github.com/name1e5s/acdc/model"
	"github.com/name1e5s/acdc/schema"
	"github.com/name1e5s/acdc/service"
	"net/http"
)

// @Summary Perform login.
// @Accept  json
// @Produce  json
// @Param userRequest body schema.AuthLoginRequest true "User request"
// @Success 200 {object} schema.AuthLoginResponse
// @Failure 401 {object} schema.CommonFailureSchema
// @Router /user/login [post]
func Login(c *gin.Context) {
	// Note that the username is phone. Not the real user name.
	middleware.JWTUserAuthenticator().LoginHandler(c)
}

// @Summary Register a new user.
// @Accept  json
// @Produce  json
// @Param userRequest body schema.UserRegisterRequest true "User request"
// @Success 200 {object} schema.CommonFailureSchema
// @Failure 401 {object} schema.CommonFailureSchema
// @Router /user/register [post]
func Register(c *gin.Context) {
	var userReq schema.UserRegisterRequest
	if err := c.ShouldBind(&userReq); err != nil {
		schema.NewCommonFailureSchema(c, http.StatusForbidden, "Incomplete user information.")
		return
	}
	userInfo := model.User{
		UserName: userReq.UserName,
		Password: userReq.Password,
		Phone:    userReq.Phone,
	}
	err := db.GetDataBase().Create(&userInfo).Error
	if err != nil {
		schema.NewCommonFailureSchema(c, http.StatusForbidden, err.Error())
		return
	}
	schema.NewCommonFailureSchema(c, http.StatusOK, "Done.")
}

// @Summary List all users.
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {array} model.User
// @Router /user/all [get]
func GetAllUser(c *gin.Context) {
	var users []model.User
	db.GetDataBase().Find(&users)
	c.JSON(http.StatusOK, users)
}

// @Summary Get current user info.
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} model.User
// @Failure 401 {object} schema.CommonFailureSchema
// @Router /user/me [get]
func GetCurrentUser(c *gin.Context) {
	user, err := service.GetUserFromClaims(c)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, user)
}

func BindUserRouters(router *gin.RouterGroup) {
	router.POST("/login", Login)
	router.POST("/register", Register)
	router.GET("/me", middleware.JWTUserAuthenticator().MiddlewareFunc(), GetCurrentUser)
	router.GET("/all", middleware.JWTSuperUserAuthenticator().MiddlewareFunc(), GetAllUser)
}
