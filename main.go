package main

import (
	"aio-server/database"
	"aio-server/pkg/auths"
	"aio-server/pkg/initializers"
	"aio-server/pkg/logger"
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

	r.Run()
}
