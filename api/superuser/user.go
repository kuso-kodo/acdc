package superuser

import (
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
// @Router /superuser/all [get]
func GetAllAdmin(c *gin.Context) {
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
// @Failure 403 {object} schema.AddNewAdminResponse
// @Router /superuser/new [post]
func CreateNewAdmin(c *gin.Context) {
	var adminRequest schema.AddNewAdminRequest
	if err := c.ShouldBind(&adminRequest); err != nil {
		schema.NewCommonStatusSchema(c, http.StatusForbidden, "Incomplete user information.")
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
		schema.NewCommonStatusSchema(c, http.StatusForbidden, "Create new user failed.")
		return
	} else {
		schema.NewCommonStatusSchema(c, http.StatusOK, "Done.")
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
// @Router /superuser/delete [post]
func DeleteAdmin(c *gin.Context) {
	var userInfo schema.DeleteAdminRequest
	if err := c.BindJSON(&userInfo); err != nil {
		schema.NewCommonStatusSchema(c, http.StatusNotAcceptable, "Incomplete user information.")
		return
	}
	db.GetDataBase().Where("username = ?", userInfo.UserName).Delete(&model.Admin{})
	schema.NewCommonStatusSchema(c, http.StatusOK, "Done.")
}

func BindSuperUserRouters(router *gin.RouterGroup) {
	user := router.Group("/superuser")
	user.Use(middleware.JWTSuperUserAuthenticator().MiddlewareFunc())
	user.GET("/all", GetAllAdmin)
	user.POST("/new", CreateNewAdmin)
	user.POST("/delete", DeleteAdmin)
}
