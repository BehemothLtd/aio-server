package msInputs

import "github.com/graph-gophers/graphql-go"

type CollectionUpdateInput struct {
	Id    graphql.ID
	Input CollectionFormInput
}
