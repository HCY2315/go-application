package config

import "net/http"

type Conf interface {
	//单库业务实现这两个接口
	SetDb(db *DBConfig)
	GetDb() *DBConfig

	//使用的路由
	SetEngine(engine http.Handler)
	GetEngine() http.Handler
}
