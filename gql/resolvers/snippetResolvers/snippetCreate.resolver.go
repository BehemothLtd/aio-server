package snippetResolvers

import (
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/gqlTypes/msTypes"
	"aio-server/gql/inputs/msInputs"
	"aio-server/services/msServices"
	"context"
)

// MsSnippetCreate resolves the mutation for creating a snippet.
func (r *Resolver) SnippetCreate(ctx context.Context, args msInputs.SnippetCreateInput) (*msTypes.MsSnippetCreatedType, error) {
	service := msServices.SnippetCreateService{
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
