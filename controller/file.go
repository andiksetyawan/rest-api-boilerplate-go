package controller

import (
	"github.com/andiksetyawan/rest-api-boilerplate-go/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DownloadFile(c *gin.Context) {
	bucketName := c.Param("bucket")
	fileID := c.Param("id")
	b,err:=storage.Download(bucketName,fileID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}
	c.Writer.Write(b)
}
