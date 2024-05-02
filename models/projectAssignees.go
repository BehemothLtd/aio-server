package models

import (
	"time"

	"gorm.io/gorm"
)

type ProjectAssignee struct {
	Id                int32
	Active            bool
	JoinDate          *time.Time
	LeaveDate         *time.Time
	LockVersion       int32 `gorm:"not null;default:1"`
	DevelopmentRoleId int32 `gorm:"not null;"`
	UserId            int32
	User              User
	ProjectId         int32
	Project           Project

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (pa *ProjectAssignee) BeforeUpdate(tx *gorm.DB) (err error) {
	if tx.Statement.Changed() {
		tx.Statement.SetColumn("lock_version", pa.LockVersion+1)
	}

	return
}
