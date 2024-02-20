package msInputs

import "github.com/graph-gophers/graphql-go"

// SnippetUpdateInput represents input for creating a snippet.
type SnippetUpdateInput struct {
	Id    graphql.ID
	Input SnippetFormInput
}
