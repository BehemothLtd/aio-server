package repository

import (
	"aio-server/models"
	"context"

	"gorm.io/gorm"
)

type UserGroupsPermissionRepository struct {
	Repository
}

// NewUserGroupsPermissionRepository initializes a new UserGroupsPermissionRepository instance.
func NewUserGroupsPermissionRepository(c *context.Context, db *gorm.DB) *UserGroupsPermissionRepository {
	return &UserGroupsPermissionRepository{
		Repository: Repository{
			db:  db,
			ctx: c,
		},
	}
}

// ListAllByUserGroup retrives a list of records based on provided UserGroup
func (r *UserGroupsPermissionRepository) ListAllByUserGroup(
	userGroup models.UserGroup,
	userGroupsPermissions *[]models.UserGroupsPermission,
) error {
	dbTables := r.db.Model(&models.UserGroupsPermission{})

	return dbTables.Where("user_group_id = ? AND active = 1", userGroup.Id).Find(&userGroupsPermissions).Error
}
