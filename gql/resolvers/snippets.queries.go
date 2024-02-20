package resolvers

import (
	"aio-server/gql/inputs"
	"aio-server/gql/payloads"
	"context"
)

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
