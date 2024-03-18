package insightInputs

import "github.com/graph-gophers/graphql-go"

type ProjectUpdateInput struct {
	Id    graphql.ID
	Input ProjectUpdateFormInput
}

type ProjectUpdateFormInput struct {
	Name            *string
	ProjectPriority *string
	Description     *string
	ClientId        *int32
	State           *string
	ProjectType     *string
	SprintDuration  *int32
	StartedAt       *string
	EndedAt         *string
	LockVersion     *int32
}
