package handler

import (
	jwt "go-application/pkg/jwtauth"

	"github.com/mojocn/base64Captcha"
)

// DefaultMemStore是CAPTCHA的共享存储，由新函数生成
// var store = base64Captcha.DefaultMemStore
var store = base64Captcha.DefaultMemStore

func PayloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(map[string]interface{}); ok {
		return jwt.MapClaims{}
	}
}
