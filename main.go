package main

import (
	"aio-server/database"
	"aio-server/pkg/auths"
	"aio-server/pkg/helpers"
	"aio-server/pkg/initializers"
	"aio-server/pkg/logger"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	initializers.LoadEnv()

	// Load DB
	db := database.InitDb()

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Use(logger.Logger(logrus.New()), gin.Recovery())

	r.Use(auths.JwtTokenCheck, auths.GinContextToContextMiddleware()).POST("/graphql", initializers.GqlHandler(db))
	r.Use(auths.JwtTokenCheck, auths.GinContextToContextMiddleware()).POST("/uploads", uploadHandler)

	r.Run()
}

func uploadHandler(c *gin.Context) {
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
		c.JSON(http.StatusInternalServerError, gin.H{
			"url": uploadedUrl,
		})
	}
}
