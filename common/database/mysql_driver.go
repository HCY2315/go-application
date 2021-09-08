package database

import (
	"database/sql"
	"go-application/app/admin/models"
	"go-application/common/config"
	"go-application/common/global"
	"go-application/tools"
	toolsconfig "go-application/tools/config"
	"sync"

	"go-application/common/log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Mysql struct {
	initOne sync.Once
	models  []TableInterface
}

// Setup 初始化数据库
func (e *Mysql) Setup() {
	global.Source = e.GetConnect()

	db, err := sql.Open("mysql", global.Source)
	if err != nil {
		log.Error(tools.Red(e.GetDriver()+"connect error:"), err)
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
		log.Fatal((tools.Red(e.GetDriver() + " connect error:")), err)
	} else {
		log.Info(tools.Green(e.GetDriver() + " connect success "))
	}

	// 注册表
	e.RegisterTableModel()
	e.CheckTable()
}

// CheckTable 检查表结构
func (e *Mysql) CheckTable() {
	e.initOne.Do(func() {
		for _, md := range e.models {
			log.Info(tools.Green("insert: " + md.TableName()))
			global.Eloquent.AutoMigrate(md)
		}
	})
}

// RegisterTableModel 注册表结构
func (e *Mysql) RegisterTableModel() {
	e.models = append(e.models, &models.SysUser{})
	e.models = append(e.models, &models.OpClassRoom{})
	e.models = append(e.models, &models.SysRole{})
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
