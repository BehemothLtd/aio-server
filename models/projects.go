package models

import (
	"aio-server/enums"
	"time"
)

type Project struct {
	Id              int32 `gorm:"not null;autoIncrement"`
	Name            string
	Code            string
	ProjectType     enums.ProjectType
	ClientId        int32
	JiraId          int32
	SprintDuration  int32
	Description     string
	CurrentSprintId int32
	ProjectPriority enums.ProjectPriority
	// Setting         string  `Todo`
	State       enums.ProjectState
	ActiveAt    time.Time
	InactiveAt  time.Time
	StartedAt   time.Time
	EndedAt     time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	LockVersion int32 `gorm:"not null;autoIncrement;default:0"`
}

type ProjectQuery struct {
	Name        string
	Description string
	ProjectType enums.ProjectType
	State       enums.ProjectState
}

type ProjectCollection struct {
	Collection []*Project
	Metadata   *Metadata
}
