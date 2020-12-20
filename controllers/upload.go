package controllers

import (
	"net/http"
	"os"
	"context"
	// "bytes"
	"crypto/sha1"
	"encoding/hex"
	// "encoding/json"
	"path/filepath"
	"io"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
)
func (idb *InDB) UploadImage(c *gin.Context){
	var hash string 
	maxSize := int64(1024000000000)
	err := c.Request.ParseMultipartForm(maxSize)
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err": err,
		})
	}
	file, err := fileHeader.Open()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"on":    "failed on file opening",
			"error": err.Error(),
		})
	}
	defer file.Close()
	checkHash := sha1.New()
	if _, err := io.Copy(checkHash, file); err != nil {
		hash = hex.EncodeToString(checkHash.Sum(nil)[:])
	}
	// put file into memory read

	pathFile := "/tmp/temporary"
	c.SaveUploadedFile(fileHeader, pathFile)
	newFile, err := os.Open(pathFile)
	defer newFile.Close()
	fileSize := fileHeader.Size

	// get file extension and mime type
	mime := fileHeader.Header.Get("Content-Type")
	ext := filepath.Ext(fileHeader.Filename)
	filename := hex.EncodeToString(checkHash.Sum(nil)[:]) + ext
	if hex.EncodeToString(checkHash.Sum(nil)[:]) == hex.EncodeToString(checkHash.Sum(nil)[:]) {
		uploadInfo, err := idb.Minio.PutObject(context.Background(),
			idb.MinioConf.Bucket,
			filename,
			newFile,
			fileSize,
			minio.PutObjectOptions{
				ContentType:        mime,
				ContentDisposition: "attachment",
			})
		
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"on":    "failed on upload",
				"error": err,
			})
		}else{
			src := "http://"+idb.MinioConf.Host+"/"+idb.MinioConf.Bucket+"/"+ filename
			c.JSON(http.StatusOK, gin.H{
				"src": src,
				// "bucket":     idb.MinioBUCKET,	
				// "file":       fileHeader,
				"uploadInfo": uploadInfo,
				// "checksum":   hash,
				// "filetype":   fileHeader.Header,
				// "mime":       mime,
				// "filename":   filename,
			})
		}
		

	} else {
		c.JSON(http.StatusOK, gin.H{
			"on": "failed on checksum checking",
			hex.EncodeToString(checkHash.Sum(nil)[:]): hash,
		})
	}
}