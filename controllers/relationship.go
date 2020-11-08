package controllers

import (
	"net/http"

	"github.com/dimasawardhana/sibanbar/structs"
	"github.com/gin-gonic/gin"
)

type relationshipJSON struct {
	Relation_type string `json:"relation_type" form:"relation_type"`
	IdFrom        uint   `json:"idFrom" form:"idFrom"`
	IdTo          uint   `json:"idTo" form:"idTo"`
}

func (idb *InDB) GetRelationship(c *gin.Context) {
	var (
		relationship []structs.Relationship
		result       gin.H
		data         relationshipJSON
		toQuery      structs.Relationship
	)
	if err := c.BindQuery(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	toQuery.IdFrom = data.IdFrom
	toQuery.IdTo = data.IdFrom
	idb.DB.
		Preload("Orang").
		Where(toQuery).
		Find(&relationship)

	if len(relationship) < 0 {
		result = gin.H{
			"status": "Not Found",
			"count":  0,
			"result": gin.H{
				"relationship": relationship,
			},
		}
	} else {
		result = gin.H{
			"status": "OK",
			"count":  len(relationship),
			"result": gin.H{
				"relationship": relationship,
			},
		}
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) GetRelationshipById(c *gin.Context) {
	var (
		relationship structs.Relationship
		result       gin.H
	)
	id := c.Param("id")

	err := idb.DB.Where("id = ?", id).Preload("To").Preload("From").First(&relationship).Error

	if err != nil {
		result = gin.H{
			"status": "Not Found",
			"error":  err,
		}
	} else {
		result = gin.H{
			"status": "ok",
			"count":  1,
			"result": relationship,
		}
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) CreateRelationship(c *gin.Context) {
	var (
		result       gin.H
		relationship structs.Relationship
		data         relationshipJSON
	)
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		relationship.Relation_type = data.Relation_type
		relationship.IdFrom = data.IdFrom
		idb.DB.First(&relationship.From, data.IdFrom)
		relationship.IdTo = data.IdTo
		idb.DB.First(&relationship.To, data.IdTo)
		ers := idb.DB.Create(&relationship).Error
		if ers != nil {
			result = gin.H{
				"status": "failed",
				"result": relationship,
				"error":  ers,
			}
		} else {
			result = gin.H{
				"status": "ok",
				"result": relationship,
			}
		}

	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) UpdateRelationship(c *gin.Context) {
	var (
		result           gin.H
		relationshipLama structs.Relationship
		relationshipBaru structs.Relationship
		data             relationshipJSON
	)
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		id := c.Param("id")
		err := idb.DB.First(&relationshipBaru, id).Error
		if err != nil {
			result = gin.H{
				"status": "failed",
				"result": "Data Not Found",
				"errors": err,
			}
		} else {
			relationshipLama = relationshipBaru
			relationshipBaru.Relation_type = data.Relation_type
			relationshipBaru.IdFrom = data.IdFrom
			relationshipBaru.IdTo = data.IdTo
			idb.DB.First(&relationshipBaru.From, data.IdFrom)
			idb.DB.First(&relationshipBaru.To, data.IdTo)
			ers := idb.DB.Save(&relationshipBaru).Error
			if ers != nil {
				result = gin.H{
					"status": "failed",
					"result": relationshipBaru,
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
						"before": relationshipLama,
						"result": relationshipBaru,
					}
				}
			}

			c.JSON(http.StatusOK, result)
		}

	}
}

func (idb *InDB) DeleteRelationship(c *gin.Context) {
	var (
		relationship structs.Relationship
		result       gin.H
	)
	id := c.Param("id")

	err := idb.DB.First(&relationship, id).Error

	if err != nil {
		result = gin.H{
			"status": "Not Found",
			"result": "Data Not Found",
		}
	} else {
		errs := idb.DB.Delete(&relationship).Error
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
