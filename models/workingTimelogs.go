package models

import (
	"time"

	"gorm.io/gorm"
)

type WorkingTimelog struct {
	Id          int32
	UserId      int32
	User        User
	ProjectId   int32
	Project     Project
	IssueId     int32
	Issue       Issue
	Minutes     int
	Description *string
	LoggedAt    time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	LockVersion int32 `gorm:"not null;default:0"`
}

func (r *WorkingTimelog) BeforeUpdate(tx *gorm.DB) (err error) {
	if tx.Statement.Changed() {
		tx.Statement.SetColumn("updated_at", time.Now())
		tx.Statement.SetColumn("lock_version", r.LockVersion+1)
	}

	return
}
