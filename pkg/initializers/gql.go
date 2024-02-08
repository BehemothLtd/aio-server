package initializers

import (
	gql "aio-server/gql/resolvers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"gorm.io/gorm"
)

func GqlHandler(db *gorm.DB) gin.HandlerFunc {
	s, err := getSchema()
	if err != nil {
		panic(err)
	}

	opts := []graphql.SchemaOpt{graphql.UseStringDescriptions(), graphql.UseFieldResolvers()}
	schema := graphql.MustParseSchema(s, &gql.Resolver{Db: db}, opts...)
	r := &relay.Handler{Schema: schema}

	return func(c *gin.Context) {
		r.ServeHTTP(c.Writer, c.Request)
	}
}

func getSchema() (string, error) {
	schemaPath := "./gql/schemas/"
	entries, err := os.ReadDir(schemaPath)
	if err != nil {
		log.Fatal(err)
	}

	schemaString := ""

	for _, e := range entries {
		filePath := schemaPath + e.Name()
		fileContent, _ := os.ReadFile(filePath)

		schemaString += string(fileContent)
	}

	return schemaString, nil
}
