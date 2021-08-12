package cache

import (
	"go-application/common/log"
	"go-application/tools"
	toolsConfig "go-application/tools/config"
	"sync"
	"time"

	"github.com/gomodule/redigo/redis"
)

type Manager struct {
	pool    redis.Pool
	initOne sync.Once
}

var defaultManager *Manager

// CreateManager 创建Manager
func CreateManager() (*Manager, error) {
	pool := redis.Pool{
		MaxIdle:     16,
		MaxActive:   500,
		IdleTimeout: 240 * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", toolsConfig.RedisConfig.Addr)
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}
	_, err := pool.Dial()
	if err != nil {
		return nil, err
	}
	manager := &Manager{
		pool:    pool,
		initOne: sync.Once{},
	}
	return manager, nil
}

// SetUp 初始化Redis
func SetUp() {
	var err error
	if defaultManager == nil {
		defaultManager, err = CreateManager()
		if err != nil {
			log.Fatal(tools.Red("setup redis manager error:"), err)
			return
		}
	}
	log.Info(tools.Green("redis " + toolsConfig.RedisConfig.Addr + " connect success."))
}

func GetManager() *Manager {
	if defaultManager == nil {
		SetUp()
	}
	return defaultManager
}
