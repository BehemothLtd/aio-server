package models

import "time"

type Pin struct {
	Id         int32
	UserId     int32
	ParentType int
	ParentId   int32
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
