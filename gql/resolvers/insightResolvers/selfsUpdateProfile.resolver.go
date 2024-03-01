package insightResolvers

import (
	"aio-server/exceptions"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/pkg/auths"
	"aio-server/services/insightServices"
	"context"
)

func (r *Resolver) SelfsUpdateProfile(ctx context.Context, args insightInputs.SelfsUpdateProfileInput) (*globalTypes.UserUpdatedType, error) {
	user, err := auths.AuthUserFromCtx(ctx)

	if err != nil {
		return nil, exceptions.NewUnauthorizedError("")
	}

	service := insightServices.SelfsUpdateProfileService{
		Ctx:  &ctx,
		Db:   r.Db,
		Args: args,
		User: &user,
	}
	if err := service.Execute(); err != nil {
		return nil, err
	}

	return &globalTypes.UserUpdatedType{User: &globalTypes.UserType{User: &user}}, nil
}
