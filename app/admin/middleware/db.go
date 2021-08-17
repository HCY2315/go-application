package middleware

import (
	"go-application/common/config"
	"go-application/common/log"
	"go-application/common/middleware"
	"go-application/tools"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var WithCountextDb = middleware.WithCountextDb

func GetOrmFromConfig(cfg config.Conf) map[string]*gorm.DB {
	gormDBs := make(map[string]*gorm.DB)
	// TODO: 对类数据库连接

	// 生成单独的数据库连接
	c := cfg.GetDb()
	db, err := GetOrmFromDB(c.Driver, c.DB, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Error(tools.Red(c.Driver+" connect error :"), err)
	}
	gormDBs["*"] = db
	return gormDBs
}
