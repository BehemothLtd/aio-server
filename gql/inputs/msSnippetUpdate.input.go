package inputs

import "github.com/graph-gophers/graphql-go"

// MsSnippetUpdateInput represents input for creating a snippet.
type MsSnippetUpdateInput struct {
	Id    graphql.ID
	Input MsSnippetFormInput
}
