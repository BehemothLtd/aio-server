package models

import (
	"time"

	"gorm.io/gorm"
)

type UserGroup struct {
	Id                    int32
	Title                 string
	Users                 []*User `gorm:"many2many:users_user_groups"`
	UserGroupsPermissions []*UserGroupsPermission
	CreatedAt             time.Time
	UpdatedAt             time.Time
}

func (ug *UserGroup) AfterCreate(tx *gorm.DB) (err error) {
	// if ug.Title == constants.SuperAdmin {
	// 	adminPermission := UserGroupsPermission{
	// 		PermissionId: constants.SystemAdminId,
	// 		UserGroupId:  ug.Id,
	// 	}

	// 	// Create or create permission `All` for super admin group
	// 	if err := tx.Table("user_groups_permissions").Find(&adminPermission).Error; err == gorm.ErrRecordNotFound {
	// 		tx.Table("user_groups_permissions").Create(adminPermission)
	// 	}
	// } else {
	// 	permissions := systems.GetPermissions()

	// 	for _, permission := range permissions {
	// 		ugp := UserGroupsPermission{
	// 			PermissionId: permission.Id,
	// 			UserGroupId:  ug.Id,
	// 		}

	// 		// Create or create permissions for user group
	// 		if err := tx.Table("user_groups_permissions").Find(&ugp).Error; err == gorm.ErrRecordNotFound {
	// 			ugp.Active = false
	// 			tx.Table("user_groups_permissions").Create(ugp)
	// 		}
	// 	}
	// }
	// seeds.InitPermissions(ug, tx)

	return
}
