package service

import (
	"github.com/name1e5s/acdc/db"
	"github.com/name1e5s/acdc/model"
)

func CheckAdminAuth(username, password string) (bool, model.Admin) {
	user := &model.Admin{
		Role: model.InvalidMask,
	}
	db.GetDataBase().Where(&model.Admin{
		Password: password,
		UserName: username,
	}).FirstOrInit(&user)
	return user.Role != model.InvalidMask, *user
}

func CheckUserAuth(phone, password string) (bool, model.User) {
	user := &model.User{UserID: 0}
	db.GetDataBase().Where(&model.User{
		Password: password,
		Phone:    phone,
	}).First(&user)
	return user.UserID != 0, *user
}
