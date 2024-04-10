package systems

import (
	"aio-server/enums"
	"aio-server/models"
	"aio-server/repository"
	"context"
	"slices"

	"gorm.io/gorm"
)

type Permission struct {
	Id     int
	Target enums.PermissionTargetType
	Action enums.PermissionActionType
}

// Define a package-level variable for the permissions.
// Although not a constant, treat this as immutable.
var defaultPermissions = []Permission{
	{Id: 9999, Target: enums.PermissionTargetTypeAll, Action: enums.PermissionActionTypeAll},

	{Id: 1, Target: enums.PermissionTargetTypeUsers, Action: enums.PermissionActionTypeRead},
	{Id: 2, Target: enums.PermissionTargetTypeUsers, Action: enums.PermissionActionTypeWrite},

	{Id: 100, Target: enums.PermissionTargetTypeProjects, Action: enums.PermissionActionTypeRead},
	{Id: 101, Target: enums.PermissionTargetTypeProjects, Action: enums.PermissionActionTypeWrite},
	{Id: 102, Target: enums.PermissionTargetTypeProjects, Action: enums.PermissionActionTypeDelete},

	{Id: 110, Target: enums.PermissionTargetTypeProjectIssues, Action: enums.PermissionActionTypeWrite},
	{Id: 111, Target: enums.PermissionTargetTypeProjectIssues, Action: enums.PermissionActionTypeDelete},

	{Id: 200, Target: enums.PermissionTargetTypeUserGroups, Action: enums.PermissionActionTypeRead},
	{Id: 201, Target: enums.PermissionTargetTypeUserGroups, Action: enums.PermissionActionTypeWrite},

	{Id: 400, Target: enums.PermissionTargetTypeClients, Action: enums.PermissionActionTypeRead},
	{Id: 401, Target: enums.PermissionTargetTypeClients, Action: enums.PermissionActionTypeWrite},
	{Id: 402, Target: enums.PermissionTargetTypeClients, Action: enums.PermissionActionTypeDelete},

	{Id: 700, Target: enums.PermissionTargetTypeLeaveDayRequests, Action: enums.PermissionActionTypeRead},
	{Id: 701, Target: enums.PermissionTargetTypeLeaveDayRequests, Action: enums.PermissionActionTypeWrite},
	{Id: 702, Target: enums.PermissionTargetTypeLeaveDayRequests, Action: enums.PermissionActionTypeChangeState},

	{Id: 800, Target: enums.PermissionTargetTypeAttendances, Action: enums.PermissionActionTypeRead},
	{Id: 801, Target: enums.PermissionTargetTypeAttendances, Action: enums.PermissionActionTypeWrite},
}

// GetPermissions returns a copy of the default permissions slice.
// This prevents the original slice from being modified.
func GetPermissions() []Permission {
	// Optionally, return a deep copy if the structs contain slice maps, or other pointers.
	return append([]Permission(nil), defaultPermissions...)
}

// findPermissionById retrieve the permission record
func findPermissionById(id int) *Permission {
	allPermissions := GetPermissions()

	if foundIdx := slices.IndexFunc(allPermissions, func(p Permission) bool { return p.Id == id }); foundIdx != -1 {
		return &allPermissions[foundIdx]
	} else {
		return nil
	}
}

// fetchPermissionsByIds retrieve list of permission record
func fetchPermissionsByIds(ids []int) []Permission {
	result := []Permission{}

	for _, id := range ids {
		permission := findPermissionById(id)

		if permission != nil {
			result = append(result, *permission)
		}
	}

	return result
}

func FetchUserPermissions(ctx context.Context, db gorm.DB, user models.User) []Permission {
	// result := []Permission{}
	repo := repository.NewUserGroupsPermissionRepository(&ctx, &db)
	permissionIds := []int{}
	repo.ListAllByUser(user, &permissionIds)

	return fetchPermissionsByIds(permissionIds)

}

func (p Permission) IsSystemAdmin() bool {
	return p.Id == 9999
}

func (p Permission) IsAuthorizedForTargetAndAction(target string, action string) bool {
	return p.IsSystemAdmin() || (p.Action.String() == action && p.Target.String() == target)
}

func AuthUserToAction(ctx context.Context, db gorm.DB, user models.User, target string, action string) bool {
	permissions := FetchUserPermissions(ctx, db, user)

	foundIdx := slices.IndexFunc(permissions, func(p Permission) bool { return p.IsAuthorizedForTargetAndAction(target, action) })

	return foundIdx != -1
}
