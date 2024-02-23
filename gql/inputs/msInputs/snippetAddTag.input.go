package msInputs

import "github.com/graph-gophers/graphql-go"

// SnippetAddTagInput represents input for creating a snippet.
type SnippetAddTagInput struct {
	Id    graphql.ID
	TagId graphql.ID
}
