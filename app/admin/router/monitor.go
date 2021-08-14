package router

import (
	"go-application/common/global"
	"go-application/tools"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Monitor() {
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
	}
	//开发环境启动监控指标
	r.GET("/metrics", Handler(promhttp.Handler()))
	r.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
}

func Handler(handler http.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}
