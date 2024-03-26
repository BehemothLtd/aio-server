package models

import (
	"aio-server/enums"
	"time"

	"gorm.io/gorm"
)

type LeaveDayRequest struct {
	Id           int32  `gorm:"not null;autoIncrement"`
	UserId       int32  `gorm:"not null;type:bigint;default:null"`
	ApproverId   *int32 `gorm:"not null;type:bigint;default:null"`
	User         User   `gorm:"foreignKey:UserId"`
	Approver     *User  `gorm:"foreignKey:ApproverId"`
	From         time.Time
	To           time.Time
	TimeOff      float64
	RequestType  enums.RequestType      `gorm:"not null;"`
	RequestState enums.RequestStateType `gorm:"not null;"`
	Reason       *string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	LockVersion  int32 `gorm:"not null;autoIncrement;default:0"`
}

func (r *LeaveDayRequest) BeforeUpdate(tx *gorm.DB) (err error) {
	if tx.Statement.Changed() {
		tx.Statement.SetColumn("lock_version", r.LockVersion+1)
	}

	return
}
