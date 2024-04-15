package msInputs

import "github.com/graph-gophers/graphql-go"

type CollectionAddSnippetInput struct {
	Id        graphql.ID
	SnippetId graphql.ID
}
