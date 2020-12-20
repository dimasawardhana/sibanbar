
package controllers

import "github.com/jinzhu/gorm"
import "github.com/minio/minio-go/v7"
import "github.com/dimasawardhana/sibanbar/config"
type InDB struct {
  DB *gorm.DB 
  Minio *minio.Client
  MinioConf config.MinioConst
}