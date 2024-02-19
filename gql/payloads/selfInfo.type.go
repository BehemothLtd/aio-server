package payloads

import (
	"aio-server/exceptions"
	"aio-server/pkg/auths"
	"context"

	"gorm.io/gorm"
)

// SelfInfoResolver resolves self information.
type SelfInfoResolver struct {
	Ctx *context.Context
	Db  *gorm.DB
}

// Resolve resolves the self information.
func (sir *SelfInfoResolver) Resolve() (*UserResolver, error) {
	user, err := auths.AuthUserFromCtx(*sir.Ctx)
	if err != nil {
		return nil, exceptions.NewUnauthorizedError("")
	}

	return &UserResolver{
		Ctx: sir.Ctx,
		Db:  sir.Db,

		User: &user,
	}, nil
}
