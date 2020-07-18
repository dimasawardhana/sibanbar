package main

import (
	"./config"
	"./controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}
	router := gin.Default()
	router.GET("/coba", inDB.DebugAPI)
	_orang := router.Group("/orang")
	{
		_orang.GET("/", inDB.GetOrang)
		_orang.GET("/:id", inDB.GetOrangById)
		_orang.DELETE(":id", inDB.DeleteOrang)
		_orang.POST("/:id", inDB.CreateOrang)
		_orang.PUT("/update/:id", inDB.UpdateOrang)
	}
	_group := router.Group("/group")
	{
		_group.GET("/", inDB.GetGroup)
		_group.GET("/:id", inDB.GetGroupByID)
		_group.GET("/type/:type", inDB.GetGroupByType)
		_group.GET("/parent/:id", inDB.GetGroupByParentId)
		_group.DELETE("/:id", inDB.DeleteGroup)
		_group.POST("/create", inDB.CreateGroup)
		_group.POST("/update/:id", inDB.UpdateGroup)
	}
	_instansi:= router.Group("/instansi")
	{
		_instansi.GET("/", inDB.GetInstansi)
		_instansi.GET("/:id", inDB.GetInstansiById)
		_instansi.DELETE("/:id", inDB.DeleteInstansi)
		_instansi.POST("/create", inDB.CreateInstansi)
		_instansi.POST("/update/:id", inDB.UpdateInstansi)
	}
	router.Run(":3000")
}
