package config

import "github.com/spf13/viper"

type Redis struct {
	Addr string
}

func InitRedis(viper *viper.Viper) *Redis {
	redis := &Redis{
		Addr: viper.GetString("addr"),
	}
	return redis
}

var RedisConfig = new(Redis)
