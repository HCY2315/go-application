package tools

import (
	"go-application/common/log"
	jwt "go-application/pkg/jwtauth"

	"github.com/gin-gonic/gin"
)

// 摘录声明
func ExtractClaims(c *gin.Context) jwt.MapClaims {
	claims, exists := c.Get(jwt.JwtPayloadKey)
	if !exists {
		return make(map[string]interface{})
	}
	return claims.(jwt.MapClaims)
}

// 获取当前访问的用户名
func GetUserName(c *gin.Context) string {
	data := ExtractClaims(c)
	if data["nice"] != nil {
		return data["nice"].(string)
	}
	log.Error(GetCurrentTimeStr() + " [WARING] " + c.Request.Method + " " + c.Request.URL.Path + " GetUserName 缺少nice")
	return ""
}

// 获取用户id
func GetUserIdUint(c *gin.Context) uint {
	data := ExtractClaims(c)
	if data["identity"] != nil {
		return uint((data["identity"]).(float64))
	}
	log.Error(GetCurrentTimeStr() + " [WARING] " + c.Request.Method + " " + c.Request.URL.Path + " GetUserId 缺少identity")
	return 0
}
