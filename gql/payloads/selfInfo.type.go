package payloads

import (
	"aio-server/exceptions"
	"aio-server/models"
	"aio-server/pkg/auths"
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

func (sir *SelfInfoResolver) Resolve() error {
	user, err := auths.AuthUserFromCtx(*sir.Ctx)

	if err != nil {
		return exceptions.NewUnauthorizedError("")
	}

	sir.User = &user

	return nil
}

func (sir *SelfInfoResolver) ID(context.Context) *graphql.ID {
	return helpers.GqlIDP(sir.User.Id)
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
