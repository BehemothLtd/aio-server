package models

import "time"

type ProjectAssignee struct {
	Id                int32 `gorm:"not null;autoIncrement"`
	Active            bool
	JoinDate          time.Time
	LeaveDate         time.Time
	LockVersion       int32 `gorm:"not null;autoIncrement;default:0"`
	DevelopmentRoleId int   `gorm:"not null;"`
	UserId            int32
	ProjectId         int32

	CreatedAt time.Time
	UpdatedAt time.Time
}
