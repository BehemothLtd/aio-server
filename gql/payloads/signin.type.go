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

func (sr *SignInResolver) Errors(ctx context.Context) *[]*ResourceModifyErrorResolver {
	errors := sr.Auth.Errors

	r := make([]*ResourceModifyErrorResolver, len(errors))
	for i := range errors {
		r[i] = &ResourceModifyErrorResolver{
			Error: errors[i],
		}
	}

	return &r
}
