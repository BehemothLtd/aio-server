package insightInputs

import graphql "github.com/graph-gophers/graphql-go"

type ClientInput struct {
	Id graphql.ID
}

type ClientCreateInput struct {
	Input ClientFormInput
}

type ClientUpdateInput struct {
	Id graphql.ID
	Input ClientFormInput
}

type ClientFormInput struct {
	Name           *string
	ShowOnHomePage *bool
	LockVersion    *int32
}
