package models

import (
	"time"
)

type UserGroup struct {
	Id                    int32   `gorm:"not null;autoIncrement"`
	Title                 string  `gorm:"not null;type:varchar(255);default:null"`
	Users                 []*User `gorm:"many2many:users_user_groups"`
	UserGroupsPermissions []*UserGroupsPermission
	CreatedAt             time.Time
	UpdatedAt             time.Time
}
