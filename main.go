package main

import (
	"aio-server/controllers"
	"aio-server/database"
	"aio-server/models"
	"aio-server/pkg/auths"
	"aio-server/pkg/initializers"
	"aio-server/pkg/logger"
	"aio-server/tasks"
	"os"
	"regexp"

	"github.com/davecgh/go-spew/spew"
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

	// message := models.Message{Id: 1}
	// db.Table("messages").Find(&message)

	input := "This is a <@{213412341}> test string <@{sdfasfsfas}> with <@{xasdfas54433445yz}> placeholders."

	// Regular expression pattern to match "<@{%s}>"
	pattern := `<@{([a-zA-Z0-9]+)}>`

	// Compile the regular expression pattern
	re := regexp.MustCompile(pattern)

	// Find all matches in the input string
	matches := re.FindAllStringSubmatch(input, -1)

	var result []string
	if len(matches) > 0 {
		matches = matches[1:]
		for _, match := range matches {
			result = append(result, match[1])
		}
	}

	request := models.LeaveDayRequest{Id: 50}
	db.Table("leave_day_requests").Preload("messages").Find(&request).First(&request)

	spew.Dump(request)

	r.Run()
}
