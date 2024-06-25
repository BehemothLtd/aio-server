package main

import (
	"aio-server/controllers"
	"aio-server/database"
	"aio-server/pkg/auths"
	"aio-server/pkg/initializers"
	"aio-server/pkg/logger"
	"aio-server/pkg/mailer"
	"aio-server/tasks"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	os.Setenv("TZ", "Asia/Ho_Chi_Minh")
	initializers.LoadEnv()

	// Load DB
	db := database.InitDb()

	// Load Asynq
	tasks.InitAsyncClient()

	// Load Mailer
	mailer.InitDialer()

	// DEMO:
	// task, err := tasks.NewDemoTask(int32(1))
	// if err != nil {
	// 	log.Fatalf("could not create task: %v", err)
	// }
	// info, err := tasks.AsynqClient.Enqueue(task, asynq.ProcessIn(10*time.Second), asynq.Queue("critical"))
	// if err != nil {
	// 	log.Fatalf("could not enqueue task: %v", err)
	// }

	r := gin.Default()

	r.Use(initializers.CorsConfig())
	r.Use(logger.Logger(logrus.New()), gin.Recovery())

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})

	r.POST("/snippetGql", auths.JwtTokenCheck, auths.GinContextToContextMiddleware(), initializers.SnippetGqlHandler(db))
	r.POST("/insightGql", auths.JwtTokenCheck, auths.GinContextToContextMiddleware(), initializers.InsightGqlHandler(db))

	r.POST("/uploads", auths.JwtTokenCheck, auths.GinContextToContextMiddleware(), controllers.UploadHandler)
	r.POST("/slack/interactives", auths.GinContextToContextMiddleware(), controllers.Interactives)

	// DEMO SEND EMAIL:
	// registerEmailData := mailer.RegisterEmailData{
	// 	Name:    "bachdx",
	// 	Message: "Test message in HTML template",
	// }
	// go registerEmailData.Send("bachdx@behemoth.vn", "Test send Email in Go with HTML template")

	r.Run() // listen and serve on 0.0.0.0:8080
}
