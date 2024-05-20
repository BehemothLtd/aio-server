package models

import "time"

type Message struct {
	Id                int32 `gorm:"not null;autoIncrement"`
	LeaveDayRequestId int32 `gorm:"not null;type:bigint;default:null"`
	Content           *string
	Timestamp         *string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
