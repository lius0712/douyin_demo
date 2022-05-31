package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Auth interface {
	GenToken() (string, error)
	ParseToken(token string) (username string, err error)
}

func AuthMiddleware(auth Auth) func(c *gin.Context) {
	return func(c *gin.Context) {
		var token string
		switch c.Request.Method {
		case "GET":
			token = c.Query("token")
		case "POST":
			token = c.PostForm("token")
		}
		username, err := auth.ParseToken(token)
		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusForbidden,
				Response{
					StatusCode: -1,
					StatusMsg:  "Session Expired, Please Relogin.",
				},
			)
		}

		c.Set("username", username)
		c.Next()
	}
}
