package gql

import (
	"aio-server/gql/payloads"
	"aio-server/services"
	"context"
)

func (r *Resolver) SignIn(ctx context.Context, args struct {
	Email    string
	Password string
}) (*payloads.SignInResolver, error) {
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

	return &payloads.SignInResolver{
		Token: service.Token,
	}, nil
}
