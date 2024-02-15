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

// GqlHandler returns a Gin middleware that handles GraphQL requests.
func GqlHandler(db *gorm.DB) gin.HandlerFunc {
	schema, err := getSchema()
	if err != nil {
		log.Fatalf("failed to get schema: %v", err)
	}

	opts := []graphql.SchemaOpt{graphql.UseStringDescriptions(), graphql.UseFieldResolvers()}
	gqlSchema := graphql.MustParseSchema(schema, &gql.Resolver{Db: db}, opts...)
	handler := &relay.Handler{Schema: gqlSchema}

	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}

func getSchema() (string, error) {
	schemaPath := "./gql/schemas/"
	entries, err := os.ReadDir(schemaPath)
	if err != nil {
		return "", err
	}

	var schemaContent []byte

	for _, entry := range entries {
		filePath := schemaPath + entry.Name()
		content, err := os.ReadFile(filePath)
		if err != nil {
			log.Printf("failed to read file %s: %v", filePath, err)
			continue
		}
		schemaContent = append(schemaContent, content...)
	}

	return string(schemaContent), nil
}
