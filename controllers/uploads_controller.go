package controllers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"aio-server/pkg/helpers"
)

func UploadHandler(c *gin.Context) {
	uploader := createUploader(c)

	uploadedUrl, err := uploader.Upload()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"url": uploadedUrl,
	})
}

func createUploader(c *gin.Context) *helpers.Uploader {
	return &helpers.Uploader{
		Ctx:        c,
		Local:      os.Getenv("UPLOAD_LOCALLY_PATH") != "",
		UploadPath: "", // Update with your desired upload path
	}
}
