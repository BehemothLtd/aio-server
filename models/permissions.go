package models

import "aio-server/enums"

type Permission struct {
	Id     int
	Target enums.PermissionTargetType
	Action enums.PermissionActionType
}
