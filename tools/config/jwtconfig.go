package config

import "github.com/spf13/viper"

type Jwt struct {
	TimeOut int64
	Secret  string
}

func InitJwt(cfg *viper.Viper) *Jwt {
	return &Jwt{
		TimeOut: cfg.GetInt64("timeout"),
		Secret:  cfg.GetString("secret"),
	}
}

var JwtConfig = new(Jwt)
