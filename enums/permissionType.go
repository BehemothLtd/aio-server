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
projects
project_issues
)
*/
type PermissionTargetType string
