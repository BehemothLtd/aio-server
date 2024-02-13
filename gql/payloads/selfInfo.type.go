package payloads

import (
	"aio-server/models"
	"aio-server/pkg/helpers"
	"context"

	graphql "github.com/graph-gophers/graphql-go"
	"gorm.io/gorm"
)

type SelfInfoResolver struct {
	Ctx *context.Context
	Db  *gorm.DB

	User *models.User
}

func (ur *SelfInfoResolver) Resolve() error {
	return nil
}

func (ur *SelfInfoResolver) ID(context.Context) *graphql.ID {
	return helpers.GqlIDP(ur.User.Id)
}

func (ur *SelfInfoResolver) Email(context.Context) *string {
	return &ur.User.Email
}

func (ur *SelfInfoResolver) AvatarUrl(context.Context) *string {
	return &ur.User.AvatarURL
}

func (ur *SelfInfoResolver) FullName(context.Context) *string {
	return &ur.User.FullName
}

func (ur *SelfInfoResolver) Name(context.Context) *string {
	return &ur.User.Name
}
