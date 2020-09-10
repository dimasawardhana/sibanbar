package controllers

import (
	"net/http"

	"github.com/dimasawardhana/sibanbar/structs"
	"github.com/gin-gonic/gin"
)

type dapuanJSON struct {
	Dapuan  string `json:"nama" form:"nama"`
	OrangId uint   `json:"orangId" form:"orangId"`
	GrupId  uint   `json:"grupId" form:"grupId"`
}

func (idb *InDB) GetDapuan(c *gin.Context) {
	var (
		dapuan []structs.Dapuan
		result gin.H
	)
	idb.DB.Preload("Grup").Preload("Orang").Find(&dapuan)

	if len(dapuan) < 0 {
		result = gin.H{
			"status": "Not Found",
			"count":  0,
			"result": gin.H{
				"dapuan": dapuan,
			},
		}
	} else {
		result = gin.H{
			"status": "OK",
			"count":  len(dapuan),
			"result": gin.H{
				"dapuan": dapuan,
			},
		}
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) GetDapuanById(c *gin.Context) {
	var (
		dapuan structs.Dapuan
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).Preload("Orang").Preload("Grup").First(&dapuan).Error

	if err != nil {
		result = gin.H{
			"status": "Not Found",
			"error":  err,
		}
	} else {
		result = gin.H{
			"status": "ok",
			"count":  1,
			"result": dapuan,
		}
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) CreateDapuan(c *gin.Context) {
	var (
		result gin.H
		dapuan structs.Dapuan
		data   dapuanJSON
	)
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {

		dapuan.Dapuan = data.Dapuan
		dapuan.OrangID = data.OrangId
		idb.DB.First(&dapuan.Orang, dapuan.OrangID)
		dapuan.GrupID = data.GrupId
		idb.DB.First(&dapuan.Grup, dapuan.GrupID)

		idb.DB.Create(&dapuan)
		result = gin.H{
			"status": "ok",
			"result": dapuan,
		}

		c.JSON(http.StatusOK, result)
	}
}

func (idb *InDB) UpdateDapuan(c *gin.Context) {
	var (
		result     gin.H
		dapuanLama structs.Dapuan
		dapuanBaru structs.Dapuan
		data       dapuanJSON
	)
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		id := c.Param("id")
		err := idb.DB.First(&dapuanBaru, id).Error
		if err != nil {
			result = gin.H{
				"status": "failed",
				"result": "Data Not Found",
				"errors": err,
			}
		} else {
			dapuanLama = dapuanBaru
			dapuanBaru.Dapuan = data.Dapuan
			dapuanBaru.OrangID = data.OrangId
			idb.DB.First(&dapuanBaru.Orang, dapuanBaru.OrangID)
			dapuanBaru.GrupID = data.GrupId
			idb.DB.First(&dapuanBaru.Grup, dapuanBaru.GrupID)

			// errs := idb.DB.Model(&dapuanLama).Updates(dapuanBaru).Error
			errs := idb.DB.Save(&dapuanBaru).Error
			if errs != nil {
				result = gin.H{
					"status": "Failed",
					"result": "failed update",
					"error":  errs,
				}
			} else {
				result = gin.H{
					"status": "OK",
					"before": dapuanLama,
					"result": dapuanBaru,
				}
			}

		}

	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) DeleteDapuan(c *gin.Context) {
	var (
		dapuan structs.Dapuan
		result gin.H
	)
	id := c.Param("id")

	err := idb.DB.First(&dapuan, id).Error

	if err != nil {
		result = gin.H{
			"status": "Not Found",
			"result": "Data Not Found",
		}
	} else {
		errs := idb.DB.Delete(&dapuan).Error
		if errs != nil {
			result = gin.H{
				"status": "Failed",
				"result": "Failed to Delete Data",
			}
		} else {
			result = gin.H{
				"status": "ok",
				"result": "data with id :" + id + " deleted successfully",
			}
		}
	}
	c.JSON(http.StatusOK, result)
}
