package globalTypes

import (
	"aio-server/models"
	"aio-server/pkg/helpers"
	"context"

	graphql "github.com/graph-gophers/graphql-go"
	"gorm.io/gorm"
)

// UserType resolves self information.
type UserType struct {
	Ctx *context.Context
	Db  *gorm.DB

	User *models.User
}

// ID returns the ID of the user.
func (sir *UserType) ID(context.Context) *graphql.ID {
	return helpers.GqlIDP(sir.User.Id)
}

// Email returns the email of the user.
func (sir *UserType) Email(context.Context) *string {
	return &sir.User.Email
}

// FullName returns the full name of the user.
func (sir *UserType) FullName(context.Context) *string {
	return &sir.User.FullName
}

// Name returns the name of the user.
func (sir *UserType) Name(context.Context) *string {
	return &sir.User.Name
}

// LockVersion returns the lock version of the user.
func (sir *UserType) LockVersion(context.Context) int32 {
	return sir.User.LockVersion
}
