package config

type Conf interface {
	//单库业务实现这两个接口
	SetDb(db *DBConfig)
	GetDb() *DBConfig
}
