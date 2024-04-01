package models

import "time"

type IssueAssignee struct {
	Id                int32
	IssueId           int32
	Issue             Issue
	UserId            int32
	User              User
	DevelopmentRoleId int32
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
