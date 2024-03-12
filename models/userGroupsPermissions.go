package models

import "time"

type UserGroupsPermission struct {
	Id           int32
	UserGroupId  int32
	PermissionId int
	Active       bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
