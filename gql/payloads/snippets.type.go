package payloads

import (
	"aio-server/models"
	"context"
)

type SnippetsResolver struct {
	SnippetsCollection *models.SnippetsCollection
}

func (sr *SnippetsResolver) Collection(ctx context.Context) *[]*SnippetResolver {
	snippets := sr.SnippetsCollection.Collection

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
		Metadata: sr.SnippetsCollection.Metadata,
	}
}
