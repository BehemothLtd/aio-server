package resolvers

import (
	"aio-server/gql/payloads"
	"context"
)

// SelfInfo resolves the query for retrieving self information.
func (r *Resolver) SelfInfo(ctx context.Context) (*payloads.UserResolver, error) {
	resolver := payloads.SelfInfoResolver{
		Ctx: &ctx,
		Db:  r.Db,
	}

	if result, err := resolver.Resolve(); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}
