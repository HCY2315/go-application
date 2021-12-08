package middleware

import (
	"go-application/common/middleware"

	"github.com/gin-gonic/gin"
)

func InitMiddleware(r *gin.Engine) {
	// 日志处理
	r.Use(LoggerToFile())

	// 自定义错误处理

	// NoCache is a middleware function that appends headers
	// NoCache是一个附加头的中间件函数

	// 跨域处理

	// 链路追踪
	r.Use(middleware.Trace())
}
