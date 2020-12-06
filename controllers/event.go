package controllers

import (
	"net/http"
	"time"

	"github.com/dimasawardhana/sibanbar/structs"
	"github.com/gin-gonic/gin"
)

type eventJSON struct {
	Nama        string  `json:"nama" form:"nama"`
	Tempat      string  `json:"tempat" form:"tempat"`
	Tanggal     string  `json:"tanggal" form:"tanggal"`
	Tempat_lat  float64 `json:"tempat_lat" form:"tempat_lat"`
	Tempat_lang float64 `json:"tempat_lang" form:"tempat_lang"`
	Poster      string  `json:"poster" form:"poster"`
}

func (idb *InDB) GetEvent(c *gin.Context) {
	var (
		event   []structs.Event
		result  gin.H
		data    eventJSON
		toQuery structs.Event
	)
	if err := c.BindQuery(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	toQuery.Nama = data.Nama
	toQuery.Tempat = data.Tempat
	idb.DB.
		Where(toQuery).
		Find(&event)

	if len(event) < 0 {
		result = gin.H{
			"status": "failed",
			"result" : "not found",
		}
	} else {
		result = gin.H{
			"status": "ok",
			"count":  len(event),
			"result": gin.H{
				"event": event,
			},
		}
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) GetEventById(c *gin.Context) {
	var (
		event  structs.Event
		result gin.H
	)
	id := c.Param("id")

	err := idb.DB.Where("id = ?", id).First(&event).Error

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
			"result": event,
		}
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) CreateEvent(c *gin.Context) {
	var (
		result gin.H
		event  structs.Event
		data   eventJSON
	)
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {

		event.Nama = data.Nama
		event.Tempat = data.Tempat
		t, ers := time.Parse("2006-01-02", data.Tanggal)
		event.Tanggal = t
		event.Tempat_lat = data.Tempat_lat
		event.Tempat_lang = data.Tempat_lang
		event.Poster = data.Poster

		if ers != nil {
			result = gin.H{
				"status": "failed",
				"result": "parsing data error",
				"error":  ers,
			}
		} else {
			idb.DB.Create(&event)
			result = gin.H{
				"status": "ok",
				"result": event,
			}
		}

		c.JSON(http.StatusOK, result)
	}
}

func (idb *InDB) UpdateEvent(c *gin.Context) {
	var (
		result    gin.H
		// eventLama structs.Event
		eventBaru structs.Event
		data      eventJSON
	)
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		id := c.Param("id")
		err := idb.DB.First(&eventBaru, id).Error
		if err != nil {
			result = gin.H{
				"status": "failed",
				"result": "data not found",
				"errors": err,
			}
		} else {
			// eventLama = eventBaru
			eventBaru.Nama = data.Nama
			eventBaru.Tempat = data.Tempat
			t, ers := time.Parse("2006-01-02", data.Tanggal)
			eventBaru.Tanggal = t
			eventBaru.Tempat_lat = data.Tempat_lat
			eventBaru.Tempat_lang = data.Tempat_lang
			eventBaru.Poster = data.Poster

			// errs := idb.DB.Model(&eventLama).Updates(eventBaru).Error
			errs := idb.DB.Save(&eventBaru).Error
			if ers != nil {
				result = gin.H{
					"status": "failed",
					"result": eventBaru,
					"error":  ers,
				}
			} else {
				if errs != nil {
					result = gin.H{
						"status": "failed",
						"result": "failed update",
						"error":  errs,
					}
				} else {
					result = gin.H{
						"status": "ok",
						// "before": eventLama,
						"result": eventBaru,
					}
				}
			}

			c.JSON(http.StatusOK, result)
		}

	}
}

func (idb *InDB) DeleteEvent(c *gin.Context) {
	var (
		event  structs.Event
		result gin.H
	)
	id := c.Param("id")

	err := idb.DB.First(&event, id).Error

	if err != nil {
		result = gin.H{
			"status": "failed",
			"result": "data not found",
		}
	} else {
		errs := idb.DB.Delete(&event).Error
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
