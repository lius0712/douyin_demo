package service

import (
	"github.com/RaymondCode/simple-demo/entity"
	"github.com/RaymondCode/simple-demo/repository"
)

//管理用户注册功能
type UserRegisterService struct {
	Username string
	Password string
}

//查询用户信息,bug：用户名不可重复，需解决
type UserInfo struct {
	Username string
	Password string
}

func (u *UserRegisterService) Register() error {
	var user entity.User
	user.Name = u.Username
	user.Password = u.Password

	err := repository.DB.Create(&user).Error

	return err
}

func (u *UserInfo) UserLogin() (entity.User, error) {
	var user entity.User
	err := repository.DB.Where(&entity.User{Name: u.Username, Password: u.Password}).First(&user).Error
	return user, err
}

func (u *UserInfo) UserInfoByName() (entity.User, error) {
	var user entity.User
	err := repository.DB.First(&user, "username=?", u.Username).Error
	return user, err
}
