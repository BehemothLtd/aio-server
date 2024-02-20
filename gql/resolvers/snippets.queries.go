package gql

import (
	"aio-server/gql/gqlTypes"
	"aio-server/gql/inputs"
	"aio-server/gql/payloads"
	"context"

	graphql "github.com/graph-gophers/graphql-go"
)

// MsSnippet resolves the query for retrieving a single snippet.
func (r *Resolver) MsSnippet(ctx context.Context, args struct{ Id graphql.ID }) (*gqlTypes.SnippetResolver, error) {
	resolver := payloads.MsSnippetResolver{
		Ctx:  &ctx,
		Db:   r.Db,
		Args: args,
	}

	if snippetResolver, err := resolver.Resolve(); err != nil {
		return nil, err
	} else {
		return snippetResolver, nil
	}
}

type SnippetsCollection struct {
	Collection *[]*gqlTypes.SnippetResolver
	Metadata   *payloads.MetadataResolver
}

// MsSnippets resolves the query for retrieving a collection of snippets.
func (r *Resolver) MsSnippets(ctx context.Context, args inputs.MsSnippetsInput) (*SnippetsCollection, error) {
	resolver := payloads.MsSnippetsResolver{
		Ctx:  &ctx,
		Db:   r.Db,
		Args: args,
	}

	if collectionResolver, metadataResolver, err := resolver.Resolve(); err != nil {
		return nil, err
	} else {
		return &SnippetsCollection{
			Collection: collectionResolver,
			Metadata:   metadataResolver,
		}, nil
	}
}

// MsSelfSnippets resolves the query for retrieving self-owned snippets.
func (r *Resolver) MsSelfSnippets(ctx context.Context, args inputs.MsSnippetsInput) (*payloads.MsSelfSnippetsResolver, error) {
	resolver := payloads.MsSelfSnippetsResolver{
		Ctx:  &ctx,
		Db:   r.Db,
		Args: args,
	}

	if err := resolver.Resolve(); err != nil {
		return nil, err
	}

	return &resolver, nil
}

func (r *Resolver) MsSnippetDecryptContent(ctx context.Context, args inputs.MsSnippetDecryptContentInput) (*string, error) {
	resolver := payloads.MsSnippetDecryptContentResolver{
		Ctx:  &ctx,
		Db:   r.Db,
		Args: args,
	}

	if decryptedContent, err := resolver.Resolve(); err != nil {
		return nil, err
	} else {
		return decryptedContent, err
	}
}
