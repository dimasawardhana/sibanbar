package controllers

import (
	"net/http"
	"time"
	"../structs"
	"github.com/gin-gonic/gin"
)
type eventJSON struct{
	nama string `json:"nama"`
	tempat string `json:"tempat"`
	tanggal string `json:"tanggal"`
	tempat_lat float64 `json:"tempat_lat"`
	tempat_lang float64 `json:"tempat_lang"`
	poster string `json:"poster"`
}

func (idb *InDB) GetEvent(c *gin.Context){
	var (
		event []structs.Event
		result gin.H
	)
	idb.DB.Find(&event)

	if len(event)<0{
		result = gin.H{
			"status" : "Not Found",
			"count" : 0,
			"result" : gin.H{
				"event" : event,
			},
		}
	}else{
		result = gin.H{
			"status" : "OK",
			"count" : len(event),
			"result" : gin.H{
				"event" : event,
			},
		}
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) GetEventById(c *gin.Context){
	var (
		event structs.Event
		result gin.H
	)
	id := c.Param("id")

	err := idb.DB.Where("id = ?", id).First(&event).Error

	if err!= nil{
		result = gin.H{
			"status" : "Not Found",
			"error" : err,
		}
	}else{
		result = gin.H{
			"status" : "ok",
			"count" : 1,
			"result" : event,
		}
	}
	c.JSON(http.StatusOK, result)
}

func (idb * InDB) CreateEvent(c *gin.Context){
	var (
		result gin.H
		event structs.Event
		data eventJSON
	)
	if err:= c.ShouldBindJSON(&data); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
	}else{
	
		event.Nama = data.nama
		event.Tempat = data.tempat
		t, ers := time.Parse("2006-01-02", data.tanggal)
		event.Tanggal = t
		event.Tempat_lat = data.tempat_lat
		event.Tempat_lang = data.tempat_lang
		event.Poster = data.poster

		if ers != nil{
			result = gin.H{
				"status" : "failed",
				"result" : event,
				"error" : ers,
			}
		}else{
			idb.DB.Create(&event)
			result = gin.H{
				"status" : "ok",
				"result" : event,
			}
		}
		
		c.JSON(http.StatusOK, result)
	}
}

func (idb *InDB) UpdateEvent( c *gin.Context){
	var (
		result gin.H
		eventLama structs.Event
		eventBaru structs.Event
		data eventJSON
	)
	if err := c.ShouldBindJSON(&data); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}else{
		id := c.Param("id")
		err := idb.DB.First(&eventLama, id).Error
		if err != nil{
			result = gin.H{
				"status": "failed",
				"result" : "Data Not Found",
				"errors" : err,
			}
		}else{
			eventBaru.Nama = data.nama
			eventBaru.Tempat = data.tempat
			t, ers := time.Parse("2006-01-02", data.tanggal)
			eventBaru.Tanggal = t
			eventBaru.Tempat_lat = data.tempat_lat
			eventBaru.Tempat_lang = data.tempat_lang
			eventBaru.Poster = data.poster

			errs := idb.DB.Model(&eventLama).Updates(eventBaru).Error
			if ers != nil{
				result = gin.H{
					"status" : "failed",
					"result" : eventBaru,
					"error" : ers,
				}
			}else{
				if errs != nil{
					result = gin.H{
						"status" : "Failed",
						"result" : "failed update",
						"error" : errs,
					}
				}else{
					result  = gin.H{
						"status": "OK",
						"result" : eventBaru,
					}
				}
			}
			
			c.JSON(http.StatusOK, result)
		}

	}
}

func (idb *InDB) DeleteEvent(c *gin.Context){
	var (
		event structs.Event
		result gin.H
	)
	id := c.Param("id")

	err := idb.DB.First(&event, id).Error

	if err != nil{
		result = gin.H{
			"status" : "Not Found",
			"result" : "Data Not Found",
		}
	}else{
		errs := idb.DB.Delete(&event).Error
		if errs != nil{
			result = gin.H{
				"status": "Failed",
				"result" : "Failed to Delete Data",
			}
		}else{
			result = gin.H{
				"status" : "ok",
				"result" : "data with id :"+ id + " deleted successfully",

			}
		}
	}
	c.JSON(http.StatusOK, result)
}