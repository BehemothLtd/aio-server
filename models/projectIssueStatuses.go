package models

import "time"

type ProjectIssueStatus struct {
	Id            int32 `gorm:"not null;autoIncrement"`
	Position      int   `gorm:"not null;default: 1"`
	IssueStatusId int
	ProjectId     int
	IssueStatus   IssueStatus

	CreatedAt time.Time
	UpdatedAt time.Time
}
