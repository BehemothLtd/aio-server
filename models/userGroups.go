package models

import (
	"time"
)

type UserGroup struct {
	Id                    int32
	Title                 string
	Users                 []*User `gorm:"many2many:users_user_groups"`
	UserGroupsPermissions []*UserGroupsPermission
	CreatedAt             time.Time
	UpdatedAt             time.Time
}
