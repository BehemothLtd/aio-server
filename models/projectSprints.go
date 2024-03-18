package models

import "time"

type ProjectSprint struct {
	Id          int32
	Title       string
	ProjectId   int32
	Project     Project
	StartDate   *time.Time
	EndDate     *time.Time
	Archived    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	LockVersion int32
}
