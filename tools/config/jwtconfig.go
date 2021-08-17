package config

import "github.com/spf13/viper"

type Jwt struct {
	TimeOut int64
}

func InitJwt(cfg *viper.Viper) *Jwt {
	return &Jwt{
		TimeOut: cfg.GetInt64("timeout"),
	}
}

var JwtConfig = new(Jwt)
