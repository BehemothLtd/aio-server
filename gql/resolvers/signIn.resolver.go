package resolvers

import (
	"aio-server/gql/gqlTypes"
	"aio-server/gql/inputs"
	"aio-server/services"
	"context"
)

func (r *Resolver) SignIn(ctx context.Context, args inputs.SignInInput) (*gqlTypes.SignInType, error) {
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

	return &gqlTypes.SignInType{
		Token: service.Token,
	}, nil
}
