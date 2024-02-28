package insightResolvers

import (
	"aio-server/exceptions"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/models"
	"aio-server/pkg/auths"
	"aio-server/pkg/systems"
	"aio-server/repository"
	"context"
	"fmt"
)

// fromSnippets converts models.Snippet slice to []*UserGroupType.
func (r *Resolver) UserGroupSliceToTypes(userGroups []*models.UserGroup) *[]*globalTypes.UserGroupType {
	resolvers := make([]*globalTypes.UserGroupType, len(userGroups))
	for i, s := range userGroups {
		resolvers[i] = &globalTypes.UserGroupType{UserGroup: s}
	}
	return &resolvers
}

func (r *Resolver) Authorize(ctx context.Context, target string, action string) (*models.User, error) {
	user, err := auths.AuthUserFromCtx(ctx)

	if err != nil {
		return nil, exceptions.NewUnauthorizedError("")
	}

	// groupsOfUser := []models.UserGroup{}
	// repo := repository.NewUserGroupRepository(&ctx, r.Db)
	// err = repo.FindAllByUser(&groupsOfUser, user.Id)
	repo := repository.NewUserRepository(&ctx, r.Db)
	repo.FindByIdWithGroups(&user, user.Id)

	fmt.Printf("USER GROUPS %+v", user.UserGroups)

	// if err != nil {
	// 	return nil, exceptions.NewUnauthorizedError("")
	// }

	return nil, nil
}

func (r *Resolver) UserGroupPermissions(ug models.UserGroup) []*systems.Permission {
	userGroupsPermissions := []models.UserGroupsPermission{}
	repo := repository.NewUserGroupsPermissionRepository(nil, r.Db)
	repo.ListAllByUserGroup(ug, &userGroupsPermissions)

	permissions := make([]systems.Permission, len(userGroupsPermissions))

	for i, ugp := range userGroupsPermissions {
		permissions[i] = *systems.FindPermissionById(ugp.PermissionId)
	}

	return []*systems.Permission{}
}
