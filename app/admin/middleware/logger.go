package middleware

import (
	"go-application/common/log"
	"time"

	"github.com/gin-gonic/gin"
)

// LoggerToFile 日志记录到文件; HandlerFunc将gin中间件使用的处理程序定义为返回值
func LoggerToFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求路由
		reqUri := c.Request.RequestURI

		// 请求方式
		reqMethod := c.Request.Method

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		// 日志格式
		logData := map[string]interface{}{
			"statusCode":  statusCode,
			"latencyTime": latencyTime,
			"clientIP":    clientIP,
			"method":      reqMethod,
			"uri":         reqUri,
		}
		log.Info(logData)

		// TODO: 写入操作日志表

	}
}
