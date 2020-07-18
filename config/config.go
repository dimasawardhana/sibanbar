package config

import (
	"../structs"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func DBInit() *gorm.DB {
	db, err := gorm.Open(DBmysql.client, DBmysql.user+":"+DBmysql.password+"@tcp("+DBmysql.hostname+":"+DBmysql.port+")/"+DBmysql.db+DBmysql.additionalSettings)
	if err != nil {
		// panic("failed to connect to database")
		panic(err)
	}

	db.AutoMigrate(structs.Orang{}, structs.Group{}, structs.Kesibukan{}, structs.Dapuan{}, structs.Relationship{}, structs.Event{}, structs.Instansi{}, structs.Kehadiran{}, structs.User{})
	return db
}
