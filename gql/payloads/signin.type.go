package payloads

import (
	"aio-server/models"
	"context"
)

type SignInResolver struct {
	Auth *models.Authentication
}

func (sr *SignInResolver) Message(ctx context.Context) *string {
	return &sr.Auth.Message
}

func (sr *SignInResolver) Token(ctx context.Context) *string {
	return &sr.Auth.Token
}
