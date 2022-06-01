package service

import (
	"encoding/base64"
	"errors"

	"github.com/RaymondCode/simple-demo/entity"
	"github.com/RaymondCode/simple-demo/repository"

	"github.com/google/uuid"
	"golang.org/x/crypto/argon2"
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

func hashPasswd(password, salt string) string {
	return base64.RawStdEncoding.EncodeToString(
		argon2.IDKey([]byte(password), []byte(salt), 1, 64*1024, 4, 32),
	)
}

func checkPasswd(password, salt, hash string) bool {
	passwdHash := argon2.IDKey([]byte(password), []byte(salt), 1, 64*1024, 4, 32)
	return hash == base64.RawStdEncoding.EncodeToString(passwdHash)
}

func (u *UserRegisterService) Register() error {
	var user entity.User
	user.Name = u.Username
	salt, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	user.Salt = salt.String()
	user.Password = hashPasswd(u.Password, user.Salt)

	err = repository.DB.Create(&user).Error

	return err
}

func (u *UserInfo) UserLogin() (entity.User, error) {
	var user entity.User
	err := repository.DB.Where(&entity.User{Name: u.Username}).First(&user).Error
	if err != nil {
		return user, err
	}

	if !checkPasswd(u.Password, user.Salt, user.Password) {
		return user, errors.New("Username or password is wrong")
	}

	return user, nil
}

func (u *UserInfo) UserInfoByName() (entity.User, error) {
	var user entity.User
	err := repository.DB.First(&user, "username=?", u.Username).Error
	return user, err
}
