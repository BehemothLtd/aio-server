package models

import "time"

type ProjectIssueStatuses struct {
	Id            int32 `gorm:"not null;autoIncrement"`
	Position      int   `gorm:"not null;default: 1"`
	IssueStatusId int
	ProjectId     int

	CreatedAt time.Time
	UpdatedAt time.Time
}
