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
		dapuan  []structs.Dapuan
		result  gin.H
		data    dapuanJSON
		toQuery structs.Dapuan
	)
	if err := c.BindQuery(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	toQuery.Dapuan = data.Dapuan
	toQuery.OrangID = data.OrangId
	toQuery.GrupID = data.GrupId
	idb.DB.
		Preload("Grup").
		Preload("Orang").
		Where(toQuery).
		Find(&dapuan)

	if len(dapuan) < 0 {
		result = gin.H{
			"status": "not found",
			"count":  0,
			"result": gin.H{
				"dapuan": dapuan,
			},
		}
	} else {
		result = gin.H{
			"status": "ok",
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
			"status": "failed",
			"result" : "not found",
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
		// dapuanLama structs.Dapuan
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
				"result": "not found",
				"errors": err,
			}
		} else {
			dapuanBaru.Dapuan = data.Dapuan
			dapuanBaru.OrangID = data.OrangId
			idb.DB.First(&dapuanBaru.Orang, dapuanBaru.OrangID)
			dapuanBaru.GrupID = data.GrupId
			idb.DB.First(&dapuanBaru.Grup, dapuanBaru.GrupID)

			// errs := idb.DB.Model(&dapuanLama).Updates(dapuanBaru).Error
			errs := idb.DB.Save(&dapuanBaru).Error
			if errs != nil {
				result = gin.H{
					"status": "failed",
					"result": "update failed",
					"error":  errs,
				}
			} else {
				result = gin.H{
					"status": "ok",
					// "before": dapuanLama,
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
			"status": "failed",
			"result": "data not found",
		}
	} else {
		errs := idb.DB.Delete(&dapuan).Error
		if errs != nil {
			result = gin.H{
				"status": "failed",
				"result": "failed to delete data",
			}
		} else {
			result = gin.H{
				"status": "ok",
				"result": "data deleted successfully",
			}
		}
	}
	c.JSON(http.StatusOK, result)
}
