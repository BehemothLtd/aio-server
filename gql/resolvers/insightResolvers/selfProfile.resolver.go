package insightResolvers

import (
	"aio-server/exceptions"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/models"
	"aio-server/pkg/auths"
	"aio-server/repository"
	"context"
)

func (r *Resolver) SelfProfile(ctx context.Context) (*globalTypes.UserType, error) {
	user, err := auths.AuthInsightUserFromCtx(ctx)

	if err != nil {
		return nil, exceptions.NewUnauthorizedError("")
	}

	repo := repository.NewUserRepository(&ctx, r.Db.Preload("ProjectAssignees.User"))
	user = models.User{Id: 1}
	repo.Find(&user)

	return &globalTypes.UserType{
		User: &user,
	}, nil
}
