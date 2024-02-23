package msInputs

import graphql "github.com/graph-gophers/graphql-go"

// TagDeleteInput represents input for deleting a tag
type TagDeleteInput struct {
	Id graphql.ID
}
