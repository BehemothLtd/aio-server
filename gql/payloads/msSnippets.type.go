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

func (msr *MsSnippetsResolver) Resolve() error {
	var snippets []*models.Snippet
	snippetsQuery, paginationData := msr.Args.ToPaginationDataAndSnippetsQuery()

	repo := repository.NewSnippetRepository(msr.Ctx, msr.Db)

	err := repo.ListSnippets(&snippets, &paginationData, &snippetsQuery)

	if err != nil {
		return err
	}

	msr.Collection = msr.FromSnippets(snippets)
	msr.Metadata = &MetadataResolver{Metadata: &paginationData.Metadata}

	return nil
}

func (msr *MsSnippetsResolver) FromSnippets(snippets []*models.Snippet) *[]*MsSnippetResolver {
	r := make([]*MsSnippetResolver, len(snippets))
	for i := range snippets {
		r[i] = &MsSnippetResolver{
			Snippet: snippets[i],
		}
	}

	return &r
}
