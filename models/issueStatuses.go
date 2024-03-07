package models

import (
	"aio-server/enums"
	"time"
)

type IssueStatus struct {
	Id         int32 `gorm:"not null;autoIncrement"`
	Color      string
	StatusType enums.IssueStatusStatusType
	Title      string

	CreatedAt time.Time
	UpdatedAt time.Time
}
