package controllers

import (
	"net/http"

	"../structs"
	"github.com/gin-gonic/gin"
)

type dataGroup struct {
	Nama        string  `json:"nama" form:"nama"`
	Masjid      string  `json:"masjid" form:"masjid"`
	Masjid_lang float64 `json:"masjid_lang" form:"masjid_lang"`
	Masjid_lat  float64 `json:"masjid_lat" form:"masjid_lat"`
	Alamat      string  `json:"alamat" form:"alamat"`
	Group_type  string  `json:"group_type" form:"group_type"`
	ParentId    uint    `json:"parentId" form:"parentId"`
}

func (idb *InDB) GetGroup(c *gin.Context) {
	var (
		result gin.H
		grup   []structs.Group
	)

	idb.DB.Preload("Parent_group").Find(&grup)
	if len(grup) > 0 {
		result = gin.H{
			"status": "ok",
			"result": grup,
		}
	} else {
		result = gin.H{
			"status": "Not Found",
			"result": result,
		}
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) GetGroupByID(c *gin.Context) {
	var (
		grup structs.Group
		res  gin.H
	)
	id := c.Param("id")

	err := idb.DB.Where("id = ? ", id).Preload("Parent_group").First(&grup).Error

	if err != nil {
		res = notFound(gin.H{
			"err": err,
		})
	} else {
		res = gin.H{
			"status": "ok",
			"result": grup,
		}
	}
	c.JSON(http.StatusOK, res)
}
func (idb *InDB) GetGroupByType(c *gin.Context) {
	var (
		grup   []structs.Group
		result gin.H
	)
	group_type := c.Param("type")

	idb.DB.Where("Group_type = ?", group_type).Find(&grup)

	if len(grup) > 0 {
		result = gin.H{
			"status": "ok",
			"result": grup,
		}
	} else {
		result = gin.H{
			"status": "Not Found",
			"result": grup,
		}
	}
	c.JSON(http.StatusOK, result)
}
func (idb *InDB) GetGroupByParentId(c *gin.Context) {
	var (
		grup   []structs.Group
		result gin.H
	)
	id := c.Param("id")

	idb.DB.Where("ParentID = ?", id).Find(&grup)

	if len(grup) <= 0 {
		result = gin.H{
			"status": "ok",
			"result": grup,
		}
	} else {
		result = gin.H{
			"status": "Not Found",
			"result": result,
		}
	}
	c.JSON(http.StatusOK, result)
}
func (idb *InDB) CreateGroup(c *gin.Context) {
	var (
		res  gin.H
		data dataGroup
		grup structs.Group
	)

	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "bad request",
			"error":  err,
		})
	}
	grup.Nama = data.Nama
	grup.Masjid = data.Masjid
	grup.Masjid_lang = data.Masjid_lang
	grup.Masjid_lat = data.Masjid_lat
	grup.Alamat = data.Alamat
	grup.Group_type = data.Group_type
	grup.ParentID = data.ParentId

	errs := idb.DB.Create(&grup).Error

	if errs != nil {
		res = gin.H{
			"status": "failed",
			"result": errs,
		}
	} else {
		res = gin.H{
			"status": "OK",
			"result": grup,
		}
	}

	c.JSON(http.StatusOK, res)
}
func (idb *InDB) UpdateGroup(c *gin.Context) {
	var (
		grup structs.Group
		lama structs.Group
		data dataGroup
		res  gin.H
	)

	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "bad request",
		})
	}
	id := c.Param("id")

	errs := idb.DB.Where("id = ? ", id).First(&grup).Error

	if errs != nil {
		res = gin.H{
			"status": "Data Not Found",
			"error":  errs,
		}
	} else {
		lama = grup
		grup.Nama = data.Nama
		grup.Masjid = data.Masjid
		grup.Masjid_lang = data.Masjid_lang
		grup.Masjid_lat = data.Masjid_lat
		grup.Alamat = data.Alamat
		grup.Group_type = data.Group_type
		grup.ParentID = data.ParentId
		errss := idb.DB.Save(grup).Error
		if errss != nil {
			res = gin.H{
				"status": "UpdateFailed",
				"error":  errss,
			}
		} else {
			res = gin.H{
				"status": "ok",
				"result": "data updated",
				"before": lama,
				"data":   grup,
			}
		}

	}
	c.JSON(http.StatusOK, res)
}
func (idb *InDB) DeleteGroup(c *gin.Context) {
	var (
		grup structs.Group
		res  gin.H
	)
	id := c.Param("id")
	err := idb.DB.Where("id = ? ", id).First(&grup).Error

	if err != nil {
		res = notFound(gin.H{
			"err": err,
		})
	} else {
		err = idb.DB.Delete(&grup).Error
		if err != nil {
			res = gin.H{
				"status": "failed to delete",
				"result": "error",
				"error":  err,
			}
		} else {
			res = gin.H{
				"status": "ok",
				"result": "data deleted successfully",
				"error":  err,
			}
		}
	}
	c.JSON(http.StatusOK, res)
}
func notFound(dor gin.H) gin.H {
	return gin.H{
		"status": "Not Found",
		"result": dor,
	}
}
