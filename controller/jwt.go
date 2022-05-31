package controller

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt"
)

type JwtAuth struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var JWT_TOKEN_KEY string = "YykbosiFS7SmYtPfPZGUzetR2$Bs5WUK"

func init() {
	if v := os.Getenv("JWT_TOKEN_KEY"); len(v) > 0 {
		JWT_TOKEN_KEY = v
	}
}

// GenToken 根据JwtAuth结构体内的信息生成jwt字符串
// 使用前应先设置结构体中的字段
// https://www.ruanyifeng.com/blog/2018/07/json_web_token-tutorial.html
func (auth *JwtAuth) GenToken() (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, auth).SignedString([]byte(JWT_TOKEN_KEY))
}

func (auth *JwtAuth) ParseToken(token string) (username string, err error) {
	tok, err := jwt.ParseWithClaims(token, auth, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(JWT_TOKEN_KEY), nil
	})

	if err != nil {
		return
	}

	if c, ok := tok.Claims.(*JwtAuth); ok && tok.Valid {
		auth = c
		return auth.Username, nil
	}

	return "", errors.New("Invalid Token")
}
