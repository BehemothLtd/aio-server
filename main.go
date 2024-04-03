package main

import (
	"aio-server/controllers"
	"aio-server/database"
	"aio-server/pkg/auths"
	"aio-server/pkg/initializers"
	"aio-server/pkg/logger"
	"aio-server/tasks"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
	"github.com/sirupsen/logrus"
)

func main() {
	initializers.LoadEnv()

	// Load DB
	db := database.InitDb()

	// Load Asynq
	tasks.InitAsyncClient()

	task, err := tasks.NewDemoTask(int32(1))
	if err != nil {
		log.Fatalf("could not create task: %v", err)
	}
	info, err := tasks.AsynqClient.Enqueue(task, asynq.ProcessIn(5*time.Second))
	if err != nil {
		log.Fatalf("could not enqueue task: %v", err)
	}
	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)

	r := gin.Default()

	r.Use(initializers.CorsConfig())
	r.Use(logger.Logger(logrus.New()), gin.Recovery())

	r.POST("/snippetGql", auths.JwtTokenCheck, auths.GinContextToContextMiddleware(), initializers.SnippetGqlHandler(db))
	r.POST("/insightGql", auths.JwtTokenCheck, auths.GinContextToContextMiddleware(), initializers.InsightGqlHandler(db))

	r.POST("/uploads", auths.JwtTokenCheck, auths.GinContextToContextMiddleware(), controllers.UploadHandler)

	r.Run()
}
