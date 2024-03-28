package models

import "time"

type Attendance struct {
	Id            int32
	UserId        int32
	User          User
	CheckinAt     time.Time
	CheckoutAt    time.Time
	CreatedUserId int32
	CreatedUser   User
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
