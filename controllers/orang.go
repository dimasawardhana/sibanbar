package controllers

import (
	"net/http"
	"time"

	"../structs"
	"github.com/gin-gonic/gin"
)

type dataJSON struct {
	Nama_lengkap  string `json:"nama"`
	Alamat        string `json:"alamat" binding:"required"`
	Tempat_lahir  string `json:"tempat_lahir"`
	Tanggal_lahir string `json:"tanggal_lahir"`
	Status        string `json:"status"`
	Photo         string `json:"photo"`
	KelompokID    uint   `json:"kelompokId"`
}

func (idb *InDB) GetOrangById(c *gin.Context) {
	var (
		orang  structs.Orang
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&orang).Error

	if err != nil {
		result = gin.H{
			"status": "ERROR",
			"result": err.Error(),
			"count":  0,
		}
	} else {
		result = gin.H{
			"status": "OK",
			"result": orang,
			"count":  1,
		}
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) GetOrang(c *gin.Context) {
	var (
		orang2 []structs.Orang
		result gin.H
	)
	// id := c.Param("id")
	idb.DB.Find(&orang2)

	if len(orang2) <= 0 {
		result = gin.H{
			"status": "Not Found",
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"status": "OK",
			"result": orang2,
			"count":  len(orang2),
		}
	}
	c.JSON(http.StatusOK, result)
}
func (idb *InDB) CreateOrang(c *gin.Context) {

	var (
		orang  structs.Orang
		data   dataJSON
		result gin.H
	)
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		orang.Nama_lengkap = data.Nama_lengkap
		orang.Alamat = data.Alamat
		orang.Tempat_lahir = data.Tempat_lahir
		t, errs := time.Parse("2006-01-02", data.Tanggal_lahir)
		orang.Tanggal_lahir = t
		orang.Status = data.Status
		orang.Photo = data.Photo
		orang.KelompokID = data.KelompokID
		if errs != nil {
			result = gin.H{
				"status": "failed",
				"result": orang,
				"error":  errs,
			}
		} else {
			idb.DB.Create(&orang)
			result = gin.H{
				"status": "ok",
				"result": orang,
			}
		}
		c.JSON(http.StatusOK, result)
	}
}
func (idb *InDB) UpdateOrang(c *gin.Context) {
	var (
		lama   structs.Orang
		orang  structs.Orang
		data   dataJSON
		result gin.H
	)
	// cari data lama

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {

		id := c.Param("id")
		errors := idb.DB.First(&lama, id).Error
		if errors != nil {
			result = gin.H{
				"status": "failed",
				"result": errors,
			}
			c.JSON(http.StatusOK, result)
		} else {
			orang.Nama_lengkap = data.Nama_lengkap
			orang.Alamat = data.Alamat
			orang.Tempat_lahir = data.Tempat_lahir
			t, errs := time.Parse("2006-01-02", data.Tanggal_lahir)
			orang.Tanggal_lahir = t
			orang.Status = data.Status
			orang.Photo = data.Photo
			if errs != nil {
				result = gin.H{
					"status": "failed",
					"result": "failed on parsing time",
					"error":  errs,
				}
			}
			errs = idb.DB.Model(&lama).Updates(orang).Error
			if errs != nil {
				result = gin.H{
					"status": "failed",
					"result": "update failed",
				}
			} else {
				result = gin.H{
					"status": "ok",
					"result": orang,
				}
			}
			c.JSON(http.StatusOK, result)
		}

	}
}
func (idb *InDB) DeleteOrang(c *gin.Context) {
	var (
		orang  structs.Orang
		result gin.H
	)

	id := c.Param("id")
	err := idb.DB.First(&orang, id).Error
	if err != nil {
		result = gin.H{
			"status": "Not Found",
			"result": "Data Not Found",
		}
	}
	err = idb.DB.Delete(&orang).Error
	if err != nil {
		result = gin.H{
			"status": "Failed",
			"result": "Delete Failed",
		}
	} else {
		result = gin.H{
			"status": "OK",
			"result": "Data " + id + " deleted successfully",
		}
	}
	c.JSON(http.StatusOK, result)
}
func (idb *InDB) DebugAPI(c *gin.Context) {
	var (
		result gin.H
	)
	result = gin.H{
		"anjay": "mabar",
	}
	c.JSON(http.StatusOK, result)
}
