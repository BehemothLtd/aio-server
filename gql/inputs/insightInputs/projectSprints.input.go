package insightInputs

import graphql "github.com/graph-gophers/graphql-go"

type ProjectSprintInput struct {
	Id graphql.ID
}

type ProjectSprintCreateInput struct {
	ProjectId graphql.ID
	Input     ProjectSprintFormInput
}

type ProjectSprintUpdateInput struct {
	Id        graphql.ID
	ProjectId graphql.ID
	Input     ProjectSprintFormInput
}

type ProjectSprintFormInput struct {
	Title     *string
	StartDate *string
}
