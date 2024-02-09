package payloads

import (
	"aio-server/gql/inputs"
	"aio-server/models"
	"aio-server/repository"
	"context"

	"gorm.io/gorm"
)

type SnippetsResolver struct {
	Ctx  *context.Context
	Db   *gorm.DB
	Args inputs.MsSnippetsInput

	Collection *[]*SnippetResolver
	Metadata   *MetadataResolver
}

func (sr *SnippetsResolver) Resolve() error {
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

func (sr *SnippetsResolver) FromSnippets(snippets []*models.Snippet) *[]*SnippetResolver {
	r := make([]*SnippetResolver, len(snippets))
	for i := range snippets {
		r[i] = &SnippetResolver{
			Snippet: snippets[i],
		}
	}

	return &r
}
