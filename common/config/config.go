package config

import (
	"database/sql"
	"net/http"
)

type Config struct {
	db     *DBConfig
	engine http.Handler
}

type DBConfig struct {
	Driver string
	DB     *sql.DB
}

// SetDb 设置单个db
func (c *Config) SetDb(db *DBConfig) {
	c.db = db
}

// GetDb 获取单个db
func (c *Config) GetDb() *DBConfig {
	return c.db
}

// SetEnging 设置引擎
func (c *Config) SetEngine(enging http.Handler) {
	c.engine = enging
}

// GetEngine 获取引擎
func (c *Config) GetEngine() http.Handler {
	return c.engine
}

func DefaultConfig() *Config {
	return &Config{}
}
