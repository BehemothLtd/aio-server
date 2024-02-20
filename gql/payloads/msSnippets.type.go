package payloads

import (
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/inputs"
	"aio-server/models"
	"aio-server/repository"
	"context"

	"gorm.io/gorm"
)

// MsSnippetsResolver resolves the querying of snippets collection.
type MsSnippetsResolver struct {
	Ctx  *context.Context
	Db   *gorm.DB
	Args inputs.MsSnippetsInput
}

// Resolve executes the snippet listing service and prepares the result for GraphQL.
func (msr *MsSnippetsResolver) Resolve() (*[]*globalTypes.SnippetType, *MetadataResolver, error) {
	var snippets []*models.Snippet
	snippetsQuery, paginationData := msr.Args.ToPaginationDataAndSnippetsQuery()

	repo := repository.NewSnippetRepository(msr.Ctx, msr.Db)

	err := repo.List(&snippets, &paginationData, &snippetsQuery)
	if err != nil {
		return nil, nil, err
	}

	return msr.fromSnippets(snippets), &MetadataResolver{Metadata: &paginationData.Metadata}, nil
}

// fromSnippets converts models.Snippet slice to []*MsSnippetType.
func (msr *MsSnippetsResolver) fromSnippets(snippets []*models.Snippet) *[]*globalTypes.SnippetType {
	resolvers := make([]*globalTypes.SnippetType, len(snippets))
	for i, s := range snippets {
		resolvers[i] = &globalTypes.SnippetType{Snippet: s}
	}
	return &resolvers
}
