package payloads

import (
	"aio-server/models"
	"aio-server/pkg/helpers"
	"context"

	graphql "github.com/graph-gophers/graphql-go"
	"gorm.io/gorm"
)

type UserResolver struct {
	Ctx *context.Context
	Db  *gorm.DB

	User *models.User
}

func (ur *UserResolver) Resolve() error {
	return nil
}

func (ur *UserResolver) ID(context.Context) *graphql.ID {
	return helpers.GqlIDP(ur.User.Id)
}

func (ur *UserResolver) Email(context.Context) *string {
	return &ur.User.Email
}

func (ur *UserResolver) AvatarUrl(context.Context) *string {
	return &ur.User.AvatarURL
}

func (ur *UserResolver) FullName(context.Context) *string {
	return &ur.User.FullName
}

func (ur *UserResolver) Name(context.Context) *string {
	return &ur.User.Name
}
