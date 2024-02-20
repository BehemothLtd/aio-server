package resolvers

import (
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/gqlTypes/msTypes"
	"aio-server/gql/inputs"
	"aio-server/services"
	"context"
)

// MsSnippetCreate resolves the mutation for creating a snippet.
func (r *Resolver) MsSnippetCreate(ctx context.Context, args inputs.MsSnippetModificationInput) (*msTypes.MsSnippetCreatedType, error) {
	service := services.SnippetCreateService{
		Ctx:  &ctx,
		Db:   r.Db,
		Args: args,
	}

	if snippet, err := service.Execute(); err != nil {
		return nil, err
	} else {
		return &msTypes.MsSnippetCreatedType{
			Snippet: &globalTypes.SnippetType{
				Snippet: snippet,
			},
		}, nil
	}
}
