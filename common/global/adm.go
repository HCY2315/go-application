package global

import (
	"go-application/common/config"

	"github.com/casbin/casbin/v2"
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
