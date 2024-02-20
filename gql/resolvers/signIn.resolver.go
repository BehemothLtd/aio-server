package resolvers

import (
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/inputs"
	"aio-server/services"
	"context"
)

func (r *Resolver) SignIn(ctx context.Context, args inputs.SignInInput) (*globalTypes.SignInType, error) {
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
