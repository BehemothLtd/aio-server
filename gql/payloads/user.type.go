package payloads

import (
	"aio-server/models"
	"aio-server/pkg/helpers"
	"context"

	graphql "github.com/graph-gophers/graphql-go"
	"gorm.io/gorm"
)

// UserResolver resolves self information.
type UserResolver struct {
	Ctx *context.Context
	Db  *gorm.DB

	User *models.User
}

// ID returns the ID of the user.
func (sir *UserResolver) ID(context.Context) *graphql.ID {
	return helpers.GqlIDP(sir.User.Id)
}

// Email returns the email of the user.
func (sir *UserResolver) Email(context.Context) *string {
	return &sir.User.Email
}

// FullName returns the full name of the user.
func (sir *UserResolver) FullName(context.Context) *string {
	return &sir.User.FullName
}

// Name returns the name of the user.
func (sir *UserResolver) Name(context.Context) *string {
	return &sir.User.Name
}
