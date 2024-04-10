package models

import (
	"aio-server/enums"
	"time"

	"gorm.io/gorm"
)

type Device struct {
	Id                    int32  `gorm:"not null;autoIncrement"`
	UserId                int32  `gorm:"not null;type:bigint;default:null"`
	User                  User   `gorm:"foreignKey:UserId"`
	Name                  string `gorm:"not null;"`
	Code                  string `gorm:"not null;"`
	Description           string
	State                 enums.DeviceStateType `gorm:"not null;"`
	DeviceTypeId          int32                 `gorm:"not null;"`
	Seller                string
	BuyAt                 time.Time
	CreatedAt             time.Time
	UpdatedAt             time.Time
	DeviceType            DeviceType `gorm:"foreignKey:DeviceTypeId"`
	LockVersion           int32      `gorm:"not null;autoIncrement;default:0"`
	DevicesUsingHistories []*DevicesUsingHistory
}

func (d *Device) BeforeUpdate(db *gorm.DB) (err error) {
	if db.Statement.Changed() {
		db.Statement.SetColumn("lock_version", d.LockVersion+1)
	}

	return
}
