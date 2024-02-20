package snippetResolvers

import (
	"aio-server/gql/gqlTypes/globalTypes"
	snippetTypes "aio-server/gql/gqlTypes/snippetTypes"
	"aio-server/gql/inputs/msInputs"
	"aio-server/services/msServices"
	"context"
)

// MsSnippetCreate resolves the mutation for creating a snippet.
func (r *Resolver) SnippetCreate(ctx context.Context, args msInputs.SnippetCreateInput) (*snippetTypes.SnippetCreatedType, error) {
	service := msServices.SnippetCreateService{
		Ctx:  &ctx,
		Db:   r.Db,
		Args: args,
	}

	if snippet, err := service.Execute(); err != nil {
		return nil, err
	} else {
		return &snippetTypes.SnippetCreatedType{
			Snippet: &globalTypes.SnippetType{
				Snippet: snippet,
			},
		}, nil
	}
}
