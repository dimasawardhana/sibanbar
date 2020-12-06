package config

import (
	"log"
	"os"

	"github.com/dimasawardhana/sibanbar/structs"
	"github.com/gin-gonic/gin"
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

	settings := userpass + `@tcp(` + hostname + ":" + port + ")/" + sqldb + "?charset=utf8&parseTime=True&loc=Local"
	// db, err := gorm.Open("mysql", "eizrael"+":"+"mysqlPengalaman354"+"@tcp("+"localhost"+":"+"3306"+")/"+"siibanbar"+"?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open("mysql", settings)
	if err != nil {
		// panic("failed to connect to database")
		panic(err)
	}

	db.AutoMigrate(structs.Orang{}, structs.Group{}, structs.Kesibukan{}, structs.Dapuan{}, structs.Relationship{}, structs.Event{}, structs.Instansi{}, structs.Kehadiran{}, structs.User{})
	return db
}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Accept, Origin, Cache-Control, X-Requested-With, access-control-allow-origin, access-control-allow-headers")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
