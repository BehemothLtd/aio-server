package models

import "time"

type Tag struct {
	Id          int32  `gorm:"not null;autoIncrement"`
	Name        string `gorm:"not null;"`
	UserId      int32  `gorm:"not null;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	LockVersion int32 `gorm:"not null;autoIncrement;default:0"`
}
