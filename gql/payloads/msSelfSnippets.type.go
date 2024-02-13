package payloads

import (
	"aio-server/exceptions"
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

	Collection *[]*MsSnippetResolver
	Metadata   *MetadataResolver
}

func (ssr *MsSelfSnippetsResolver) Resolve() error {
	var snippets []*models.Snippet

	user, err := auths.AuthUserFromCtx(*ssr.Ctx)

	if err != nil {
		return exceptions.NewUnauthorizedError("")
	}

	snippetsQuery, paginationData := ssr.Args.ToPaginationDataAndSnippetsQuery()

	repo := repository.NewSnippetRepository(ssr.Ctx, ssr.Db)

	fetchErr := repo.ListSnippetsByUser(&snippets, &paginationData, &snippetsQuery, &user)

	if fetchErr != nil {
		return exceptions.NewUnprocessableContentError(nil)
	}

	ssr.Collection = ssr.FromSnippets(snippets)
	ssr.Metadata = &MetadataResolver{Metadata: &paginationData.Metadata}

	return nil
}

func (sr *MsSelfSnippetsResolver) FromSnippets(snippets []*models.Snippet) *[]*MsSnippetResolver {
	r := make([]*MsSnippetResolver, len(snippets))
	for i := range snippets {
		r[i] = &MsSnippetResolver{
			Snippet: snippets[i],
		}
	}

	return &r
}
