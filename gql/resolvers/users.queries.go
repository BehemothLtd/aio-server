package gql

import (
	"aio-server/gql/payloads"
	"context"
)

func (r *Resolver) SelfInfo(ctx context.Context) (*payloads.UserResolver, error) {
	user, err := AuthUserFromCtx(ctx)

	if err != nil {
		return nil, err
	}

	resolver := payloads.UserResolver{
		Ctx:  &ctx,
		Db:   r.Db,
		User: &user,
	}

	return &resolver, nil
}
