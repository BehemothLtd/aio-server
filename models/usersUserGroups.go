package models

import "time"

type UsersUserGroup struct {
	Id          int32
	UserId      int32
	UserGroupId int32
	User        User
	UserGroup   UserGroup
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
