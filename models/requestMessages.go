package models

import "time"

type RequestMessage struct {
	Id        int32 `gorm:"not null;autoIncrement"`
	ParentId  int32 `gorm:"not null;type:bigint;default:null"`
	Content   *string
	Timestamp *string
	CreatedAt time.Time
	UpdatedAt time.Time
}
