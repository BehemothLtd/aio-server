package insightInputs

import "github.com/graph-gophers/graphql-go"

type ProjectSprintIssueModifyInput struct {
	ProjectId graphql.ID
	Id        graphql.ID
	IssueId   graphql.ID
}
