package insightInputs

import graphql "github.com/graph-gophers/graphql-go"

type ProjectSprintInput struct {
	Id graphql.ID
}

type ProjectSprintCreateInput struct {
	Input ProjectSprintFormInput
}

type ProjectSprintFormInput struct {
	Title     *string
	ProjectId *int32
	StartDate *string
}
