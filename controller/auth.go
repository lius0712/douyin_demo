package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Auth interface {
	GenToken() (string, error)
	ParseToken(token string) error
	Username() string
	Uid() int64
}

func AuthMiddleware(auth Auth) func(c *gin.Context) {
	return func(c *gin.Context) {
		var token string
		if token = c.Query("token"); len(token) == 0 {
			token = c.PostForm("token")
		}
		err := auth.ParseToken(token)
		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				Response{
					StatusCode: -1,
					StatusMsg:  "Session Expired, Please Relogin.",
				},
			)
		}

		c.Set("username", auth.Username())
		c.Set("uid", auth.Uid())
		c.Next()
	}
}
