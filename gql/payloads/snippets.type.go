package payloads

import (
	"aio-server/models"
	"context"
)

type SnippetsResolver struct {
	C *models.SnippetsCollection
}

func (sr *SnippetsResolver) Collection(ctx context.Context) *[]*SnippetResolver {
	snippets := sr.C.Collection

	r := make([]*SnippetResolver, len(snippets))
	for i := range snippets {
		r[i] = &SnippetResolver{
			Snippet: snippets[i],
		}
	}

	return &r
}

func (sr *SnippetsResolver) Metadata(ctx context.Context) *MetadataResolver {
	return &MetadataResolver{
		M: sr.C.Metadata,
	}
}
