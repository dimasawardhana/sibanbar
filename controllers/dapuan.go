package controllers

import (
	"net/http"
	"../structs"
	"github.com/gin-gonic/gin"
)
type dapuanJSON struct{
	dapuan string `json:"nama"`
	orangId uint `json:"orangId"`
	grupId uint `json:"grupId"`
	keterangan string `json:"poster"`
}

func (idb *InDB) GetDapuan(c *gin.Context){
	var (
		dapuan []structs.Dapuan
		result gin.H
	)
	idb.DB.Find(&dapuan)

	if len(dapuan)<0{
		result = gin.H{
			"status" : "Not Found",
			"count" : 0,
			"result" : gin.H{
				"dapuan" : dapuan,
			},
		}
	}else{
		result = gin.H{
			"status" : "OK",
			"count" : len(dapuan),
			"result" : gin.H{
				"dapuan" : dapuan,
			},
		}
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) GetDapuanById(c *gin.Context){
	var (
		dapuan structs.Dapuan
		result gin.H
	)
	id := c.Param("id")

	err := idb.DB.Where("id = ?", id).First(&dapuan).Error

	if err!= nil{
		result = gin.H{
			"status" : "Not Found",
			"error" : err,
		}
	}else{
		result = gin.H{
			"status" : "ok",
			"count" : 1,
			"result" : dapuan,
		}
	}
	c.JSON(http.StatusOK, result)
}

func (idb * InDB) CreateDapuan(c *gin.Context){
	var (
		result gin.H
		dapuan structs.Dapuan
		data dapuanJSON
	)
	if err:= c.ShouldBindJSON(&data); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
	}else{
	
		dapuan.Dapuan = data.dapuan
		dapuan.OrangID = data.orangId
		dapuan.GrupID = data.grupId
		dapuan.Keterangan = data.keterangan

		
		idb.DB.Create(&dapuan)
		result = gin.H{
			"status" : "ok",
			"result" : dapuan,
		}
		
		c.JSON(http.StatusOK, result)
	}
}

func (idb *InDB) UpdateDapuan( c *gin.Context){
	var (
		result gin.H
		dapuanLama structs.Dapuan
		dapuanBaru structs.Dapuan
		data dapuanJSON
	)
	if err := c.ShouldBindJSON(&data); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}else{
		id := c.Param("id")
		err := idb.DB.First(&dapuanLama, id).Error
		if err != nil{
			result = gin.H{
				"status": "failed",
				"result" : "Data Not Found",
				"errors" : err,
			}
		}else{
			dapuanBaru.Dapuan = data.dapuan
			dapuanBaru.OrangID = data.orangId
			dapuanBaru.GrupID = data.grupId
			dapuanBaru.Keterangan = data.keterangan

			errs := idb.DB.Model(&dapuanLama).Updates(dapuanBaru).Error
			
			if errs != nil{
				result = gin.H{
					"status" : "Failed",
					"result" : "failed update",
					"error" : errs,
				}
			}else{
				result  = gin.H{
					"status": "OK",
					"result" : dapuanBaru,
				}
			}
		
			
			c.JSON(http.StatusOK, result)
		}

	}
}

func (idb *InDB) DeleteDapuan(c *gin.Context){
	var (
		dapuan structs.Dapuan
		result gin.H
	)
	id := c.Param("id")

	err := idb.DB.First(&dapuan, id).Error

	if err != nil{
		result = gin.H{
			"status" : "Not Found",
			"result" : "Data Not Found",
		}
	}else{
		errs := idb.DB.Delete(&dapuan).Error
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