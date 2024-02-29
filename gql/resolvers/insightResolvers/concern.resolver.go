package insightResolvers

import (
	"aio-server/exceptions"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/models"
	"aio-server/pkg/auths"
	"aio-server/pkg/systems"
	"context"
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

	if !systems.AuthUserToAction(ctx, *r.Db, user, target, action) {
		return nil, exceptions.NewUnauthorizedError("You dont have authorization for this action")
	}

	return nil, nil
}
