package models

import (
	"aio-server/enums"
	"time"

	"gorm.io/gorm"
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
	IssueStatusId   int32
	Position        int
	ProjectSprintId *int32
	ProjectSprint   *ProjectSprint
	StartDate       time.Time
	Deadline        time.Time
	Archived        bool
	CreatorId       int32
	Creator         User `gorm:"foreignkey:CreatorId"`
	Data            *string
	IssueAssignees  []IssueAssignee
	Children        []Issue `gorm:"foreignkey:ParentId"`
	Parent          *Issue
	LockVersion     int32
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func (i *Issue) BeforeCreate(tx *gorm.DB) (err error) {
	i.SetPosition(tx)

	return
}

func (i *Issue) SetPosition(tx *gorm.DB) *Issue {
	maxPosition := 0
	tx.Model(&i).Select("MAX(position)").Where("issue_status_id = ?", i.IssueStatusId).Where("project_id = ?", i.ProjectId).Scan(&maxPosition)

	i.Position = maxPosition + 1
	return i
}
