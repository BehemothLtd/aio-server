package models

import (
	"aio-server/enums"
	"time"

	"gorm.io/gorm"
)

type DevicesUsingHistory struct {
	Id          int32
	UserId      int32 `gorm:"default:null"`
	User        User
	DeviceId    int32
	Device      Device
	State       enums.DeviceStateType
	LockVersion int32
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (d *DevicesUsingHistory) BeforeUpdate(db *gorm.DB) (err error) {
	if db.Statement.Changed() {
		db.Statement.SetColumn("lock_version", d.LockVersion+1)
	}
	return
}
