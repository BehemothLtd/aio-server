package insightResolvers

import (
	"aio-server/exceptions"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/pkg/auths"
	"aio-server/services/insightServices"
	"context"
)

func (r *Resolver) SelfUpdatePassword(ctx context.Context, args insightInputs.SelfUpdatePasswordInput) (*string, error) {
	user, err := auths.AuthUserFromCtx(ctx)

	if err != nil {
		return nil, exceptions.NewUnauthorizedError("")
	}

	service := insightServices.SelfUpdatePasswordService{
		Ctx:  &ctx,
		Db:   r.Db,
		Args: args,
		User: user,
	}

	if err := service.Execute(); err != nil {
		return nil, err
	}

	message := "Update Password Succesfully"

	return &message, nil
}
