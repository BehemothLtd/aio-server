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

	// groupTitles := []string{"Super Admin", "BOD", "Team Leader", "Member", "test"}

	// for _, title := range groupTitles {
	// 	ug := models.UserGroup{Title: title}

	// 	if err := db.Table("user_groups").Where(&ug).First(&ug).Error; err == gorm.ErrRecordNotFound {
	// 		db.Table("user_groups").Create(&ug)
	// 	}

	// }
	// abc := systems.GetPermissions()

	// for _, p := range abc {
	// 	fmt.Print(p.Id)
	// }
	// seeds.InitPermissions(ug models.UserGroup, tx *gorm.DB)

	r.Run()
}
