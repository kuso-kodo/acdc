package service

import (
	"github.com/gin-gonic/gin"
	"github.com/name1e5s/acdc/db"
	"github.com/name1e5s/acdc/model"
	"github.com/name1e5s/acdc/schema"
	"net/http"
)

type ReceptionistHandler func(c *gin.Context, userID uint)

func ReceptionistHandlerWrapper(c *gin.Context, handler ReceptionistHandler) {
	phone, err := GetQueryUserPhone(c)
	if err != nil {
		schema.NewCommonStatusSchema(c, http.StatusForbidden, err.Error())
		return
	}
	user := &model.User{UserID: 0}
	err = db.GetDataBase().Where("phone = ?", phone).First(&user).Error
	if err != nil {
		schema.NewCommonStatusSchema(c, http.StatusForbidden, err.Error())
		return
	}
	handler(c, user.UserID)
}
