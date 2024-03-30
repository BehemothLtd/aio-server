package insightInputs

import "github.com/graph-gophers/graphql-go"

type SelfWorkingTimelogUpdateInput struct {
	Input   *SelfWorkingTimelogFormInput
	IssueId *int32
	Id      graphql.ID
}
