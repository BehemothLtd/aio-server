package insightInputs

import "github.com/graph-gophers/graphql-go"

type IssueStatusUpdateInput struct {
	Id    graphql.ID
	Input IssueStatusUpdateFormInput
}

type IssueStatusUpdateFormInput struct {
	Title       *string
	Color       *string
	StatusType  *string
	LockVersion *int32
}
