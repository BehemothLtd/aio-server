package models

import "aio-server/enums"

type Device struct {
	Id                    int32
	UserId                int32
	User                  User
	Name                  string
	Code                  string
	Description           string
	State                 enums.DeviceStateType
	DeviceTypeId          int32
	DeviceType            DeviceType
	DevicesUsingHistories []*DevicesUsingHistory
}
