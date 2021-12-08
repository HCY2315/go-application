package middleware

import (
	"go-application/app/admin/models"
	"go-application/app/admin/models/system"
	"go-application/app/admin/service"
	"go-application/common/log"
	"go-application/tools"
	"go-application/tools/config"
	"strings"
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

	// 判断菜单是否存在
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
	sysOperaLog.Status = tools.IntToString(statusCode)
	sysOperaLog.OperName = tools.GetUserName(c)
	sysOperaLog.RequestMethod = reqMethod
	sysOperaLog.OperUrl = reqUri
	if strings.Contains(reqUri, "/login") {
		sysOperaLog.BusinessType = "10"
		sysOperaLog.Title = "用户登录"
		sysOperaLog.OperName = "-"
	} else if strings.Contains(reqUri, "/api/v1/logout") {
		sysOperaLog.BusinessType = "11"
	} else if strings.Contains(reqUri, "/api/v1/getCaptcha") {
		sysOperaLog.BusinessType = "12"
		sysOperaLog.Title = "验证码"
	} else {
		if reqMethod == "POST" {
			sysOperaLog.BusinessType = "1"
		} else if reqMethod == "PUT" {
			sysOperaLog.BusinessType = "2"
		} else if reqMethod == "DELETE" {
			sysOperaLog.BusinessType = "3"
		}
	}
	sysOperaLog.Method = reqMethod
	if len(menuList) > 0 {
		sysOperaLog.Title = menuList[0].Title
	}
	b, _ := c.Get("body")
	sysOperaLog.OperParam, _ = tools.StructToJsonStr(b)
	sysOperaLog.CreateBy = tools.GetUserIdUint(c)
	sysOperaLog.OperTime = tools.GetCurrentTime()
	sysOperaLog.LatencyTime = (latencyTime).String()
	// UserAgent() 如果在请求中发送，UserAgent将返回客户端的用户代理
	sysOperaLog.UserAgent = c.Request.UserAgent()
	// 判断接口返回是否为空
	if c.Err() == nil {
		sysOperaLog.Status = "0"
	} else {
		sysOperaLog.Status = "1"
	}
	msgID := tools.GenerateMsgIDFromContext(c)
	db, err := tools.GetOrm(c)
	serviceOperaLog := service.SysOperaLog{}
	serviceOperaLog.MsgID = msgID
	serviceOperaLog.Orm = db
	_ = serviceOperaLog.InsertSysOperaLog(sysOperaLog.Generate())
}
