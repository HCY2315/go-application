package database

import (
	"database/sql"
	"fmt"
	"go-application/common/config"
	"go-application/common/global"
	toolsconfig "go-application/tools/config"

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
		//
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
	fmt.Println("bbbb")
	if err != nil {
		// log.Fatal(tools.Red(e.GetDriver()+" connect error :"), err)
	} else {
		// log.Info(tools.Green(e.GetDriver() + " connect success !"))
		fmt.Println("connect success!!!")
	}
}

// 打开数据库连接
func (e *Mysql) Open(db *sql.DB, cfg *gorm.Config) (*gorm.DB, error) {
	return gorm.Open(mysql.New(mysql.Config{Conn: db}), cfg)
}

// 获取数据库连接
func (e *Mysql) GetConnect() string {
	return toolsconfig.DatabaseConfig.Source
}
