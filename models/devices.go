package models

import (
	"aio-server/enums"
	"time"

	"gorm.io/gorm"
)

type Device struct {
	Id                    int32
	UserId                int32 `gorm:"default:null"`
	User                  User
	Name                  string
	Code                  string
	Description           string
	State                 enums.DeviceStateType
	DeviceTypeId          int32
	Seller                string
	BuyAt                 time.Time `gorm:"default:null"`
	CreatedAt             time.Time
	UpdatedAt             time.Time
	DeviceType            DeviceType
	LockVersion           int32
	DevicesUsingHistories []*DevicesUsingHistory
}

func (d *Device) BeforeUpdate(db *gorm.DB) (err error) {
	if db.Statement.Changed() {
		db.Statement.SetColumn("lock_version", d.LockVersion+1)
	}

	return
}
