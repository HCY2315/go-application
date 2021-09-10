package middleware

import (
	"go-application/app/admin/models"
	"go-application/app/admin/models/system"
	"go-application/common/log"
	"go-application/tools"
	"go-application/tools/config"
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

		// 写入操作日志表
		if c.Request.Method != "GET" && c.Request.Method != "OPTIONS" && config.LoggerConfig.EnabledDB {
			SetDBOperLog(c, clientIP, statusCode, reqUri, reqMethod, latencyTime)
		}
	}
}

// SetDBOperLog 写入操作日志表
func SetDBOperLog(c *gin.Context, clientIP string, statusCode int, reqUri string, reqMethod string, latencyTime time.Duration) {
	menu := models.Menu{}
	menu.Path = reqUri
	menu.Action = reqMethod
	menuList, err := menu.Get()
	if err != nil {
		log.Error("获取菜单数据失败！err:", err)
		return
	}
	log.Info("打印查询的菜单目录", menuList)
	sysOperaLog := system.SysOperaLog{}
	sysOperaLog.OperIp = clientIP
	sysOperaLog.OperLocation, err = tools.GetLocation(clientIP)
	if err != nil {
		log.Error("获取外网IP地址失败， err", err)
		return
	}
}
