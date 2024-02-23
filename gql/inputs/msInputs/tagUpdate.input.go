package msInputs

import "github.com/graph-gophers/graphql-go"

// TagUpdateInput represents input for updating a tag
type TagUpdateInput struct {
	Id    graphql.ID
	Input TagFormInput
}
