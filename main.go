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

	// time := time.Now()
	// project := models.Project{
	// 	Name:            "Project test",
	// 	Code:            "PTT",
	// 	ProjectType:     enums.ProjectTypeKanban,
	// 	ProjectPriority: enums.ProjectPriorityMedium,
	// 	State:           enums.ProjectStateActive,
	// 	ProjectAssignees: []models.ProjectAssignee{
	// 		{
	// 			Active:            true,
	// 			DevelopmentRoleId: 1,
	// 			UserId:            1,
	// 			JoinDate:          &time,
	// 		},
	// 		{
	// 			Active:            true,
	// 			DevelopmentRoleId: 1,
	// 			UserId:            2,
	// 			JoinDate:          &time,
	// 		},
	// 	},
	// 	ProjectIssueStatuses: []models.ProjectIssueStatus{
	// 		{
	// 			IssueStatusId: 1,
	// 			Position:      1,
	// 		},
	// 		{
	// 			IssueStatusId: 2,
	// 			Position:      2,
	// 		},
	// 	},
	// }
	// db.Table("projects").Create(&project)

	r.Run()
}
