package models

import "time"

type RequestMessage struct {
	Id            int32 `gorm:"not null;autoIncrement"`
	UserId        int32 `gorm:"not null;type:bigint;default:null"`
	User          User  `gorm:"foreignKey:UserId"`
	ParentId      int32 `gorm:"not null;type:bigint;default:null"`
	UnixTimestamp *string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
