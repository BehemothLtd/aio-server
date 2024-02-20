package snippetResolvers

import (
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/models"
)

// fromSnippets converts models.Snippet slice to []*MsSnippetType.
func (r *Resolver) SnippetSliceToTypes(snippets []*models.Snippet) *[]*globalTypes.SnippetType {
	resolvers := make([]*globalTypes.SnippetType, len(snippets))
	for i, s := range snippets {
		resolvers[i] = &globalTypes.SnippetType{Snippet: s}
	}
	return &resolvers
}