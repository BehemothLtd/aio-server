package models

import "time"

type Pin struct {
	Id         int32
	UserId     int32
	ParentType int
	ParentID   int32 `gorm:"column:parent_id"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
