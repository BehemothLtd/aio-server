package main

import (
	"aio-server/controllers"
	"aio-server/database"
	"aio-server/models"
	"aio-server/pkg/auths"
	"aio-server/pkg/initializers"
	"aio-server/pkg/logger"
	"aio-server/tasks"
	"fmt"
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

	r.POST("/snippetGql", auths.JwtTokenCheck, auths.GinContextToContextMiddleware(), initializers.SnippetGqlHandler(db))
	r.POST("/insightGql", auths.JwtTokenCheck, auths.GinContextToContextMiddleware(), initializers.InsightGqlHandler(db))

	r.POST("/uploads", auths.JwtTokenCheck, auths.GinContextToContextMiddleware(), controllers.UploadHandler)
	r.POST("/slack/interactives", auths.GinContextToContextMiddleware(), controllers.Interactives)

	results := []models.RequestReport{}

	requestSubQuery := db.Table("leave_day_requests").Select(`leave_day_requests.user_id, leave_day_requests.request_state, SUM(time_off) as total_time_off`).
		Group("leave_day_requests.user_id, leave_day_requests.request_state")

	db.Table("(?) as Subquery", requestSubQuery).
		Select("user_id," +
			"SUM(CASE WHEN request_state = 1 THEN total_time_off ELSE 0 END) as approved_time, " +
			"SUM(CASE WHEN request_state = 2 THEN total_time_off ELSE 0 END) as pending_time, " +
			"SUM(CASE WHEN request_state = 3 THEN total_time_off ELSE 0 END) as rejected_time").
		Preload("User").
		Group("user_id").
		Order("user_id").
		Scan(&results)

	for _, result := range results {
		fmt.Printf("result : >>>>>>>>>> %+v \n\n", result)
	}

	fmt.Printf(">>>>>>>>>>> len results : \n %+v \n\n", len(results))

	r.Run()
}
