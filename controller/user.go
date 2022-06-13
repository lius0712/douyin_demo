package controller

import (
	"net/http"

	"github.com/RaymondCode/simple-demo/entity"
	"github.com/RaymondCode/simple-demo/middleware"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
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
	User User `json:"user"`
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
		u, err := registerService.Register()

		if err != nil {
			c.JSON(http.StatusOK, UserRegisterResponse{
				Response: Response{StatusCode: 1, StatusMsg: "Insert failed"},
			})
		} else {
			jwt := middleware.JwtAuth{
				Username: username,
				Uid:      u.ID,
			}
			token, err := jwt.GenToken()
			if err != nil {
				c.AbortWithStatusJSON(http.StatusOK, UserRegisterResponse{
					Response: Response{StatusCode: 1, StatusMsg: err.Error()},
				})
				return
			}
			userInfo := service.UserInfo{
				Username: username,
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

	userLogin := service.UserInfo{
		Username: username,
		Password: password, //密码明文，后续优化进行加密
	}
	user, err := userLogin.UserLogin()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	} else {
		jwt := middleware.JwtAuth{
			Username: username,
			Uid:      user.ID,
		}
		token, err := jwt.GenToken()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 1, StatusMsg: err.Error()},
			})
			return
		}
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   user.ID,
			Token:    token,
		})
	}
}

func UserInfo(c *gin.Context) {
	//username := c.GetString("username")
	//userId := c.Query("user_id") //个人信息页用户id
	loginId := c.GetInt64("uid") //登录用户id
	var user entity.User
	var err error
	userInfo := service.UserInfo{Uid: loginId}
	user, err = userInfo.UserInfoByUid()
	if err != nil {
		c.JSON(http.StatusUnauthorized, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	} else {
		//fmt.Println(user)
		//user.Name = token

		userDemo := UserByEntity(user)
		userDemo.IsFollow = true //自己对自己默认设置已关注

		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0},
			User:     userDemo,
		})
	}
}
