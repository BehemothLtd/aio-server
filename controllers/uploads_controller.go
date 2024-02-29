package controllers

import (
	"net/http"
	"os"

	"aio-server/pkg/helpers"

	"github.com/gin-gonic/gin"
)

func UploadHandler(c *gin.Context) {
	uploader := createUploader(c)

	uploaded, err := uploader.Upload()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": uploaded,
	})

}

func createUploader(c *gin.Context) *helpers.Uploader {
	return &helpers.Uploader{
		Ctx:        c,
		Local:      os.Getenv("UPLOAD_LOCALLY_PATH") != "",
		UploadPath: "", // Update with your desired upload path
	}
}
