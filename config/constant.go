package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

type JwtConst struct {
	SecretKey string
	Algorithm string
}
type MinioConst struct {
	Host string
	Port string
	SecretKey string
	AccessKey string
	Bucket string
}
type MysqlConst struct{
	Userpass string
	Host string
	Port string
	DB string
}
type ConfigConst struct{
	Mysql MysqlConst
	Minio MinioConst
	Jwt JwtConst
}

func InitConst () ConfigConst{
	err := godotenv.Load()
	if err!= nil {
		log.Fatal("Error loading .env file")
	}
	return ConfigConst{
		Mysql : MysqlConst{
			Userpass: os.Getenv("MYSQL_USERNAME") +":"+ os.Getenv("MYSQL_PASSWORD"),
			Host: os.Getenv("MYSQL_HOSTNAME"),
			Port: os.Getenv("MYSQL_PORT"),
			DB: os.Getenv("MYSQL_DB"),
		},
		Minio : MinioConst{
			Host : os.Getenv("MINIO_HOST")+":"+os.Getenv("MINIO_PORT"),
			// Port : os.Getenv("MINIO_PORT"),
			SecretKey : os.Getenv("MINIO_SECRETKEY"),
			AccessKey : os.Getenv("MINIO_ACCESSKEY"),
			Bucket : os.Getenv("MINIO_BUCKET"),
		},
		Jwt : JwtConst{
			SecretKey : os.Getenv("JWT_SECRETKEY"),
			Algorithm : os.Getenv("JWT_ALGORITHM"),
		},
	}
}