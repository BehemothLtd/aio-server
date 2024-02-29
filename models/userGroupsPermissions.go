package models

import "time"

type UserGroupsPermission struct {
	Id           int32 `gorm:"not null;autoIncrement"`
	UserGroupId  int32 `gorm:"not null"`
	PermissionId int   `gorm:"not null"`
	Active       bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
