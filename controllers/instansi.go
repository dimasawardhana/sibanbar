package controllers

import (
	"net/http"

	"github.com/dimasajiwardhana/sibanbar/structs"
	"github.com/gin-gonic/gin"
)

type dataInstansi struct {
	Nama   string `json:"nama" form:"nama" binding:required`
	Alamat string `json:"alamat" form:"alamat" binding:required`
}

func (idb *InDB) GetInstansi(c *gin.Context) {
	var (
		instansi []structs.Instansi
		result   gin.H
	)
	idb.DB.Find(&instansi)
	if len(instansi) > 0 {
		result = gin.H{
			"status": "Not Found",
			"count":  0,
			"result": gin.H{
				"instansi": instansi,
			},
		}
	} else {
		result = gin.H{
			"status": "OK",
			"count":  len(instansi),
			"result": gin.H{
				"instansi": instansi,
			},
		}
	}
	c.JSON(http.StatusOK, result)
}
func (idb *InDB) GetInstansiById(c *gin.Context) {
	var (
		instansi structs.Instansi
		result   gin.H
	)
	id := c.Param("id")
	err := idb.DB.First(&instansi, id).Error
	if err != nil {
		result = gin.H{
			"status": "Not Found",
			"error":  err,
		}
	} else {
		result = gin.H{
			"status": "ok",
			"count":  1,
			"result": instansi,
		}
	}
	c.JSON(http.StatusOK, result)
}
func (idb *InDB) CreateInstansi(c *gin.Context) {
	var (
		result   gin.H
		instansi structs.Instansi
		data     dataInstansi
	)
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		instansi.Nama_Instansi = data.Nama
		instansi.Alamat_Instansi = data.Alamat
		errs := idb.DB.Create(&instansi).Error
		if errs != nil {
			result = gin.H{
				"result": instansi,
				"error":  errs,
				"status": "failed",
			}
		} else {
			result = gin.H{
				"result": instansi,
				"check":  data,
				"status": "ok",
			}
		}
		c.JSON(http.StatusOK, result)
	}
}
func (idb *InDB) UpdateInstansi(c *gin.Context) {
	var (
		result   gin.H
		lama     structs.Instansi
		instansi structs.Instansi
		data     dataInstansi
	)
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		id := c.Param("id")
		err := idb.DB.First(&instansi, id).Error
		if err != nil {
			result = gin.H{
				"status": "failed",
				"result": err,
			}
			c.JSON(http.StatusOK, result)
		} else {

		}
		lama = instansi
		instansi.Nama_Instansi = data.Nama
		instansi.Alamat_Instansi = data.Alamat

		errs := idb.DB.Model(&lama).Updates(instansi).Error
		if errs != nil {
			result = gin.H{
				"status": "failed",
				"result": "update failed",
			}
		} else {
			result = gin.H{
				"result":       instansi,
				"beforeUpdate": lama,
				"status":       "ok",
			}
		}

		c.JSON(http.StatusOK, result)
	}
}
func (idb *InDB) DeleteInstansi(c *gin.Context) {
	var (
		instansi structs.Instansi
		result   gin.H
	)
	id := c.Param("id")
	err := idb.DB.First(&instansi, id).Error
	if err != nil {
		result = gin.H{
			"status": "Not Found",
			"result": "Data Not Found",
		}
	}
	ers := idb.DB.Delete(&instansi).Error
	if ers != nil {
		result = gin.H{
			"status": "Failed",
			"result": "Delete Failed",
		}
	} else {
		result = gin.H{
			"status": "OK",
			"result": "Data" + id + "deleted successfully",
		}
	}
	c.JSON(http.StatusOK, result)
}
