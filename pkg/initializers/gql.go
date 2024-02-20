package initializers

import (
	"aio-server/gql/resolvers/insightResolvers"
	"aio-server/gql/resolvers/snippetResolvers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"gorm.io/gorm"
)

// SnippetGqlHandler returns a Gin middleware that handles GraphQL requests.
func SnippetGqlHandler(db *gorm.DB) gin.HandlerFunc {
	schema, err := fetchSchema("./gql/schemas/snippet/")

	if err != nil {
		log.Fatalf("failed to get schema: %v", err)
	}
	opts := []graphql.SchemaOpt{graphql.UseStringDescriptions(), graphql.UseFieldResolvers()}
	gqlSchema := graphql.MustParseSchema(schema, &snippetResolvers.Resolver{Db: db}, opts...)

	return ginSchemaHandler(schema, db, gqlSchema)
}

// InsightGqlHandler returns a Gin middleware that handles GraphQL requests.
func InsightGqlHandler(db *gorm.DB) gin.HandlerFunc {
	schema, err := fetchSchema("./gql/schemas/insight/")

	if err != nil {
		log.Fatalf("failed to get schema: %v", err)
	}

	opts := []graphql.SchemaOpt{graphql.UseStringDescriptions(), graphql.UseFieldResolvers()}
	gqlSchema := graphql.MustParseSchema(schema, &insightResolvers.Resolver{Db: db}, opts...)

	return ginSchemaHandler(schema, db, gqlSchema)
}

func ginSchemaHandler(schema string, db *gorm.DB, gqlSchema *graphql.Schema) gin.HandlerFunc {
	handler := &relay.Handler{Schema: gqlSchema}

	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}

func fetchSchema(schemaPath string) (string, error) {
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
