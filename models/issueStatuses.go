package models

import (
	"aio-server/enums"
	"time"
)

type IssueStatus struct {
	Id         int32
	Color      string
	StatusType enums.IssueStatusStatusType
	Title      string

	CreatedAt   time.Time
	UpdatedAt   time.Time
	LockVersion int32 `gorm:"default:1"`
}
