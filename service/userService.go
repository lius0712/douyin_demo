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
	Uid      int64
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

func (u *UserRegisterService) Register() (entity.User, error) {
	var user entity.User
	user.Name = u.Username
	salt, err := uuid.NewRandom()
	if err != nil {
		return user, err
	}
	user.Salt = salt.String()
	user.Password = hashPasswd(u.Password, user.Salt)

	err = repository.NewUserDao().Register(&user)

	return user, err
}

func (u *UserInfo) UserLogin() (entity.User, error) {
	user, err := repository.NewUserDao().UserInfoByName(u.Username)
	if err != nil {
		return user, errors.New("User does not exist")
	}

	if !checkPasswd(u.Password, user.Salt, user.Password) {
		return user, errors.New("Username or password is wrong")
	}

	return user, nil
}

//根据用户名获取用户信息

func (u *UserInfo) UserInfoByName() (entity.User, error) {
	user, err := repository.NewUserDao().UserInfoByName(u.Username)
	return user, err
}

//根据用户id获取用户信息

func (u *UserInfo) UserInfoByUid() (entity.User, error) {
	user, err := repository.NewUserDao().UserInfoByUid(u.Uid)
	return user, err
}
