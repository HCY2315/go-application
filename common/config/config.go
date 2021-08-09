package config

import "database/sql"

type Config struct {
	db *DBConfig
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

func DefaultConfig() *Config {
	return &Config{}
}
