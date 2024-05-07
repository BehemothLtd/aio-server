package main

import (
	"aio-server/controllers"
	"aio-server/database"
	"aio-server/pkg/auths"
	"aio-server/pkg/initializers"
	"aio-server/pkg/logger"
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

	// var result []struct {
	// 		UserId            int32
	// 		User              models.User
	// 		TotalTimeApproved float64
	// 		TotalTimePending  float64
	// 		TotalTimeRejected float64
	// 	}
	// 	db.Raw(`
	// 		SELECT
	// 			user_id,
	// 			MAX(CASE WHEN request_state = 1 THEN total_time_off ELSE 0 END) AS total_time_pending,
	// 			MAX(CASE WHEN request_state = 2 THEN total_time_off ELSE 0 END) AS total_time_approved,
	// 			MAX(CASE WHEN request_state = 3 THEN total_time_off ELSE 0 END) AS total_time_rejected
	// 		FROM
	// 			(SELECT
	// 				leave_day_requests.user_id,
	// 				leave_day_requests.request_state,
	// 				SUM(time_off) as total_time_off
	// 			FROM
	// 				leave_day_requests
	// 			JOIN users ON
	// 				leave_day_requests.user_id = users.id
	// 			GROUP BY
	// 				leave_day_requests.user_id,
	// 				leave_day_requests.request_state) AS subquery
	// 		GROUP BY
	// 			user_id
	// 		ORDER BY
	// 			user_id;`).Scan(&result)

	r.Run()
}
