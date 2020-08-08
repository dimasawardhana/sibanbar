package config

import (
	"log"
	"os"

	"github.com/dimasawardhana/sibanbar/structs"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

func DBInit() *gorm.DB {
	errs := godotenv.Load()
	if errs != nil {
		log.Fatal("Error loading .env file")
	}
	// user := os.Getenv("MYSQL_USERNAME")
	// pass := os.Getenv("MYSQL_PASSWORD")
	userpass := os.Getenv("MYSQL_USERPASS")
	hostname := os.Getenv("MYSQL_HOSTNAME")
	port := os.Getenv("MYSQL_PORT")
	sqldb := os.Getenv("MYSQL_DB")

	settings := userpass + `@tcp(` + hostname + ":" + port + ")/" + sqldb + "?reconnect=true&charset=utf8&parseTime=True&loc=Local"
	// db, err := gorm.Open("mysql", "eizrael"+":"+"mysqlPengalaman354"+"@tcp("+"localhost"+":"+"3306"+")/"+"siibanbar"+"?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open("mysql", settings)
	if err != nil {
		// panic("failed to connect to database")
		panic(err)
	}

	db.AutoMigrate(structs.Orang{}, structs.Group{}, structs.Kesibukan{}, structs.Dapuan{}, structs.Relationship{}, structs.Event{}, structs.Instansi{}, structs.Kehadiran{}, structs.User{})
	return db
}
