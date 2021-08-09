package global

import (
	"go-application/common/config"

	"gorm.io/gorm"
)

var Cfg config.Conf = config.DefaultConfig()

var Eloquent *gorm.DB

var (
	Source string
	Driver string
	DBName string
)
