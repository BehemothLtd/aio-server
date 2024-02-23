package snippetResolvers

import (
	"aio-server/exceptions"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/gqlTypes/snippetTypes"
	"aio-server/gql/inputs/msInputs"
	"aio-server/models"
	"aio-server/pkg/auths"
	"aio-server/repository"
	"context"
)

// Tags resolves the query for retrieving a collection of tags
func (r *Resolver) Tags(ctx context.Context, args msInputs.TagsInput) (*snippetTypes.TagsType, error) {
	// Authenticate the user
	_, authErr := auths.AuthUserFromCtx(ctx)
	if authErr != nil {
		return nil, exceptions.NewUnauthorizedError("")
	}

	var tags []*models.Tag
	tagsQuery, paginationData := args.ToPaginationDataAndSnippetsQuery()

	repo := repository.NewTagRepository(&ctx, r.Db)

	err := repo.List(&tags, &paginationData, &tagsQuery)
	if err != nil {
		return nil, err
	}

	return &snippetTypes.TagsType{
		Collection: r.TagSliceToTypes(tags),
		Metadata: &globalTypes.MetadataType{
			Metadata: &paginationData.Metadata,
		},
	}, nil
}
