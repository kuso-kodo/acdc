package service

import (
	"github.com/name1e5s/acdc/db"
	"github.com/name1e5s/acdc/model"
)

func CheckAuth(phone, password string) (bool, *model.User) {
	var user model.User
	db.GetDataBase().Where(&model.User{
		Password: password,
		Phone:    phone,
	}).FirstOrInit(&user, &model.User{
		Role: model.InvalidMask,
	})
	return user.Role == model.InvalidMask, &user
}
