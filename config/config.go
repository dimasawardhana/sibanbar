package config

import (
	"../structs"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "eizrael"+":"+"mysqlPengalaman354"+"@tcp("+"localhost"+":"+"3306"+")/"+"siibanbar"+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		// panic("failed to connect to database")
		panic(err)
	}

	db.AutoMigrate(structs.Orang{}, structs.Group{}, structs.Kesibukan{}, structs.Dapuan{}, structs.Relationship{}, structs.Event{}, structs.Instansi{}, structs.Kehadiran{}, structs.User{})
	return db
}
