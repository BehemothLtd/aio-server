package models

import (
	"aio-server/enums"
	"time"
)

type LeaveDayRequest struct {
	Id           int32 `gorm:"not null;autoIncrement"`
	UserId       int32 `gorm:"not null;type:bigint;default:null"`
	ApproverId   int32 `gorm:"not null;type:bigint;default:null"`
	User         User
	Approver     User
	From         time.Time
	To           time.Time
	TimeOff      float32
	RequestType  enums.RequestType      `gorm:"not null;"`
	RequestState enums.RequestStateType `gorm:"not null;"`
	Reason       string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	LockVersion  int32 `gorm:"not null;autoIncrement;default:0"`
}
