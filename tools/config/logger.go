package config

import "github.com/spf13/viper"

type Logger struct {
	Enableddb string
}

func InitLogger(cfg *viper.Viper) *Logger {
	return &Logger{
		Enableddb: viper.GetString("enableddb"),
	}
}

var LoggerConfig = new(Logger)
