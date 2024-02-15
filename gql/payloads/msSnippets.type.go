package payloads

import (
	"aio-server/gql/inputs"
	"aio-server/models"
	"aio-server/repository"
	"context"

	"gorm.io/gorm"
)

// MsSnippetsResolver resolves the querying of snippets collection.
type MsSnippetsResolver struct {
	Ctx        *context.Context
	Db         *gorm.DB
	Args       inputs.MsSnippetsInput
	Collection *[]*MsSnippetResolver
	Metadata   *MetadataResolver
}

// Resolve executes the snippet listing service and prepares the result for GraphQL.
func (msr *MsSnippetsResolver) Resolve() error {
	var snippets []*models.Snippet
	snippetsQuery, paginationData := msr.Args.ToPaginationDataAndSnippetsQuery()

	repo := repository.NewSnippetRepository(msr.Ctx, msr.Db)

	err := repo.List(&snippets, &paginationData, &snippetsQuery)
	if err != nil {
		return err
	}

	msr.Collection = msr.fromSnippets(snippets)
	msr.Metadata = &MetadataResolver{Metadata: &paginationData.Metadata}

	return nil
}

// fromSnippets converts models.Snippet slice to []*MsSnippetResolver.
func (msr *MsSnippetsResolver) fromSnippets(snippets []*models.Snippet) *[]*MsSnippetResolver {
	resolvers := make([]*MsSnippetResolver, len(snippets))
	for i, s := range snippets {
		resolvers[i] = &MsSnippetResolver{Snippet: s}
	}
	return &resolvers
}
