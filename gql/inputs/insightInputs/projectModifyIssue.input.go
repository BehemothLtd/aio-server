package insightInputs

import graphql "github.com/graph-gophers/graphql-go"

type ProjectCreateIssueInput struct {
	Id    graphql.ID
	Input ProjectModifyIssueFormInput
}

type ProjectModifyIssueFormInput struct {
	Title           *string
	Description     *string
	IssueStatus     *string
	IssueType       *string
	Priority        *string
	Archived        *bool
	Deadline        *string
	StartDate       *string
	IssueStatusId   *int32
	ParentId        *int32
	ProjectSprintId *int32
	IssueAssignees  *[]IssueAssigneeInputForIssueCreate
}

type IssueAssigneeInputForIssueCreate struct {
	UserId            *int32
	DevelopmentRoleId *int32
}
