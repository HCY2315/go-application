package config

import "github.com/spf13/viper"

type Logger struct {
	EnabledDB bool `default:"true"`
}

func InitLogger(cfg *viper.Viper) *Logger {
	return &Logger{
		EnabledDB: viper.GetBool("enabled_db"),
	}
}

var LoggerConfig = new(Logger)
