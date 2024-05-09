package seeds

import (
	"aio-server/models"
	"aio-server/pkg/constants"
	"aio-server/pkg/systems"

	"gorm.io/gorm"
)

func InitPermissions(ug models.UserGroup, tx *gorm.DB) (err error) {
	if ug.Title == constants.SuperAdmin {
		adminPermission := models.UserGroupsPermission{
			PermissionId: constants.SystemAdminId,
			UserGroupId:  ug.Id,
		}

		// Create or create permission `All` for super admin group
		if err := tx.Table("user_groups_permissions").Find(&adminPermission).Error; err == gorm.ErrRecordNotFound {
			tx.Table("user_groups_permissions").Create(adminPermission)
		}
	} else {
		permissions := systems.GetPermissions()

		for _, permission := range permissions {
			ugp := models.UserGroupsPermission{
				PermissionId: permission.Id,
				UserGroupId:  ug.Id,
			}

			// Create or create permissions for user group
			if err := tx.Table("user_groups_permissions").Find(&ugp).Error; err == gorm.ErrRecordNotFound {
				ugp.Active = false
				tx.Table("user_groups_permissions").Create(ugp)
			}
		}
	}
	return
}
