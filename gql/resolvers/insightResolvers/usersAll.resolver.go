package insightResolvers

import (
	"aio-server/exceptions"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/models"
	"aio-server/pkg/auths"
	"aio-server/repository"
	"context"
)

func (r *Resolver) UsersAll(ctx context.Context) (*[]*globalTypes.UserType, error) {
	_, err := auths.AuthInsightUserFromCtx(ctx)

	if err != nil {
		return nil, exceptions.NewUnauthorizedError("")
	}

	users := []*models.User{}
	repo := repository.NewUserRepository(&ctx, r.Db)

	if err := repo.All(&users); err != nil {
		return nil, err
	} else {
		result := make([]*globalTypes.UserType, len(users))

		for i, user := range users {
			result[i] = &globalTypes.UserType{
				User: user,
			}
		}
		return &result, nil
	}
}
