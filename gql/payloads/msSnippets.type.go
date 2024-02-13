package payloads

import (
	"aio-server/gql/inputs"
	"aio-server/models"
	"aio-server/repository"
	"context"

	"gorm.io/gorm"
)

type MsSnippetsResolver struct {
	Ctx  *context.Context
	Db   *gorm.DB
	Args inputs.MsSnippetsInput

	Collection *[]*MsSnippetResolver
	Metadata   *MetadataResolver
}

func (sr *MsSnippetsResolver) Resolve() error {
	var snippets []*models.Snippet
	snippetsQuery, paginationData := sr.Args.ToPaginationDataAndSnippetsQuery()

	repo := repository.NewSnippetRepository(sr.Ctx, sr.Db)

	err := repo.ListSnippets(&snippets, &paginationData, &snippetsQuery)

	if err != nil {
		return err
	}

	sr.Collection = sr.FromSnippets(snippets)
	sr.Metadata = &MetadataResolver{Metadata: &paginationData.Metadata}

	return nil
}

func (sr *MsSnippetsResolver) FromSnippets(snippets []*models.Snippet) *[]*MsSnippetResolver {
	r := make([]*MsSnippetResolver, len(snippets))
	for i := range snippets {
		r[i] = &MsSnippetResolver{
			Snippet: snippets[i],
		}
	}

	return &r
}
