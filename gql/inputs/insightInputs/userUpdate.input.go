package insightInputs

import "github.com/graph-gophers/graphql-go"

type UserUpdateInput struct {
	Id    graphql.ID
	Input UserFormInput
}
