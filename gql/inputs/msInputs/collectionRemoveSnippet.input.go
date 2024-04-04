package msInputs

import "github.com/graph-gophers/graphql-go"

type CollectionRemoveSnippetInput struct {
	Id        graphql.ID
	SnippetId graphql.ID
}
