package gql

import (
	"aio-server/gql/inputs"
	"aio-server/gql/payloads"
	"context"

	graphql "github.com/graph-gophers/graphql-go"
)

// MsSnippetFavorite resolves the mutation for favoriting a snippet.
func (r *Resolver) MsSnippetFavorite(ctx context.Context, args struct {
	Id graphql.ID
}) (*payloads.MsSnippetFavoriteResolver, error) {
	resolver := payloads.MsSnippetFavoriteResolver{
		Ctx:  &ctx,
		Db:   r.Db,
		Args: args,
	}

	if err := resolver.Resolve(); err != nil {
		return nil, err
	}

	return &resolver, nil
}

// MsSnippetCreate resolves the mutation for creating a snippet.
func (r *Resolver) MsSnippetCreate(ctx context.Context, args struct {
	Input inputs.MsSnippetInput
}) (*payloads.MsSnippetCreateResolver, error) {
	resolver := payloads.MsSnippetCreateResolver{
		Ctx:  &ctx,
		Db:   r.Db,
		Args: args,
	}

	if err := resolver.Resolve(); err != nil {
		return nil, err
	}

	return &resolver, nil
}

// MsSnippetUpdate resolves the mutation for updating a snippet.
func (r *Resolver) MsSnippetUpdate(ctx context.Context, args struct {
	Id    graphql.ID
	Input inputs.MsSnippetInput
}) (*payloads.MsSnippetUpdateResolver, error) {
	resolver := payloads.MsSnippetUpdateResolver{
		Ctx:  &ctx,
		Db:   r.Db,
		Args: args,
	}

	if err := resolver.Resolve(); err != nil {
		return nil, err
	}

	return &resolver, nil
}
