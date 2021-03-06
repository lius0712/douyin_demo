package middleware

import (
	"github.com/gin-gonic/gin"
)

type Auth interface {
	GenToken() (string, error)
	ParseToken(token string) error
	GetUsername() string
	GetUid() int64
}

func AuthMiddleware(auth Auth, optional bool) func(c *gin.Context) {
	return func(c *gin.Context) {
		var token string
		if token = c.Query("token"); len(token) == 0 {
			token = c.PostForm("token")
		}
		err := auth.ParseToken(token)
		if err != nil && !optional {
			//c.AbortWithStatusJSON(
			//	http.StatusUnauthorized,
			//	controller.Response{
			//		StatusCode: -1,
			//		StatusMsg:  "Session Expired, Please Relogin.",
			//	},
			//)
			return
		}

		c.Set("username", auth.GetUsername())
		c.Set("uid", auth.GetUid())
		c.Next()
	}
}
