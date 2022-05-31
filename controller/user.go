package controller

import (
	"github.com/RaymondCode/simple-demo/entity"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserRegisterResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User entity.User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	userByNameService := service.UserInfo{ //根据用户名进行查找
		Username: username,
	}

	_, errSelect := userByNameService.UserInfoByName()

	if errSelect == nil { //如果存在该用户名
		c.JSON(http.StatusOK, UserRegisterResponse{
			Response: Response{StatusCode: 1, StatusMsg: "用户名已存在！"},
		})
	} else {
		registerService := service.UserRegisterService{
			Username: username,
			Password: password,
		}
		token := username //暂定token为username
		//fmt.Println(registerService)
		err := registerService.Register()

		if err != nil {
			c.JSON(http.StatusOK, UserRegisterResponse{
				Response: Response{StatusCode: 1, StatusMsg: "Insert failed"},
			})
		} else {
			userInfo := service.UserInfo{
				Username: token,
			}
			user, _ := userInfo.UserInfoByName()
			c.JSON(http.StatusOK, UserRegisterResponse{
				Response: Response{StatusCode: 0},
				UserId:   user.ID,
				Token:    token,
			})
		}
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username

	userLogin := service.UserInfo{
		Username: username,
		Password: password, //密码明文，后续优化进行加密
	}
	user, err := userLogin.UserLogin()
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "username or password err"},
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   user.ID,
			Token:    token,
		})
	}
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")
	//userId := c.Query("user_id") //小修中...
	//fmt.Println("&&&&&&&&&&&&&&&&&")
	//fmt.Println(userId)
	//fmt.Println("&&&&&&&&&&&&&&&&&")
	var user entity.User
	var err error
	userInfo := service.UserInfo{Username: token}
	user, err = userInfo.UserInfoByName()
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	} else {
		//fmt.Println(user)
		//user.Name = token
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0},
			User:     user,
		})
	}
}
