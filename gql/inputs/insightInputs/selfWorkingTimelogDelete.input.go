package insightInputs

import "github.com/graph-gophers/graphql-go"

type SelfWorkingTimelogDeleteInput struct {
	IssueId *int32
	Id      graphql.ID
}
