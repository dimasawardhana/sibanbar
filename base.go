package main

import (
	"github.com/dimasawardhana/sibanbar/config"
	"github.com/dimasawardhana/sibanbar/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}
	router := gin.Default()
	apiVersion := "/api/v1.0"
	router.GET("/coba", inDB.DebugAPI)
	_orang := router.Group(apiVersion + "/orang")
	{
		_orang.GET("/all/", inDB.GetOrang)
		_orang.GET("/byId/:id", inDB.GetOrangById)
		_orang.DELETE(":id", inDB.DeleteOrang)
		_orang.POST("/create", inDB.CreateOrang)
		_orang.PUT("/update/:id", inDB.UpdateOrang)
	}
	_group := router.Group(apiVersion + "/group")
	{
		_group.GET("/all/", inDB.GetGroup)
		_group.GET("/byId/:id", inDB.GetGroupByID)
		_group.GET("/byType/:type", inDB.GetGroupByType)
		_group.GET("/byParentId/:id", inDB.GetGroupByParentId)
		_group.DELETE("/:id", inDB.DeleteGroup)
		_group.POST("/create", inDB.CreateGroup)
		_group.PUT("/update/:id", inDB.UpdateGroup)
	}
	_instansi := router.Group(apiVersion + "/instansi")
	{
		_instansi.GET("/all/", inDB.GetInstansi)
		_instansi.GET("/byId/:id", inDB.GetInstansiById)
		_instansi.DELETE("/:id", inDB.DeleteInstansi)
		_instansi.POST("/create", inDB.CreateInstansi)
		_instansi.PUT("/update/:id", inDB.UpdateInstansi)
	}
	_user := router.Group(apiVersion + "/user")
	{
		_user.GET("/all/", inDB.GetUsers)
		_user.GET("/byId/:id", inDB.GetUserById)
		_user.DELETE("/:id", inDB.DeleteUser)
		_user.POST("/create", inDB.CreateUser)
		_user.PUT("/update/:id", inDB.UpdateUser)
	}
	_kesibukan := router.Group(apiVersion + "/kesibukan")
	{
		_kesibukan.GET("/all/", inDB.GetKesibukan)
		_kesibukan.GET("/byId/:id", inDB.GetKesibukanById)
		_kesibukan.DELETE("/:id", inDB.DeleteKesibukan)
		_kesibukan.POST("/create", inDB.CreateKesibukan)
		_kesibukan.PUT("/update/:id", inDB.UpdateKesibukan)
	}
	_relationship := router.Group(apiVersion + "/relationship")
	{
		_relationship.GET("/all/", inDB.GetRelationship)
		_relationship.GET("/byId/:id", inDB.GetRelationshipById)
		_relationship.DELETE("/:id", inDB.DeleteRelationship)
		_relationship.POST("/create", inDB.CreateRelationship)
		_relationship.PUT("/update/:id", inDB.UpdateRelationship)
	}
	_event := router.Group(apiVersion + "/event")
	{
		_event.GET("/all/", inDB.GetEvent)
		_event.GET("/byId/:id", inDB.GetEventById)
		_event.DELETE("/:id", inDB.DeleteEvent)
		_event.POST("/create", inDB.CreateEvent)
		_event.PUT("/update/:id", inDB.UpdateEvent)
	}
	_kehadiran := router.Group(apiVersion + "/kehadiran")
	{
		_kehadiran.GET("/all/", inDB.GetKehadiran)
		_kehadiran.GET("/byId/:id", inDB.GetKehadiranById)
		_kehadiran.DELETE("/:id", inDB.DeleteKehadiran)
		_kehadiran.POST("/create", inDB.CreateKehadiran)
		_kehadiran.PUT("/update/:id", inDB.UpdateKehadiran)
	}
	_dapuan := router.Group(apiVersion + "/dapuan")
	{
		_dapuan.GET("/all/", inDB.GetDapuan)
		_dapuan.GET("/byId/:id", inDB.GetDapuanById)
		_dapuan.DELETE("/:id", inDB.DeleteDapuan)
		_dapuan.POST("/create", inDB.CreateDapuan)
		_dapuan.PUT("/update/:id", inDB.UpdateDapuan)
	}
	router.Run(":3000")
}
