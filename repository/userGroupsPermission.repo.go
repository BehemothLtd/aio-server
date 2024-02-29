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

// ListAllByUser retrives a list of records based on provided User
func (r *UserGroupsPermissionRepository) ListAllByUser(
	user models.User,
	permissionIds *[]int,
) error {
	usersUserGroupTable := r.db.Model(&models.UsersUserGroup{})
	mainTable := r.db.Model(&models.UserGroupsPermission{})

	return mainTable.Select("DISTINCT permission_id").Where(
		"user_group_id IN (?) AND active = 1",
		usersUserGroupTable.Select("user_group_id").Where("user_id = ?", user.Id),
	).Find(&permissionIds).Error
}
