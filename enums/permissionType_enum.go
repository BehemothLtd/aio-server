// Code generated by go-enum DO NOT EDIT.
// Version: 0.6.0
// Revision: 919e61c0174b91303753ee3898569a01abb32c97
// Build Date: 2023-12-18T15:54:43Z
// Built By: goreleaser

package enums

import (
	"errors"
	"fmt"
)

const (
	// PermissionActionTypeAll is a PermissionActionType of type all.
	PermissionActionTypeAll PermissionActionType = "all"
	// PermissionActionTypeRead is a PermissionActionType of type read.
	PermissionActionTypeRead PermissionActionType = "read"
	// PermissionActionTypeWrite is a PermissionActionType of type write.
	PermissionActionTypeWrite PermissionActionType = "write"
	// PermissionActionTypeDelete is a PermissionActionType of type delete.
	PermissionActionTypeDelete PermissionActionType = "delete"
	// PermissionActionTypeChangeState is a PermissionActionType of type change_state.
	PermissionActionTypeChangeState PermissionActionType = "change_state"
)

var ErrInvalidPermissionActionType = errors.New("not a valid PermissionActionType")

// String implements the Stringer interface.
func (x PermissionActionType) String() string {
	return string(x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x PermissionActionType) IsValid() bool {
	_, err := ParsePermissionActionType(string(x))
	return err == nil
}

var _PermissionActionTypeValue = map[string]PermissionActionType{
	"all":          PermissionActionTypeAll,
	"read":         PermissionActionTypeRead,
	"write":        PermissionActionTypeWrite,
	"delete":       PermissionActionTypeDelete,
	"change_state": PermissionActionTypeChangeState,
}

// ParsePermissionActionType attempts to convert a string to a PermissionActionType.
func ParsePermissionActionType(name string) (PermissionActionType, error) {
	if x, ok := _PermissionActionTypeValue[name]; ok {
		return x, nil
	}
	return PermissionActionType(""), fmt.Errorf("%s is %w", name, ErrInvalidPermissionActionType)
}

// MarshalText implements the text marshaller method.
func (x PermissionActionType) MarshalText() ([]byte, error) {
	return []byte(string(x)), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *PermissionActionType) UnmarshalText(text []byte) error {
	tmp, err := ParsePermissionActionType(string(text))
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

const (
	// PermissionTargetTypeAll is a PermissionTargetType of type all.
	PermissionTargetTypeAll PermissionTargetType = "all"
	// PermissionTargetTypeUsers is a PermissionTargetType of type users.
	PermissionTargetTypeUsers PermissionTargetType = "users"
	// PermissionTargetTypeUserGroups is a PermissionTargetType of type user_groups.
	PermissionTargetTypeUserGroups PermissionTargetType = "user_groups"
	// PermissionTargetTypeProjects is a PermissionTargetType of type projects.
	PermissionTargetTypeProjects PermissionTargetType = "projects"
	// PermissionTargetTypeProjectIssues is a PermissionTargetType of type project_issues.
	PermissionTargetTypeProjectIssues PermissionTargetType = "project_issues"
	// PermissionTargetTypeLeaveDayRequests is a PermissionTargetType of type leave_day_requests.
	PermissionTargetTypeLeaveDayRequests PermissionTargetType = "leave_day_requests"
	// PermissionTargetTypeClients is a PermissionTargetType of type clients.
	PermissionTargetTypeClients PermissionTargetType = "clients"
	// PermissionTargetTypeIssueStatuses is a PermissionTargetType of type issue_statuses.
	PermissionTargetTypeIssueStatuses PermissionTargetType = "issue_statuses"
	// PermissionTargetTypeDevices is a PermissionTargetType of type devices.
	PermissionTargetTypeDevices PermissionTargetType = "devices"
	// PermissionTargetTypeTimesheetTemplates is a PermissionTargetType of type timesheet_templates.
	PermissionTargetTypeTimesheetTemplates PermissionTargetType = "timesheet_templates"
)

var ErrInvalidPermissionTargetType = errors.New("not a valid PermissionTargetType")

// String implements the Stringer interface.
func (x PermissionTargetType) String() string {
	return string(x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x PermissionTargetType) IsValid() bool {
	_, err := ParsePermissionTargetType(string(x))
	return err == nil
}

var _PermissionTargetTypeValue = map[string]PermissionTargetType{
	"all":                 PermissionTargetTypeAll,
	"users":               PermissionTargetTypeUsers,
	"user_groups":         PermissionTargetTypeUserGroups,
	"projects":            PermissionTargetTypeProjects,
	"project_issues":      PermissionTargetTypeProjectIssues,
	"leave_day_requests":  PermissionTargetTypeLeaveDayRequests,
	"clients":             PermissionTargetTypeClients,
	"issue_statuses":      PermissionTargetTypeIssueStatuses,
	"devices":             PermissionTargetTypeDevices,
	"timesheet_templates": PermissionTargetTypeTimesheetTemplates,
}

// ParsePermissionTargetType attempts to convert a string to a PermissionTargetType.
func ParsePermissionTargetType(name string) (PermissionTargetType, error) {
	if x, ok := _PermissionTargetTypeValue[name]; ok {
		return x, nil
	}
	return PermissionTargetType(""), fmt.Errorf("%s is %w", name, ErrInvalidPermissionTargetType)
}

// MarshalText implements the text marshaller method.
func (x PermissionTargetType) MarshalText() ([]byte, error) {
	return []byte(string(x)), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *PermissionTargetType) UnmarshalText(text []byte) error {
	tmp, err := ParsePermissionTargetType(string(text))
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}
