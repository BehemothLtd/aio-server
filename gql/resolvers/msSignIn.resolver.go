package resolvers

import (
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/inputs/msInputs"
	"aio-server/services"
	"context"
)

func (r *Resolver) MsSignIn(ctx context.Context, args msInputs.SignInInput) (*globalTypes.SignInType, error) {
	service := services.AuthService{
		Email:    args.Email,
		Password: args.Password,
		Ctx:      &ctx,
		Db:       r.Db,
	}

	err := service.Execute()

	if err != nil {
		return nil, err
	}

	return &globalTypes.SignInType{
		Token: service.Token,
	}, nil
}
