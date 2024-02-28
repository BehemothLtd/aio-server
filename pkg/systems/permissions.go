package systems

import (
	"aio-server/enums"
	"slices"
)

type Permission struct {
	Id     int
	Target enums.PermissionTargetType
	Action enums.PermissionActionType
}

// Define a package-level variable for the permissions.
// Although not a constant, treat this as immutable.
var defaultPermissions = []Permission{
	{Id: 999, Target: enums.PermissionTargetTypeAll, Action: enums.PermissionActionTypeAll},

	{Id: 1, Target: enums.PermissionTargetTypeUsers, Action: enums.PermissionActionTypeRead},
	{Id: 2, Target: enums.PermissionTargetTypeUsers, Action: enums.PermissionActionTypeWrite},

	{Id: 100, Target: enums.PermissionTargetTypeProjects, Action: enums.PermissionActionTypeRead},
	{Id: 101, Target: enums.PermissionTargetTypeProjects, Action: enums.PermissionActionTypeWrite},
	{Id: 102, Target: enums.PermissionTargetTypeProjects, Action: enums.PermissionActionTypeDelete},
}

// GetPermissions returns a copy of the default permissions slice.
// This prevents the original slice from being modified.
func GetPermissions() []Permission {
	// Optionally, return a deep copy if the structs contain slice maps, or other pointers.
	return append([]Permission(nil), defaultPermissions...)
}

// FindPermissionById retrieve the permission record
func FindPermissionById(id int) *Permission {
	allPermissions := GetPermissions()

	if foundIdx := slices.IndexFunc(allPermissions, func(p Permission) bool { return p.Id == id }); foundIdx != -1 {
		return &allPermissions[foundIdx]
	} else {
		return nil
	}
}
