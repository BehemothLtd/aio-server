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

func (sir *SelfInfoResolver) Email(context.Context) *string {
	return &sir.User.Email
}

func (sir *SelfInfoResolver) AvatarURL(context.Context) *string {
	return &sir.User.AvatarURL
}

func (sir *SelfInfoResolver) FullName(context.Context) *string {
	return &sir.User.FullName
}

func (sir *SelfInfoResolver) Name(context.Context) *string {
	return &sir.User.Name
}
