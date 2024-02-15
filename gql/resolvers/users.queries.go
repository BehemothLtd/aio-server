package gql

import (
	"aio-server/gql/payloads"
	"context"
)

// SelfInfo resolves the query for retrieving self information.
func (r *Resolver) SelfInfo(ctx context.Context) (*payloads.SelfInfoResolver, error) {
	resolver := payloads.SelfInfoResolver{
		Ctx: &ctx,
		Db:  r.Db,
	}

	if err := resolver.Resolve(); err != nil {
		return nil, err
	}

	return &resolver, nil
}
