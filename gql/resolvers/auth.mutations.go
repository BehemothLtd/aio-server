package gql

import (
	"aio-server/database"
	"aio-server/gql/payloads"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"context"
)

func (r *Resolver) SignIn(ctx context.Context, args struct {
	Email    string
	Password string
}) (*payloads.SignInResolver, error) {
	repo := repository.NewUserRepository(&ctx, database.Db)

	user, err := repo.AuthUser(args.Email, args.Password)

	if err != nil {
		return nil, err
	}

	token, err := helpers.GenerateJwtToken(user.GenerateJwtClaims())

	if err != nil {
		return nil, err
	}

	s := payloads.SignInResolver{
		Auth: &models.Authentication{
			Message: "Successfully Signed In",
			Token:   token,
		},
	}

	return &s, nil
}
