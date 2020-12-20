package config

import (
	"log"

	"github.com/dimasawardhana/sibanbar/structs"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func DBInit(Mysql MysqlConst) *gorm.DB {
	userpass := Mysql.Userpass
	hostname := Mysql.Host
	port := Mysql.Port
	sqldb := Mysql.DB
	settings := userpass + `@tcp(` + hostname + ":" + port + ")/" + sqldb + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", settings)
	if err != nil {
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
func MinioInit(minioConst MinioConst) *minio.Client{
	useSSL := false
	minioClient, err := minio.New(minioConst.Host, &minio.Options{
		Creds:  credentials.NewStaticV4(minioConst.AccessKey, minioConst.SecretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatal(err, "shit")
	}
	return minioClient
}