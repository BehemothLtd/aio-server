package main

import (
	"aio-server/database"
	"aio-server/gql"
	"aio-server/pkg/logger"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func main() {
	// Load ENV
	err := godotenv.Load()

	if err != nil {
		panic("Cant Load .env file")
	}

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

	log := logrus.New()
	logger := logger.Logger(log)

	r.Use(logger, gin.Recovery())

	r.POST("/graphql", gophersGraphQLHandler(db))

	r.Run()
}

func gophersGraphQLHandler(db *gorm.DB) gin.HandlerFunc {
	s, err := getSchema("./gql/schema.graphql")
	if err != nil {
		panic(err)
	}

	schema := graphql.MustParseSchema(s, &gql.Resolver{Db: db}, graphql.UseStringDescriptions())
	r := &relay.Handler{Schema: schema}

	return func(c *gin.Context) {
		r.ServeHTTP(c.Writer, c.Request)
	}
}

func getSchema(path string) (string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
