package gql

import (
	"aio-server/database"
	"aio-server/gql/payloads"
	"aio-server/models"
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
		Db:       database.Db,
		Result: &models.Authentication{
			Errors: []*models.ResourceModifyErrors{},
		},
	}

	service.Execute()

	return &payloads.SignInResolver{
		Auth: service.Result,
	}, nil
}
