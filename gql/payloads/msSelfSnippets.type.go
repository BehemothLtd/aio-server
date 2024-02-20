package payloads

import (
	"aio-server/exceptions"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/inputs"
	"aio-server/models"
	"aio-server/pkg/auths"
	"aio-server/repository"
	"context"

	"gorm.io/gorm"
)

type MsSelfSnippetsResolver struct {
	Ctx  *context.Context
	Db   *gorm.DB
	Args inputs.MsSnippetsInput

	Collection *[]*globalTypes.SnippetType
	Metadata   *MetadataResolver
}

func (mssr *MsSelfSnippetsResolver) Resolve() error {
	var snippets []*models.Snippet

	user, err := auths.AuthUserFromCtx(*mssr.Ctx)

	if err != nil {
		return exceptions.NewUnauthorizedError("")
	}

	snippetsQuery, paginationData := mssr.Args.ToPaginationDataAndSnippetsQuery()

	repo := repository.NewSnippetRepository(mssr.Ctx, mssr.Db)

	fetchErr := repo.ListByUser(&snippets, &paginationData, &snippetsQuery, &user)

	if fetchErr != nil {
		return exceptions.NewUnprocessableContentError("", nil)
	}

	mssr.Collection = mssr.FromSnippets(snippets)
	mssr.Metadata = &MetadataResolver{Metadata: &paginationData.Metadata}

	return nil
}

func (mssr *MsSelfSnippetsResolver) FromSnippets(snippets []*models.Snippet) *[]*globalTypes.SnippetType {
	r := make([]*globalTypes.SnippetType, len(snippets))
	for i := range snippets {
		r[i] = &globalTypes.SnippetType{
			Snippet: snippets[i],
		}
	}

	return &r
}
