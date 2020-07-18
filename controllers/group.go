package controllers

import (
	"net/http"
	"../structs"
	"github.com/gin-gonic/gin"
)


type dataGroup struct{
	Nama string `json:"nama"`
	Masjid string `json:"masjid"`
	Masjid_lang float64 `json:"masjid_lang"`
	Masjid_lat float64 `json:"masjid_lat"`
	Alamat string `json:"alamat"`
	Group_type int `json:"group_type"`
	ParentId uint `json:"parentId"`
}
func (idb *InDB) GetGroup( c *gin.Context){
	var (
		result gin.H
		grup []structs.Group
	)

	idb.DB.Find(&grup)
	if len(grup)<=0{
		result = gin.H{
			"status" : "ok",
			"result" : grup,
		}
	}	else{
		result = gin.H{
			"status" : "Not Found",
			"result" : result,
		}
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) GetGroupByID(c *gin.Context){
	var (
		grup structs.Group
		res gin.H
	)
	id := c.Param("id")

	err := idb.DB.First(&grup,id).Error

	if err!=nil{
		res = notFound()
	}else{
		res = gin.H{
			"status": "ok",
			"result" : res,
		}
	}
}
func (idb *InDB) GetGroupByType(c *gin.Context){
	var (
		grup []structs.Group
		result gin.H
	)
	group_type := c.Param("type")

	idb.DB.Where("Group_type = ?", group_type).Find(&grup)

	if len(grup)<=0{
		result = gin.H{
			"status" : "ok",
			"result" : grup,
		}
	}	else{
		result = gin.H{
			"status" : "Not Found",
			"result" : result,
		}
	}
	c.JSON(http.StatusOK, result)
}
func (idb *InDB) GetGroupByParentId(c *gin.Context){
	var (
		grup []structs.Group
		result gin.H
	)
	id := c.Param("id")

	idb.DB.Where("ParentID = ?", id).Find(&grup)

	if len(grup)<=0{
		result = gin.H{
			"status" : "ok",
			"result" : grup,
		}
	}	else{
		result = gin.H{
			"status" : "Not Found",
			"result" : result,
		}
	}
	c.JSON(http.StatusOK, result)
}
func (idb *InDB) CreateGroup(c *gin.Context){
	var (
		res gin.H
		data dataGroup
		grup structs.Group
	)

	if err:= c.ShouldBindJSON(&data); err !=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"status" : "bad request",
			"error": err,
		})
	}
	grup.Nama = data.Nama
	grup.Masjid = data.Masjid
	grup.Masjid_lang = data.Masjid_lang
	grup.Masjid_lat = data.Masjid_lat
	grup.Alamat = data.Alamat
	grup.Group_type = data.Group_type
	grup.ParentID = data.ParentId
	
	idb.DB.Create(&grup)
	res = gin.H{
		"status" : "OK",
		"result" : grup,
		
	}
	c.JSON(http.StatusOK, res)
}
func (idb *InDB) UpdateGroup(c *gin.Context){
	var (
		grup structs.Group
		lama structs.Group
		data dataGroup
		res gin.H
	)

	if err := c.ShouldBindJSON(&data); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "bad request",
		})
	}
	grup.Nama = data.Nama
	grup.Masjid = data.Masjid
	grup.Masjid_lang = data.Masjid_lang
	grup.Masjid_lat = data.Masjid_lat
	grup.Alamat = data.Alamat
	grup.Group_type = data.Group_type
	grup.ParentID = data.ParentId

	err := idb.DB.First(&lama).Error

	if err != nil{
		res = gin.H{
			"status" : "Data Not Found",
			"error" : err,
		}
	}else{
		err = idb.DB.Model(&lama).Updates(grup).Error
		if err != nil{
			res = gin.H{
				"status" : "Data Not Found",
				"error" : err,
			}	
		}else{
			res = gin.H{
				"status" : "ok",
				"result" : "data updated",
				"data" : grup,
			}
		}
		
	}
	c.JSON(http.StatusOK, res) 
}
func (idb *InDB) DeleteGroup(c *gin.Context){
	var (
		grup structs.Group
		res gin.H
	)
	id := c.Param("id")
	err := idb.DB.First(&grup,id).Error

	if err != nil{
		res = notFound();
	}else{

		err = idb.DB.Delete(&grup).Error
		if err != nil{
			res = gin.H{
				"status" : "failed to delete",
				"result" : "error",
				"error" : err,
			}
		}else{
			res = gin.H{
				"status" : "ok",
				"result" : "data deleted successfully",
				"error" : err,
			}
		}
	}
	c.JSON(http.StatusOK, res)
}
func notFound() gin.H{
	return gin.H{
		"status" : "Not Found",
		"result" : gin.H{},
	}
}