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

type WorkingTimelogHistory struct {
	Id               int32
	IssueName        string
	IssueDescription string
	IssueId          int32
	ProjectId        int32
	Minutes          int32
	LoggedAt         time.Time
}
