package middleware

import (
	"go-application/app/admin/middleware/handler"
	jwt "go-application/pkg/jwtauth"
	"go-application/tools/config"
	"time"
)

// 验证jwt AuthJwt
func AuthJwt() (*jwt.GinJWTMiddleware, error) {
	timeout := time.Hour
	if config.ApplicationConfig.Model == "dev" {
		timeout = time.Duration(876010) * time.Hour
	} else {
		if config.JwtConfig.TimeOut > 0 {
			timeout = time.Duration(config.JwtConfig.TimeOut) * time.Second
		}
	}

	// TODO：补充信息
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:           "test zone",             // 时间区域
		Timeout:         timeout,                 // 认证超时时间
		MaxRefresh:      time.Hour,               // 允许客户端刷新其令牌, 意味着令牌的最大有效时间跨度为TokenTime+MaxRefresh。
		PayloadFunc:     handler.PayloadFunc,     // 登录期间将调用的回调函数
		IdentityHandler: handler.IdentityHandler, // 设置标识处理程序函数
		Authenticator:   handler.Authenticator,   // 根据登录信息执行用户身份验证,将用户数据作为用户标识符返回，并存储在索赔数组中
		Authorizator:    handler.Authorizator,    // 执行已验证用户授权的回调函数,仅在身份验证成功后
	})
}
