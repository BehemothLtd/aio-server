package insightResolvers

import (
	"aio-server/exceptions"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/pkg/auths"
	"aio-server/repository"
	"context"
)

func (r *Resolver) SelfGeneralInfo(ctx context.Context) (*globalTypes.UserType, error) {
	user, err := auths.AuthInsightUserFromCtx(ctx)

	if err != nil {
		return nil, exceptions.NewUnauthorizedError("")
	}

	repo := repository.NewUserRepository(&ctx, r.Db.Preload("Avatar"))
	repo.FindWithAvatar(&user)

	return &globalTypes.UserType{
		User: &user,
		Db:   r.Db,
	}, nil
}
