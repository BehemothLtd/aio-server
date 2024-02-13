package gql

import (
	"aio-server/gql/payloads"
	"aio-server/pkg/auths"
	"context"
)

func (r *Resolver) SelfInfo(ctx context.Context) (*payloads.SelfInfoResolver, error) {
	user, err := auths.AuthUserFromCtx(ctx)

	if err != nil {
		return nil, err
	}

	resolver := payloads.SelfInfoResolver{
		Ctx:  &ctx,
		Db:   r.Db,
		User: &user,
	}

	return &resolver, nil
}
