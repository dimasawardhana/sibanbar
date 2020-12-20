package main

import (
	"github.com/dimasawardhana/sibanbar/config"
	"github.com/dimasawardhana/sibanbar/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	confConst := config.InitConst()
	db := config.DBInit(confConst.Mysql)
	minio := config.MinioInit(confConst.Minio)
	inDB := &controllers.InDB{DB: db, Minio: minio, MinioConf: confConst.Minio}
	router := gin.Default()
	apiVersion := "/api/v1"
	router.GET("/coba", inDB.DebugAPI)
	router.Use(config.CORSMiddleware())
	_upload := router.Group(apiVersion+"/upload")
	{
		_upload.PUT("", inDB.UploadImage )
	}
	_orang := router.Group(apiVersion + "/orang")
	{
		_orang.GET("", inDB.GetOrang)
		_orang.GET("/:id", inDB.GetOrangById)
		_orang.DELETE("/:id", inDB.DeleteOrang)
		_orang.POST("", inDB.CreateOrang)
		_orang.PUT("/:id", inDB.UpdateOrang)
	}
	_group := router.Group(apiVersion + "/group")
	{
		_group.GET("", inDB.GetGroup)
		_group.GET("/:id", inDB.GetGroupByID)		
		_group.DELETE("/:id", inDB.DeleteGroup)
		_group.POST("", inDB.CreateGroup)
		_group.PUT("/:id", inDB.UpdateGroup)
	}
	_instansi := router.Group(apiVersion + "/instansi")
	{
		_instansi.GET("", inDB.GetInstansi)
		_instansi.GET("/:id", inDB.GetInstansiById)
		_instansi.DELETE("/:id", inDB.DeleteInstansi)
		_instansi.POST("", inDB.CreateInstansi)
		_instansi.PUT("/:id", inDB.UpdateInstansi)
	}
	_user := router.Group(apiVersion + "/user")
	{
		_user.GET("", inDB.GetUsers)
		_user.GET("/:id", inDB.GetUserById)
		_user.DELETE("/:id", inDB.DeleteUser)
		_user.POST("", inDB.CreateUser)
		_user.PUT("/:id", inDB.UpdateUser)
	}
	_kesibukan := router.Group(apiVersion + "/kesibukan")
	{
		_kesibukan.GET("", inDB.GetKesibukan)
		_kesibukan.GET("/:id", inDB.GetKesibukanById)
		_kesibukan.DELETE("/:id", inDB.DeleteKesibukan)
		_kesibukan.POST("", inDB.CreateKesibukan)
		_kesibukan.PUT("/:id", inDB.UpdateKesibukan)
	}
	_relationship := router.Group(apiVersion + "/relationship")
	{
		_relationship.GET("", inDB.GetRelationship)
		_relationship.GET("/:id", inDB.GetRelationshipById)
		_relationship.DELETE("/:id", inDB.DeleteRelationship)
		_relationship.POST("", inDB.CreateRelationship)
		_relationship.PUT("/:id", inDB.UpdateRelationship)
	}
	_event := router.Group(apiVersion + "/event")
	{
		_event.GET("", inDB.GetEvent)
		_event.GET("/:id", inDB.GetEventById)
		_event.DELETE("/:id", inDB.DeleteEvent)
		_event.POST("", inDB.CreateEvent)
		_event.PUT("/:id", inDB.UpdateEvent)
	}
	_kehadiran := router.Group(apiVersion + "/kehadiran")
	{
		_kehadiran.GET("", inDB.GetKehadiran)
		_kehadiran.GET("/:id", inDB.GetKehadiranById)
		_kehadiran.DELETE("/:id", inDB.DeleteKehadiran)
		_kehadiran.POST("", inDB.CreateKehadiran)
		_kehadiran.PUT("/:id", inDB.UpdateKehadiran)
	}
	_dapuan := router.Group(apiVersion + "/dapuan")
	{
		_dapuan.GET("", inDB.GetDapuan)
		_dapuan.GET("/:id", inDB.GetDapuanById)
		_dapuan.DELETE("/:id", inDB.DeleteDapuan)
		_dapuan.POST("", inDB.CreateDapuan)
		_dapuan.PUT("/:id", inDB.UpdateDapuan)
	}
	router.Run(":3000")
}
