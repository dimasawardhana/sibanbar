package controllers

import "github.com/gin-gonic/gin"
func responseSuccess(data gin.H) gin.H{
	return gin.H{
		"status": STATUS_OK,
		"result": data,
	}
}
func responseNotFound(message string) gin.H{
	return gin.H{
		"status": STATUS_OK,
		"result": message,
	}
}
func responseBadRequest(err error) gin.H{
	return gin.H{
		"status": STATUS_ERROR,
		"error": err,
	}
}
func responseFailed(err error, message string) gin.H{
	return gin.H{
		"status" : STATUS_FAILED,
		"result" : err,
		"message" : message,
	}
}