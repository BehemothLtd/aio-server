package models

import (
	"time"
)

type ProjectIssueStatus struct {
	Id            int32
	Position      int `gorm:"not null;default: 1"`
	IssueStatusId int32
	ProjectId     int32
	IssueStatus   IssueStatus

	CreatedAt time.Time
	UpdatedAt time.Time
}
