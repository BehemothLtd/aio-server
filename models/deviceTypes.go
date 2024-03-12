package models

import "time"

type DeviceType struct {
	Id        int32
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
