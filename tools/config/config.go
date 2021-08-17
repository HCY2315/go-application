package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

// mysql 连接
var cfgDatabase *viper.Viper

// redis 连接
var cfgRedis *viper.Viper

// application 配置
var cfgApplication *viper.Viper

// jwtconfig 配置
var cfgJwt *viper.Viper

func Setup(path string) {
	// 显式定义配置文件的路径、名称和扩展名
	viper.SetConfigFile(path)

	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(fmt.Sprintf("Read config file fail: %s", err.Error()))
	}

	// ReadConfig：将读取配置文件，如果文件中不存在密钥，则将现有密钥设置为nil。
	// NewReader：返回从s读取的新读取器。它类似于bytes.NewBufferString，但效率更高且为只读。
	// ExpandEnv：根据当前环境变量的值替换字符串中的${var}或$var，对未定义变量的引用将替换为空字符串。

	// Replace environment variables 替换环境变量
	err = viper.ReadConfig(strings.NewReader(os.ExpandEnv(string(content))))
	if err != nil {
		log.Fatal(fmt.Sprintf("Parse config file fail: %s", err.Error()))
	}

	cfgDatabase = viper.Sub("settings.database")
	if cfgDatabase == nil {
		panic("No found settings.database in the configuration")
	}
	DatabaseConfig = InitDatabase(cfgDatabase)

	cfgRedis = viper.Sub("settings.redis")
	if cfgRedis == nil {
		panic("No found settings.redis in the configuration")
	}
	RedisConfig = InitRedis(cfgRedis)

	cfgApplication = viper.Sub("settings.application")
	if cfgApplication == nil {
		panic("No found settings.application in the configuration")
	}
	ApplicationConfig = InitApplication(cfgApplication)

	cfgJwt = viper.Sub("settings.jwt")
	if cfgJwt == nil {
		panic("No found settings.jwtconfig in the configuration")
	}
	JwtConfig = InitJwt(cfgJwt)
}
