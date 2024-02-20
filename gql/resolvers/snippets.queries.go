package resolvers

import (
	"aio-server/gql/inputs"
	"aio-server/gql/payloads"
	"context"
)

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
