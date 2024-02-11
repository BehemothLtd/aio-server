package gql

import (
	"aio-server/gql/payloads"
	"context"

	graphql "github.com/graph-gophers/graphql-go"
)

func (r *Resolver) MsSnippetFavorite(ctx context.Context, args struct{ Id graphql.ID }) (*payloads.SnippetFavoriteResolver, error) {
	resolver := payloads.SnippetFavoriteResolver{
		Ctx:  &ctx,
		Db:   r.Db,
		Args: args,
	}

	err := resolver.Resolve()

	if err != nil {
		return nil, err
	}

	return &resolver, nil
}
