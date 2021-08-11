package database

import (
	"database/sql"
	"fmt"
	"go-application/common/config"
	"go-application/common/global"
	toolsconfig "go-application/tools/config"

	"go-application/common/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// mysql 结构体
type Mysql struct {
}

func (e *Mysql) Setup() {
	global.Source = e.GetConnect()
	// log.Info(tools.Green(global.Source))

	db, err := sql.Open("mysql", global.Source)
	if err != nil {
		log.Error(e.GetDriver(), "connect failed, err:", err)
	}
	global.Cfg.SetDb(&config.DBConfig{
		Driver: "mysql",
		DB:     db,
	})

	global.Eloquent, err = e.Open(db, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		log.Fatal((e.GetDriver() + " connect error :"), err)
	} else {
		log.Info(e.GetDriver() + " connect success !")
		fmt.Println("connect success!!!")
	}

	// TODO: 自动注册数据库中的表
}

// 打开数据库连接
func (e *Mysql) Open(db *sql.DB, cfg *gorm.Config) (*gorm.DB, error) {
	// Open: 基于方言器打开初始化数据库会话
	return gorm.Open(mysql.New(mysql.Config{Conn: db}), cfg)
}

// 获取数据库连接
func (e *Mysql) GetConnect() string {
	return toolsconfig.DatabaseConfig.Source
}

// 获取数据库类型
func (e *Mysql) GetDriver() string {
	return toolsconfig.DatabaseConfig.Driver
}
