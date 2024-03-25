package models

import (
	"aio-server/enums"
	"time"
)

type Issue struct {
	Id              int32
	ProjectId       int32
	Project         Project
	IssueType       enums.IssueType
	ParentId        int32
	Title           string
	Description     string
	Code            string
	Priority        enums.IssuePriority
	IssueStatusId   int
	Position        int
	ProjectSprintId *int32
	ProjectSprint   *ProjectSprint
	StartDate       time.Time
	Deadline        time.Time
	Archived        bool
	CreatorId       int32
	Creator         User
	Data            string
	IssueAssignees  []IssueAssignee
	Children        []Issue
	Parent          *Issue
	LockVersion     int32
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
