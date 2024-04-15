package insightResolvers

import (
	"aio-server/exceptions"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/pkg/auths"
	"aio-server/pkg/systems"
	"context"
)

func (r *Resolver) SelfPermission(ctx context.Context) ([]*globalTypes.PermissionType, error) {
	user, err := auths.AuthUserFromCtx(ctx)
	if err != nil {
		return nil, exceptions.NewUnauthorizedError("")
	}

	userPermissions := systems.FetchUserPermissions(ctx, *r.Db, user)
	mappedPermissions := make([]*globalTypes.PermissionType, len(userPermissions))

	for i, permission := range userPermissions {
		mappedPermissions[i] = &globalTypes.PermissionType{
			Permission: permission,
		}
	}

	return mappedPermissions, nil
}
