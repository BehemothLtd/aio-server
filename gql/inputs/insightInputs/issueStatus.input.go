package insightInputs

import graphql "github.com/graph-gophers/graphql-go"

type IssueStatusInput struct {
	Id graphql.ID
}

type IssueStatusCreateInput struct {
	Input IssueStatusFormInput
}

type IssueStatusUpdateInput struct {
	Id    graphql.ID
	Input IssueStatusFormInput
}

type IssueStatusFormInput struct {
	Color       *string
	Title       *string
	StatusType  *string
	LockVersion *int32
}
