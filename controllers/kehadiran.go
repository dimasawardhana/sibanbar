package controllers

import (
	"net/http"
	"time"
	"../structs"
	"github.com/gin-gonic/gin"
)
type kehadiranJSON struct{
	jam_hadir string `json:"jam_hadir"`
	keterangan int `json:"keterangan"`
	eventId uint `json:"eventId"`
	orangId uint `json:"orangId"`
}

func (idb *InDB) GetKehadiran(c *gin.Context){
	var (
		kehadiran []structs.Kehadiran
		result gin.H
	)
	idb.DB.Find(&kehadiran)

	if len(kehadiran)<0{
		result = gin.H{
			"status" : "Not Found",
			"count" : 0,
			"result" : gin.H{
				"kehadiran" : kehadiran,
			},
		}
	}else{
		result = gin.H{
			"status" : "OK",
			"count" : len(kehadiran),
			"result" : gin.H{
				"kehadiran" : kehadiran,
			},
		}
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) GetKehadiranById(c *gin.Context){
	var (
		kehadiran structs.Kehadiran
		result gin.H
	)
	id := c.Param("id")

	err := idb.DB.Where("id = ?", id).First(&kehadiran).Error

	if err!= nil{
		result = gin.H{
			"status" : "Not Found",
			"error" : err,
		}
	}else{
		result = gin.H{
			"status" : "ok",
			"count" : 1,
			"result" : kehadiran,
		}
	}
	c.JSON(http.StatusOK, result)
}

func (idb * InDB) CreateKehadiran(c *gin.Context){
	var (
		result gin.H
		kehadiran structs.Kehadiran
		data kehadiranJSON
	)
	if err:= c.ShouldBindJSON(&data); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
	}else{
	
		kehadiran.Keterangan = data.keterangan
		kehadiran.OrangID = data.orangId
		t, ers := time.Parse("2006-01-02", data.jam_hadir)
		kehadiran.Jam_hadir = t
		kehadiran.EventID = data.eventId
		

		if ers != nil{
			result = gin.H{
				"status" : "failed",
				"result" : kehadiran,
				"error" : ers,
			}
		}else{
			idb.DB.Create(&kehadiran)
			result = gin.H{
				"status" : "ok",
				"result" : kehadiran,
			}
		}
		
		c.JSON(http.StatusOK, result)
	}
}

func (idb *InDB) UpdateKehadiran ( c *gin.Context){
	var (
		result gin.H
		kehadiranLama structs.Kehadiran
		kehadiranBaru structs.Kehadiran
		data kehadiranJSON
	)
	if err := c.ShouldBindJSON(&data); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}else{
		id := c.Param("id")
		err := idb.DB.First(&kehadiranLama, id).Error
		if err != nil{
			result = gin.H{
				"status": "failed",
				"result" : "Data Not Found",
				"errors" : err,
			}
		}else{
			kehadiranBaru.Keterangan = data.keterangan
			kehadiranBaru.OrangID = data.orangId
			t, ers := time.Parse("2006-01-02", data.jam_hadir)
			kehadiranBaru.Jam_hadir = t
			kehadiranBaru.EventID = data.eventId

			errs := idb.DB.Model(&kehadiranLama).Updates(kehadiranBaru).Error
			if ers != nil{
				result = gin.H{
					"status" : "failed",
					"result" : kehadiranBaru,
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
						"result" : kehadiranBaru,
					}
				}
			}
			
			c.JSON(http.StatusOK, result)
		}

	}
}

func (idb *InDB) DeleteKehadiran(c *gin.Context){
	var (
		kehadiran structs.Kehadiran
		result gin.H
	)
	id := c.Param("id")

	err := idb.DB.First(&kehadiran, id).Error

	if err != nil{
		result = gin.H{
			"status" : "Not Found",
			"result" : "Data Not Found",
		}
	}else{
		errs := idb.DB.Delete(&kehadiran).Error
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