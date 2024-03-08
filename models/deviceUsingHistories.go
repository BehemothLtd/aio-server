package models

import (
	"aio-server/enums"
	"time"
)

type DevicesUsingHistory struct {
	Id          int32
	UserId      int32
	User        User
	DeviceId    int32
	Device      Device
	State       enums.DeviceStateType
	LockVersion int32
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
