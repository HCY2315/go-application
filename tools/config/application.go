package config

import "github.com/spf13/viper"

type Application struct {
	Model string
	Host  string
	Port  string
}

func InitApplication(cfg *viper.Viper) *Application {
	app := &Application{
		Model: cfg.GetString("model"),
		Host:  cfg.GetString("host"),
		Port:  cfg.GetString("port"),
	}
	return app
}

var ApplicationConfig = new(Application)
