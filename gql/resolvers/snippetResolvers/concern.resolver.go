package snippetResolvers

import (
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/models"
)

// SnippetSliceToTypes converts models.Snippet slice to []*SnippetType.
func (r *Resolver) SnippetSliceToTypes(snippets []*models.Snippet) *[]*globalTypes.SnippetType {
	resolvers := make([]*globalTypes.SnippetType, len(snippets))
	for i, s := range snippets {
		resolvers[i] = &globalTypes.SnippetType{Snippet: s}
	}
	return &resolvers
}

// TagSliceToTypes converts models.Tag slice to []*TagType.
func (r *Resolver) TagSliceToTypes(tags []*models.Tag) *[]*globalTypes.TagType {
	resolvers := make([]*globalTypes.TagType, len(tags))
	for i, t := range tags {
		resolvers[i] = &globalTypes.TagType{Tag: t}
	}
	return &resolvers
}
