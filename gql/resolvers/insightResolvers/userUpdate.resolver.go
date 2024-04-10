package insightResolvers

import (
	"aio-server/exceptions"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/pkg/auths"
	"aio-server/services/insightServices"
	"context"
)

func (r *Resolver) UserUpdate(ctx context.Context, args insightInputs.UserUpdateInput) (*globalTypes.UserUpdatedType, error) {
	user, err := auths.AuthUserFromCtx(ctx)

	if err != nil {
		return nil, exceptions.NewUnauthorizedError("")
	}

	service := insightServices.UserUpdateService{
		Ctx:  &ctx,
		Db:   r.Db,
		Args: args,
		User: &user,
	}
	if err := service.Execute(); err != nil {
		return nil, err
	}

	return nil, nil
}
