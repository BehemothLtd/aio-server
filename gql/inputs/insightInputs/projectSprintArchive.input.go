package insightInputs

import "github.com/graph-gophers/graphql-go"

type ProjectSprintArchiveInput struct {
	ProjectId graphql.ID
	Id        graphql.ID
	MoveToId  graphql.ID
}
