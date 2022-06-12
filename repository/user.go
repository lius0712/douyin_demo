package repository

import (
	"github.com/RaymondCode/simple-demo/entity"
	"sync"
)

type UserDao struct {
}

var userDao *UserDao
var userOnce sync.Once

func NewUserDao() *UserDao {
	userOnce.Do(func() {
		userDao = new(UserDao)
	})
	return userDao
}

//注册用户信息
func (u *UserDao) Register(user *entity.User) error {
	err := DB.Create(&user).Error
	return err
}

//根据用户名查找用户信息
func (u *UserDao) UserInfoByName(username string) (entity.User, error) {
	var user entity.User
	err := DB.Where(&entity.User{Name: username}).First(&user).Error
	return user, err
}

//根据用户id查找用户信息
func (u *UserDao) UserInfoByUid(uid int64) (entity.User, error) {
	var user entity.User
	err := DB.Where(&entity.User{ID: uid}).First(&user).Error
	return user, err
}
