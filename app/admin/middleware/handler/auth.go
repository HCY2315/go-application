package handler

import (
	"go-application/app/admin/models"
	jwt "go-application/pkg/jwtauth"

	"github.com/gin-gonic/gin"
	// "github.com/mojocn/base64Captcha"
)

// DefaultMemStore是CAPTCHA的共享存储，由新函数生成
// var store = base64Captcha.DefaultMemStore

func PayloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(map[string]interface{}); ok {
		v, _ := v["user"].(models.SysUser)
		return jwt.MapClaims{
			jwt.IdentityKey: v.UserId,
			// TODO：角色等认证
		}
	}
	return jwt.MapClaims{}
}

func IdentityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return map[string]interface{}{
		"IdentityKey": claims["identity"],
		// TODO：
	}
}
