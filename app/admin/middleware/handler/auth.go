package handler

import (
	"fmt"
	"go-application/app/admin/models"
	jwt "go-application/pkg/jwtauth"
	"go-application/tools/config"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

// DefaultMemStore是CAPTCHA的共享存储，由新函数生成
var store = base64Captcha.DefaultMemStore

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

// @Summary 登陆
// @Description 获取token
// @Description LoginHandler can be used by clients to get a jwt token.
// @Description Payload needs to be json in the form of {"username": "USERNAME", "password": "PASSWORD"}.
// @Description Reply will be of the form {"token": "TOKEN"}.
// @Description dev mode：It should be noted that all fields cannot be empty, and a value of 0 can be passed in addition to the account password
// @Description 注意：开发模式：需要注意全部字段不能为空，账号密码外可以传入0值
// @Accept  application/json
// @Product application/json
// @Param account body models.Login  true "account"
// @Success 200 {string} string "{"code": 200, "expire": "2019-08-07T12:45:48+08:00", "token": ".eyJleHAiOjE1NjUxNTMxNDgsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTU2NTE0OTU0OH0.-zvzHvbg0A" }"
// @Router /login [post]
func Authenticator(c *gin.Context) (interface{}, error) {
	var LoginVals models.Login
	// var status = "0"
	// var msg = "登录成功"
	// var username = ""

	if err := c.ShouldBind(&LoginVals); err != nil {
		// TODO: 登录失败的日志
		fmt.Println(err.Error())
		return nil, jwt.ErrMissingLoginValues
	}

	if config.ApplicationConfig.Model != "dev" {
		if !store.Verify(LoginVals.UUID, LoginVals.Code, true) {
			//TODO: 登录失败的日志
			return nil, jwt.ErrInvalidVerificationode
		}
	}

	user, role, e := LoginVals.GetUser()
	if e == nil {
		// TODO: 登录成功的日志
		return map[string]interface{}{"user": user, "role": role}, nil
	} else {
		//TODO：日志输出
	}
	return nil, jwt.ErrFailedAuthentication
}
