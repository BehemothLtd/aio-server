package inputs

import "github.com/graph-gophers/graphql-go"

type MsSnippetUpdateInput struct {
	Id    graphql.ID
	Input MsSnippetFormInput
}
