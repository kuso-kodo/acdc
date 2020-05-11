package service

import (
	"github.com/name1e5s/acdc/db"
	"github.com/name1e5s/acdc/model"
)

func CheckAuth(phone, password string) (bool, model.User) {
	user := &model.User{
		Role: model.InvalidMask,
	}
	db.GetDataBase().Where(&model.User{
		Password: password,
		Phone:    phone,
	}).FirstOrInit(&user)
	return user.Role != model.InvalidMask, *user
}
