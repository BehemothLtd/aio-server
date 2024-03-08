package insightInputs

import "time"

type ProjectCreateInput struct {
	Input ProjectCreateFormInput
}

type ProjectCreateFormInput struct {
	Name                  *string
	Code                  *string
	Description           *string
	ProjectType           *string
	ProjectIssueStatusIds *[]int32
	ProjectAssignees      *[]ProjectAssigneeInputForProjectCreate
}

type ProjectAssigneeInputForProjectCreate struct {
	UserId            *int32
	DevelopmentRoleId *int
	Active            bool
	JoinDate          *time.Time
}
