package resolvers

import (
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/gqlTypes/msTypes"
	"aio-server/gql/inputs/msInputs"
	"aio-server/services/msServices"
	"context"
)

// MsSnippetUpdate resolves the mutation for updating a snippet.
func (r *Resolver) MsSnippetUpdate(ctx context.Context, args msInputs.SnippetUpdateInput) (*msTypes.MsSnippetUpdatedType, error) {
	service := msServices.SnippetUpdateService{
		Ctx:  &ctx,
		Db:   r.Db,
		Args: args,
	}

	if snippet, err := service.Execute(); err != nil {
		return nil, err
	} else {
		return &msTypes.MsSnippetUpdatedType{
			Snippet: &globalTypes.SnippetType{
				Snippet: snippet,
			},
		}, nil
	}
}
