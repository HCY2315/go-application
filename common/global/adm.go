package global

import (
	"go-application/common/config"

	"github.com/casbin/casbin/v2"
	// "go-application/pkg/logger"
	"gorm.io/gorm"
)

var Cfg config.Conf = config.DefaultConfig()

var Eloquent *gorm.DB

var CasbinEnforcer *casbin.SyncedEnforcer

var (
	Source string
	Driver string
	DBName string
)

// TODO:日志系统
// var (
// 	Logger        = &logger.Logger{}
// 	JobLogger     = &logger.Logger{}
// 	RequestLogger = &logger.Logger{}
// )
