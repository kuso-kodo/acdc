package admin

import (
	"encoding/json"
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/name1e5s/acdc/db"
	"github.com/name1e5s/acdc/middleware"
	"github.com/name1e5s/acdc/model"
	"github.com/name1e5s/acdc/schema"
	"net/http"
)

// @Summary List all users.
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {array} model.Admin
// @Router /admin/user/all [get]
func GetAllUsers(c *gin.Context) {
	var users []model.Admin
	db.GetDataBase().Find(&users)
	c.JSON(http.StatusOK, users)
}

// @Summary Add a new user.
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param userRequest body schema.AddNewAdminRequest true "User request"
// @Success 200 {object} schema.AddNewAdminResponse
// @Failure 406 {object} schema.AddNewAdminResponse
// @Failure 500 {object} schema.AddNewAdminResponse
// @Router /admin/user/new [post]
func CreateNewUser(c *gin.Context) {
	var adminRequest schema.AddNewAdminRequest
	if err := c.ShouldBind(&adminRequest); err != nil {
		schema.NewCommonFailureSchema(c, http.StatusNotAcceptable, "Incomplete user information.")
		return
	}
	adminModel := model.Admin{
		UserName: adminRequest.UserName,
		Password: adminRequest.Password,
		Role:     adminRequest.Role,
	}
	db.GetDataBase().Create(&adminModel)
	user := &model.Admin{
		Role: model.InvalidMask,
	}
	db.GetDataBase().Where(&adminModel).FirstOrInit(&user)
	if user.Role == model.InvalidMask {
		schema.NewCommonFailureSchema(c, http.StatusInternalServerError, "Create new user failed.")
		return
	} else {
		schema.NewCommonFailureSchema(c, http.StatusOK, "Done.")
		return
	}
}

// @Summary Delete a user.
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param userInfo body schema.DeleteAdminRequest true "User request"
// @Success 200 {object} schema.DeleteAdminResponse
// @Failure 406 {object} schema.DeleteAdminResponse
// @Router /admin/user/delete [post]
func DeleteUser(c *gin.Context) {
	var userInfo schema.DeleteAdminRequest
	if err := c.BindJSON(&userInfo); err != nil {
		schema.NewCommonFailureSchema(c, http.StatusNotAcceptable, "Incomplete user information.")
		return
	}
	db.GetDataBase().Where("username = ?", userInfo.UserName).Delete(&model.Admin{})
	schema.NewCommonFailureSchema(c, http.StatusOK, "Done.")
}

// @Summary Get current user info.
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} model.Admin
// @Failure 401 {object} schema.CommonFailureSchema
// @Router /admin/me [get]
func GetCurrentUser(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user := model.Admin{}
	err := json.Unmarshal([]byte(claims["payload"].(string)), &user)
	if err != nil {
		schema.NewCommonFailureSchema(c, http.StatusUnauthorized, "Invalid user token.")
		return
	}
	c.JSON(http.StatusOK, user)
}

func BindUserRouters(router *gin.RouterGroup) {
	user := router.Group("/user")
	user.Use(middleware.JWTSuperUserAuthenticator().MiddlewareFunc())
	user.GET("/all", GetAllUsers)
	user.POST("/new", CreateNewUser)
	user.POST("/delete", DeleteUser)
	router.GET("/me", middleware.JWTBaseAuthenticator().MiddlewareFunc(), GetCurrentUser)
}
