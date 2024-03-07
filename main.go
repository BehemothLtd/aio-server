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

	r.POST("/snippetGql", auths.JwtTokenCheck, auths.GinContextToContextMiddleware(), initializers.SnippetGqlHandler(db))
	r.POST("/insightGql", auths.JwtTokenCheck, auths.GinContextToContextMiddleware(), initializers.InsightGqlHandler(db))

	r.POST("/uploads", auths.JwtTokenCheck, auths.GinContextToContextMiddleware(), controllers.UploadHandler)

	// project := models.Project{Id: 1}

	// db.Model(&project).Preload("ProjectAssignees").Preload("ProjectIssueStatuses").First(&project)

	// fmt.Printf("PROJECT %+v\n", project)

	r.Run()
}
