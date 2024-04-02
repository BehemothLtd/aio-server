package models

import "time"

type WorkingTimelog struct {
	Id          int32
	UserId      int32
	User        User
	ProjectId   int32
	Project     Project
	IssueId     int32
	Issue       Issue
	Minutes     int
	Description string
	LoggedAt    time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
