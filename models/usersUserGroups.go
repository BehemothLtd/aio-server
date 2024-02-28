package models

import "time"

type UsersUserGroup struct {
	Id          int32 `gorm:"not null;autoIncrement"`
	UserId      int32 `gorm:"not null"`
	UserGroupId int32 `gorm:"not null"`
	User        User
	UserGroup   UserGroup
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
