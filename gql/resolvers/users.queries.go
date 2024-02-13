package gql

import (
	"aio-server/gql/payloads"
	"context"
)

func (r *Resolver) SelfInfo(ctx context.Context) (*payloads.SelfInfoResolver, error) {
	resolver := payloads.SelfInfoResolver{
		Ctx: &ctx,
		Db:  r.Db,
	}

	err := resolver.Resolve()

	if err != nil {
		return nil, err
	}

	return &resolver, nil
}
