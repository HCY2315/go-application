package middleware

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func WithCountextDb(dbMap map[string]*gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		if db, ok := dbMap["*"]; ok {
			c.Set("db", db)
		} else {
			// TODO: 配置dbs 使用
			// c.Set("db", c.Request.Host)
		}
		c.Next()
	}
}
