package snippetResolvers

import (
	"aio-server/gql/gqlTypes/globalTypes"
	snippetTypes "aio-server/gql/gqlTypes/snippetTypes"
	"aio-server/gql/inputs/msInputs"
	"aio-server/services/msServices"
	"context"
)

// MsSnippetUpdate resolves the mutation for updating a snippet.
func (r *Resolver) SnippetUpdate(ctx context.Context, args msInputs.SnippetUpdateInput) (*snippetTypes.SnippetUpdatedType, error) {
	service := msServices.SnippetUpdateService{
		Ctx:  &ctx,
		Db:   r.Db,
		Args: args,
	}

	if snippet, err := service.Execute(); err != nil {
		return nil, err
	} else {
		return &snippetTypes.SnippetUpdatedType{
			Snippet: &globalTypes.SnippetType{
				Snippet: snippet,
			},
		}, nil
	}
}
