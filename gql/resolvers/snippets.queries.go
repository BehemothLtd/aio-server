package resolvers

import (
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/inputs"
	"aio-server/gql/payloads"
	"context"
)

type SnippetsCollection struct {
	Collection *[]*globalTypes.SnippetType
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
