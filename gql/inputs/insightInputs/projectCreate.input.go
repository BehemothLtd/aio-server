package insightInputs

import "time"

type ProjectCreateInput struct {
	Input ProjectCreateFormInput
}

type ProjectCreateFormInput struct {
	Name                 *string
	Code                 *string
	Description          *string
	ProjectType          *string
	SprintDuration       *int32
	ProjectIssueStatuses []ProjectIssueStatusInputForProjectCreate
	ProjectAssignees     *[]ProjectAssigneeInputForProjectCreate
}

type ProjectIssueStatusInputForProjectCreate struct {
	IssueStatusId int32
}

type ProjectAssigneeInputForProjectCreate struct {
	UserId            *int32
	DevelopmentRoleId *int
	Active            bool
	JoinDate          *time.Time
}
