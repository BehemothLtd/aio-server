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

// SelfInfoResolver resolves self information.
type SelfInfoResolver struct {
	Ctx  *context.Context
	Db   *gorm.DB
	User *models.User
}

// Resolve resolves the self information.
func (sir *SelfInfoResolver) Resolve() error {
	user, err := auths.AuthUserFromCtx(*sir.Ctx)
	if err != nil {
		return exceptions.NewUnauthorizedError("")
	}

	sir.User = &user
	return nil
}

// ID returns the ID of the user.
func (sir *SelfInfoResolver) ID(context.Context) *graphql.ID {
	return helpers.GqlIDP(sir.User.Id)
}

// Email returns the email of the user.
func (sir *SelfInfoResolver) Email(context.Context) *string {
	return &sir.User.Email
}

// AvatarURL returns the avatar URL of the user.
func (sir *SelfInfoResolver) AvatarURL(context.Context) *string {
	return &sir.User.AvatarURL
}

// FullName returns the full name of the user.
func (sir *SelfInfoResolver) FullName(context.Context) *string {
	return &sir.User.FullName
}

// Name returns the name of the user.
func (sir *SelfInfoResolver) Name(context.Context) *string {
	return &sir.User.Name
}
