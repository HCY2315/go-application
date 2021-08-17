package router

import (
	"go-application/app/admin/middleware"
	"go-application/common/global"
	"go-application/common/log"
	"go-application/tools"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	var r *gin.Engine
	h := global.Cfg.GetEngine()
	if h == nil {
		h = gin.New()
		global.Cfg.SetEngine(h)
	}
	switch h := h.(type) {
	case *gin.Engine:
		r = h
	default:
		log.Fatal(tools.Red("not support other engine"))
		os.Exit(-1)
	}

	// 跨域
	r.Use(cors())

	// 生成数据库连接
	r.Use(middleware.WithCountextDb(middleware.GetOrmFromConfig(global.Cfg)))

	// TODO：限流

	// TODO：初始化中间件
	var err error
	authMiddleware, err := middleware.AuthJwt()
	if err != nil {
		log.Error(err)
	}

	// TODO：注册系统路由

	// 注册业务路由
	InitBusinessRouter(r, authMiddleware)
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")

		// var headerKeys []string
		// for k, _ := range c.Request.Header {
		// 	headerKeys = append(headerKeys, k)
		// }
		// headerStr := strings.Join(headerKeys, ",")
		// if headerStr != "" {
		// 	headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		// } else {
		// 	headerStr = "access-control-allow-origin, access-control-allow-headers"
		// }
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*")
			// 服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			// header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			// 允许跨域设置,可以返回其他子段
			// 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")
			// 缓存请求信息 单位为秒
			c.Header("Access-Control-Max-Age", "172800")
			// 跨域请求是否需要带cookie信息 默认设置为true
			c.Header("Access-Control-Allow-Credentials", "false")
			// 设置返回格式是json
			c.Set("content-type", "application/json")
		}
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		c.Next()
	}
}
