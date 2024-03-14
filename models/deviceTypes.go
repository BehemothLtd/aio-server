package models

import "time"

type DeviceType struct {
	Id        int32
	Name      string
	Devices   []*Device
	CreatedAt time.Time
	UpdatedAt time.Time
}
