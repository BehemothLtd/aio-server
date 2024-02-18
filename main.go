package main

import (
	"aio-server/controllers"
	"aio-server/database"
	"aio-server/pkg/auths"
	"aio-server/pkg/initializers"
	"aio-server/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	initializers.LoadEnv()

	// Load DB
	db := database.InitDb()

	r := gin.Default()

	r.Use(initializers.CorsConfig())
	r.Use(logger.Logger(logrus.New()), gin.Recovery())

	r.Use(auths.JwtTokenCheck, auths.GinContextToContextMiddleware()).POST("/graphql", initializers.GqlHandler(db))
	r.Use(auths.JwtTokenCheck, auths.GinContextToContextMiddleware()).POST("/uploads", controllers.UploadHandler)

	r.Run()
}
