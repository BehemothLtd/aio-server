package models

import (
	"aio-server/enums"
	"time"
)

type Project struct {
	Id                   int32 `gorm:"not null;autoIncrement"`
	Name                 string
	Code                 string
	ProjectType          enums.ProjectType
	ProjectPriority      enums.ProjectPriority
	State                enums.ProjectState
	ActivedAt            time.Time
	InactivedAt          time.Time
	StartedAt            time.Time
	EndedAt              time.Time
	CreatedAt            time.Time
	UpdatedAt            time.Time
	SprintDuration       int
	ClientId             int
	CurrentSprintId      int
	ProjectAssignees     []ProjectAssignee
	ProjectIssueStatuses []ProjectIssueStatus
}
