package insightInputs

import "github.com/graph-gophers/graphql-go"

type ProjectModifyProjectAssigneeInput struct {
	Id    graphql.ID
	Input ProjectModifyProjectAssigneeFormInput
}

type ProjectModifyProjectAssigneeFormInput struct {
	UserId            int32
	DevelopmentRoleId int32
	Active            bool
	JoinDate          string
	LeaveDate         *string
}
