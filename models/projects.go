package models

import (
	"aio-server/enums"
	"time"
)

type Project struct {
	Id                   int32
	Name                 string
	Code                 string
	ProjectType          enums.ProjectType
	ProjectPriority      enums.ProjectPriority
	State                enums.ProjectState
	ActivedAt            *time.Time
	InactivedAt          *time.Time
	StartedAt            *time.Time
	EndedAt              *time.Time
	CreatedAt            time.Time
	UpdatedAt            time.Time
	SprintDuration       int
	ClientId             int32
	CurrentSprintId      int
	ProjectAssignees     []ProjectAssignee
	ProjectIssueStatuses []ProjectIssueStatus
	IssueStatuses        []IssueStatus `gorm:"many2many:project_issue_statuses;"`
	LockVersion          int32         `gorm:"default:1"`
}
