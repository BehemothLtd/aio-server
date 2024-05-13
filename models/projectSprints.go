package models

import (
	"time"

	"gorm.io/gorm"
)

type ProjectSprint struct {
	Id          int32
	Title       string
	ProjectId   int32
	Project     Project
	StartDate   time.Time
	EndDate     *time.Time
	Archived    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	LockVersion int32
}

func (p *ProjectSprint) BeforeUpdate(tx *gorm.DB) (err error) {
	if tx.Statement.Changed() {
		tx.Statement.SetColumn("lock_version", p.LockVersion+1)
	}

	return
}
