package database

import (
	"database/sql"
	"go-application/common/global"
	toolsconfig "go-application/tools/config"
)

// mysql 结构体
type Mysql struct {
}

func (m *Mysql) Setup() {
	global.Source = m.GetConnect()
	// log.Info(tools.Green(global.Source))

	db, err := sql.Open("mysql", global.Source)
	if err != nil {

	}

}

// 获取数据库连接
func (e *Mysql) GetConnect() string {
	return toolsconfig.DatabaseConfig.Source
}
