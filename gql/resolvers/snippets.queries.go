package gql

import (
	"aio-server/gql/inputs"
	"aio-server/gql/payloads"
	"context"

	graphql "github.com/graph-gophers/graphql-go"
)

func (r *Resolver) MsSnippet(ctx context.Context, args struct{ Id graphql.ID }) (*payloads.MsSnippetResolver, error) {
	resolver := payloads.MsSnippetResolver{
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

func (r *Resolver) MsSnippets(ctx context.Context, args inputs.MsSnippetsInput) (*payloads.SnippetsResolver, error) {
	resolver := payloads.SnippetsResolver{
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

func (r *Resolver) MsSelfSnippets(ctx context.Context, args inputs.MsSnippetsInput) (*payloads.SelfSnippetsResolver, error) {
	resolver := payloads.SelfSnippetsResolver{
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
