package middleware

import (
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
		Realm:   "test zone",
		Timeout: timeout,
	})
}
