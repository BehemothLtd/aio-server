package msInputs

import "github.com/graph-gophers/graphql-go"

// SnippetRemoveTagInput represents input for creating a snippet.
type SnippetRemoveTagInput struct {
	Id    graphql.ID
	TagId graphql.ID
}
