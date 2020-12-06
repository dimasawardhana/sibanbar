package controllers

import (
	"net/http"

	"github.com/dimasawardhana/sibanbar/structs"
	"github.com/gin-gonic/gin"
)

type kesibukanJSON struct {
	Status     string `json:"status" form:"status"`
	Kedudukan  string `json:"kedudukan" form:"kedudukan"`
	InstansiId uint   `json:"instansiId" form:"instansiId"`
	OrangId    uint   `json:"orangId" form:"orangId"`
}

func (idb *InDB) GetKesibukan(c *gin.Context) {
	var (
		kesibukan []structs.Kesibukan
		result    gin.H
		data      kesibukanJSON
		toQuery   structs.Kesibukan
	)
	if err := c.BindQuery(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	toQuery.Status = data.Status
	toQuery.Kedudukan = data.Kedudukan
	toQuery.InstansiID = data.InstansiId
	toQuery.OrangID = data.OrangId
	idb.DB.
		Preload("instansi").
		Preload("Orang").
		Where(toQuery).
		Find(&kesibukan)

	if len(kesibukan) < 0 {
		result = gin.H{
			"status": "Not Found",
			"count":  0,
			"result": gin.H{
				"kesibukan": kesibukan,
			},
		}
	} else {
		result = gin.H{
			"status": "OK",
			"count":  len(kesibukan),
			"result": gin.H{
				"kesibukan": kesibukan,
			},
		}
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) GetKesibukanById(c *gin.Context) {
	var (
		kesibukan structs.Kesibukan
		result    gin.H
	)
	id := c.Param("id")

	err := idb.DB.Where("id = ?", id).Preload("instansi").Preload("Orang").First(&kesibukan).Error

	if err != nil {
		result = gin.H{
			"status": "Not Found",
			"error":  err,
		}
	} else {
		result = gin.H{
			"status": "ok",
			"count":  1,
			"result": kesibukan,
		}
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) CreateKesibukan(c *gin.Context) {
	var (
		result    gin.H
		kesibukan structs.Kesibukan
		data      kesibukanJSON
	)
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {

		kesibukan.Kedudukan = data.Kedudukan
		kesibukan.OrangID = data.OrangId
		kesibukan.Status = data.Status
		kesibukan.InstansiID = data.InstansiId
		idb.DB.First(&kesibukan.Orang, data.OrangId)
		idb.DB.First(&kesibukan.Instansi, data.InstansiId)
		ers := idb.DB.Create(&kesibukan).Error
		if ers != nil {
			result = gin.H{
				"status": "failed",
				"result": kesibukan,
				"error":  ers,
			}
		} else {

			result = gin.H{
				"status": "ok",
				"result": kesibukan,
			}
		}

		c.JSON(http.StatusOK, result)
	}
}

func (idb *InDB) UpdateKesibukan(c *gin.Context) {
	var (
		result        gin.H
		kesibukanLama structs.Kesibukan
		kesibukanBaru structs.Kesibukan
		data          kesibukanJSON
	)
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		id := c.Param("id")
		err := idb.DB.First(&kesibukanBaru, id).Error
		if err != nil {
			result = gin.H{
				"status": "failed",
				"result": "Data Not Found",
				"errors": err,
			}
		} else {
			kesibukanLama = kesibukanBaru
			kesibukanBaru.Kedudukan = data.Kedudukan
			kesibukanBaru.OrangID = data.OrangId
			idb.DB.First(&kesibukanBaru.Orang, data.OrangId)

			kesibukanBaru.Status = data.Status
			kesibukanBaru.InstansiID = data.InstansiId
			idb.DB.First(&kesibukanBaru.Instansi, data.InstansiId)

			ers := idb.DB.Save(&kesibukanBaru).Error
			if ers != nil {
				result = gin.H{
					"status": "failed",
					"result": kesibukanBaru,
					"error":  ers,
				}
			} else {
				if ers != nil {
					result = gin.H{
						"status": "Failed",
						"result": "failed update",
						"error":  ers,
					}
				} else {
					result = gin.H{
						"status": "OK",
						"before": kesibukanLama,
						"result": kesibukanBaru,
					}
				}
			}

			c.JSON(http.StatusOK, result)
		}

	}
}

func (idb *InDB) DeleteKesibukan(c *gin.Context) {
	var (
		kesibukan structs.Kesibukan
		result    gin.H
	)
	id := c.Param("id")

	err := idb.DB.First(&kesibukan, id).Error

	if err != nil {
		result = gin.H{
			"status": "Not Found",
			"result": "Data Not Found",
		}
	} else {
		errs := idb.DB.Delete(&kesibukan).Error
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
