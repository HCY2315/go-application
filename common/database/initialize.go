package database

func Setup(driver string) {
	switch driver {
	case "mysql":
		var db = new(Mysql)
		db.Setup()
	default:
		var db = new(Mysql)
		db.Setup()
	}
}
