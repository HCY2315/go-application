package router

import (
	"github.com/gin-gonic/gin"
	jwt "go-application/pkg/jwtauth"
)

var routerNoCheckRoleV1 = make([]func(*gin.RouterGroup), 0)

// 初始化业务路由 InitBusinessRouter
func InitBusinessRouter(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) *gin.Engine {
	// TODO:需要认证的路由
	// BusinessNoCheckRoleRouter(r, authMiddleware)

	// 无需认证的路由
	BusinessNoCheckRoleRouter(r)

	return r
}

// 无需认证的业务路由 BusinessNoCheckRoleRouter
func BusinessNoCheckRoleRouter(r *gin.Engine) {
	v1 := r.Group("/api/v1")

	// 测试接口
	v1.GET("/nilcheckrole", nil)

	for _, f := range routerNoCheckRoleV1 {
		f(v1)
	}
}
