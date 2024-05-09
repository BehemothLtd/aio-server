//go:generate go-enum --marshal
package enums

/*
ENUM(
all
read
write
delete
change_state
)
*/
type PermissionActionType string

/*
ENUM(
all
users
user_groups
projects
project_issues
project_issue_statuses
project_assignees
project_sprints
leave_day_requests
clients
issue_statuses
devices
timesheet_templates
attendances
project_experiences
)
*/
type PermissionTargetType string
