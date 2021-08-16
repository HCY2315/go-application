package config

import (
	"database/sql"
	"net/http"
)

type Config struct {
	saas   bool
	db     *DBConfig
	dbs    map[string]*DBConfig
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

// SetEngine 设置引擎
func (c *Config) SetEngine(engine http.Handler) {
	c.engine = engine
}

// GetEngine 获取引擎
func (c *Config) GetEngine() http.Handler {
	return c.engine
}

// SetSaas 设置是否是saas应用
func (c *Config) SetSaas(saas bool) {
	c.saas = saas
}

// GetSaas 获取是否是saas应用
func (c *Config) GetSaas() bool {
	return c.saas
}

// SetDbs 设置对应key的db
func (c *Config) SetDbs(key string, db *DBConfig) {
	c.dbs[key] = db
}

// GetDbs 获取所有map里的db数据
func (c *Config) GetDbs() map[string]*DBConfig {
	return c.dbs
}

func DefaultConfig() *Config {
	return &Config{}
}
