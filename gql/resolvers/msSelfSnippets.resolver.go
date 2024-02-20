package resolvers

import (
	"aio-server/exceptions"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/gqlTypes/msTypes"
	"aio-server/gql/inputs/msInputs"
	"aio-server/models"
	"aio-server/pkg/auths"
	"aio-server/repository"
	"context"
)

// MsSelfSnippets resolves the query for retrieving self-owned snippets.
func (r *Resolver) MsSelfSnippets(ctx context.Context, args msInputs.SnippetsInput) (*msTypes.MsSnippetsType, error) {
	var snippets []*models.Snippet

	user, err := auths.AuthUserFromCtx(ctx)

	if err != nil {
		return nil, exceptions.NewUnauthorizedError("")
	}

	snippetsQuery, paginationData := args.ToPaginationDataAndSnippetsQuery()

	repo := repository.NewSnippetRepository(&ctx, r.Db)

	fetchErr := repo.ListByUser(&snippets, &paginationData, &snippetsQuery, &user)

	if fetchErr != nil {
		return nil, exceptions.NewBadRequestError("")
	}

	return &msTypes.MsSnippetsType{
		Collection: r.SnippetSliceToTypes(snippets),
		Metadata: &globalTypes.MetadataType{
			Metadata: &paginationData.Metadata,
		},
	}, nil
}
