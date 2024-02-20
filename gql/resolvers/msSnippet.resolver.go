package resolvers

import (
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/payloads"
	"context"

	graphql "github.com/graph-gophers/graphql-go"
)

// MsSnippet resolves the query for retrieving a single snippet.
func (r *Resolver) MsSnippet(ctx context.Context, args struct{ Id graphql.ID }) (*globalTypes.SnippetType, error) {
	resolver := payloads.MsSnippetType{
		Ctx:  &ctx,
		Db:   r.Db,
		Args: args,
	}

	if SnippetType, err := resolver.Resolve(); err != nil {
		return nil, err
	} else {
		return SnippetType, nil
	}
}
