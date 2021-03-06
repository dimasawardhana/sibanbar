package controllers

import (
	"net/http"
	"time"
	"strings"
	"github.com/dimasawardhana/sibanbar/structs"
	"github.com/gin-gonic/gin"
)

type dataJSON struct {
	Nama_lengkap  string `json:"nama" form:"nama"`
	Alamat        string `json:"alamat" form:"alamat"`
	Tempat_lahir  string `json:"tempat_lahir" form:"tempat_lahir"`
	Tanggal_lahir string `json:"tanggal_lahir" form:"tanggal_lahir"`
	Status        string `json:"status" form:"status"`
	Photo         string `json:"photo" form:"photo"`
	KelompokID    uint   `json:"kelompokId" form:"kelompokId"`
	Phone         string `json:"phone" form:"phone"`
	Email         string `json:"email" form:"email"`
	JenisKelamin  string `json:"jenis_kelamin" form:"jenis_kelamin"`
}

func (idb *InDB) GetOrangById(c *gin.Context) {
	var (
		orang  structs.Orang
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).Preload("Kelompok").First(&orang).Error

	if err != nil {
		result = gin.H{
			"status": STATUS_ERROR,
			"result": err.Error(),			
		}
	} else {
		result = gin.H{
			"status": STATUS_OK,
			"result": orang,
		}
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) GetOrang(c *gin.Context) {
	var (
		orang2 []structs.Orang
		result gin.H
		data   dataJSON
		// orang  structs.Orang
	)
	var query = []string{}
	var values = []interface{}{}
	// id := c.Param("id")
	if err := c.BindQuery(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if data.Alamat != ""{
		query = append(query, "Alamat LIKE ?")
		values = append(values, data.Alamat)
		}
	if data.Nama_lengkap != ""{
		query = append(query, "Nama_lengkap LIKE ?")
		values = append(values, data.Nama_lengkap)
		}
	if data.Status != ""{
		query = append(query, "Status = ?")
		values = append(values, data.Status)
		}
	if data.KelompokID != 0{
		query = append(query, "KelompokID = ?")
		values = append(values, data.KelompokID)
		}
	// orang.Alamat = data.Alamat
	// orang.Nama_lengkap = data.Nama_lengkap
	// orang.Status = data.Status
	// orang.KelompokID = data.KelompokID
	idb.DB.Preload("Kelompok").
	// Where(&orang).
	Where(strings.Join(query, " AND "), values...).
	Find(&orang2)
	if len(orang2) <= 0 {
		result = gin.H{
			"status": STATUS_FAILED,
			"count":   0,
		}
	} else {
		result = gin.H{
			"status":  STATUS_OK,
			"result":  orang2,
			"count":   len(orang2),
		}
	}
	c.JSON(http.StatusOK, result)
}
func (idb *InDB) CreateOrang(c *gin.Context) {

	var (
		result gin.H
		orang  structs.Orang
		data   dataJSON
	)
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		orang.Nama_lengkap = data.Nama_lengkap
		orang.Alamat = data.Alamat
		orang.Tempat_lahir = data.Tempat_lahir
		t, errs := time.Parse("2006-01-02", data.Tanggal_lahir)
		orang.Tanggal_lahir = t
		orang.Status = data.Status
		orang.Photo = data.Photo
		err := idb.DB.Where("id = ?", data.KelompokID).First(&orang.Kelompok).Error
		orang.KelompokID = data.KelompokID
		orang.Phone = data.Phone
		orang.Email = data.Email
		orang.Jenis_kelamin = data.JenisKelamin
		if err != nil {
			result = gin.H{
				"status": STATUS_ERROR,
				"result": orang,
				"error":  err,
			}
			c.JSON(http.StatusBadRequest, result)
		}
		if errs != nil {
			result = gin.H{
				"status": STATUS_FAILED,
				"error":  errs,
			}
			c.JSON(http.StatusBadRequest, result)
		} else {
			idb.DB.Create(&orang)
			result = gin.H{
				"status": STATUS_OK,
				"result": orang,
			}
			c.JSON(http.StatusOK, result)
		}
		
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

	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {

		id := c.Param("id")
		errors := idb.DB.First(&orang, id).Error
		if errors != nil {
			result = gin.H{
				"status": STATUS_FAILED,
				"result": errors,
			}
			c.JSON(http.StatusOK, result)
		} else {
			lama = orang
			orang.Nama_lengkap = data.Nama_lengkap
			orang.Alamat = data.Alamat
			orang.Tempat_lahir = data.Tempat_lahir
			t, errs := time.Parse("2006-01-02", data.Tanggal_lahir)
			orang.Tanggal_lahir = t
			orang.Status = data.Status
			orang.Photo = data.Photo
			orang.Phone = data.Phone
			orang.Email = data.Email
			orang.Jenis_kelamin = data.JenisKelamin
			if errs != nil {
				result = gin.H{
					"status": STATUS_ERROR,
					"result": "failed on parsing time",
					"error":  errs,
				}
			}
			err := idb.DB.Where("id = ?", data.KelompokID).First(&orang.Kelompok).Error
			if err != nil {
				result = gin.H{
					"status": STATUS_FAILED,
					"result": orang,
					"error":  err,
				}
			}
			
			errs = idb.DB.Save(&orang).Error
			if errs != nil {
				result = gin.H{
					"status": STATUS_FAILED,
					"result": "update failed",
					"errs":   errs,
				}
			} else {
				result = gin.H{
					"status":        STATUS_OK,
					"beforeUpdated": lama,
					"result":        orang,
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
			"status": STATUS_FAILED,
			"result": "data not found",
		}
	}
	err = idb.DB.Delete(&orang).Error
	if err != nil {
		result = gin.H{
			"status": STATUS_FAILED,
			"result": "delete failed",
		}
	} else {
		result = gin.H{
			"status": STATUS_OK,
			"result": "data deleted successfully",
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

