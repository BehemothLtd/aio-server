package models

import (
	"time"
)

type LeaveDayRequest struct {
	Id         int32 `gorm:"not null;autoIncrement"`
	UserId     int32 `gorm:"not null;type:bigint;default:null"`
	ApproverId int32 `gorm:"not null;type:bigint;default:null"`
	From       time.Time
	To         time.Time
	TimeOff    float32
	// RequestType
	// RequestState
	Reason      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	LockVersion int32 `gorm:"not null;autoIncrement;default:0"`
}
