package insightInputs

import "github.com/graph-gophers/graphql-go"

type ProjectExperienceInput struct {
	Id graphql.ID
}

type ProjectExperienceFormInput struct {
	Title       string
	ProjectId   int32
	Description string
}

type ProjectExperienceCreateInput struct {
	Input ProjectExperienceFormInput
}

type ProjectExperienceUpdateInput struct {
	Id    graphql.ID
	Input ProjectExperienceFormInput
}
