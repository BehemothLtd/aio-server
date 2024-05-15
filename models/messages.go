package models

import "time"

type Message struct {
	Id                int32 `gorm:"not null;autoIncrement"`
	ParentId          int32 `gorm:"not null;type:bigint;default:null"`
	LeaveDayRequestId int32 `gorm:"column:parent_id"`
	Content           *string
	Timestamp         *string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
