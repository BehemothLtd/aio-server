package models

import "time"

type Issue struct {
	Id        int32
	ProjectId int32
	Project   Project
	// IssueType enums.IssueType
	ParentId    int32
	Title       string
	Description string
	Code        string
	// Priority enums.PriorityType
	IssueStatusId   int
	Position        int
	ProjectSprintId int
	StartDate       time.Time
	Deadline        time.Time
	Archived        int
	CreatorId       int
	Data            string
	IssueAssignees  []IssueAssignee
	LockVersion     int32
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
