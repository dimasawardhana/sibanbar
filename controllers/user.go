package controllers

import (
	"net/http"

	"github.com/dimasajiwardhana/sibanbar/structs"
	"github.com/gin-gonic/gin"
)

type orangJSON struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	OrangID  uint   `json:"orangId" form:"orangId"`
}

func (idb *InDB) GetUsers(c *gin.Context) {
	var (
		users  []structs.User
		result gin.H
	)

	idb.DB.Preload("Orang").Find(&users)

	if len(users) > 0 {
		result = gin.H{
			"status": "OK",
			"result": users,
			"count":  len(users),
		}
	} else {
		result = gin.H{
			"status": "Not Found",
			"result": nil,
			"count":  0,
		}
	}
	c.JSON(http.StatusOK, result)
}
func (idb *InDB) GetUserById(c *gin.Context) {

	var (
		user structs.User
		res  gin.H
	)
	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).Preload("Orang").First(&user).Error

	if err != nil {
		res = gin.H{
			"status": "Not Found",
			"error":  err,
		}
	} else {
		res = gin.H{
			"status": "ok",
			"result": user,
		}
	}
	c.JSON(http.StatusOK, res)
}
func (idb *InDB) CreateUser(c *gin.Context) {
	var (
		user     structs.User
		newOrang orangJSON
		res      gin.H
	)

	if err := c.ShouldBindJSON(&newOrang); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		user.Username = newOrang.Username
		user.Password = newOrang.Password
		user.OrangID = newOrang.OrangID

		idb.DB.Create(&user)
		res = gin.H{
			"status": "OK",
			"result": user,
		}
		c.JSON(http.StatusOK, res)
	}
}

func (idb *InDB) UpdateUser(c *gin.Context) {
	var (
		user     structs.User
		newOrang orangJSON
		oldOrang structs.User
		res      gin.H
	)
	if err := c.ShouldBindJSON(&newOrang); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		user.Username = newOrang.Username
		user.Password = newOrang.Password
		id := c.Param("id")
		errs := idb.DB.First(&oldOrang, id).Error
		if errs != nil {
			// not found
		} else {
			ers := idb.DB.Model(&oldOrang).Updates(user).Error
			if ers != nil {
				res = gin.H{
					"error":  ers,
					"result": "data is not updated",
				} // update failed
			} else {
				// update success
				res = gin.H{
					"status": "ok",
					"result": newOrang,
				}
			}
		}

		c.JSON(http.StatusOK, res)
	}

}
func (idb *InDB) DeleteUser(c *gin.Context) {
	var (
		user   structs.User
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.First(&user, id).Error

	if err != nil {
		result = gin.H{
			"status": "Not Found",
			"result": "Data Not Found",
		}
	} else {
		errs := idb.DB.Delete(&user).Error
		if errs != nil {
			result = gin.H{
				"status": "Failed",
				"result": "Failed to Delete Data",
			}
		} else {
			result = gin.H{
				"status": "OK",
				"result": "data id " + id + " deleted successfully",
			}
		}
	}

	c.JSON(http.StatusOK, result)
}

func NotFound() gin.H {
	return gin.H{
		"status": "Not Found",
		"result": "User Not Found",
	}
}
