package controllers

import (
	"github.com/gin-gonic/gin"
	"aio-server/pkg/helpers"
	"os"
	"net/http"
)

func UploadHandler(c *gin.Context) {
	uploader := helpers.Uploader{
		Ctx: c,
		Local: os.Getenv("UPLOAD_LOCALLY_PATH") != "",
	}

	uploadedUrl, err := uploader.Upload()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"url": uploadedUrl,
		})
	}
}
