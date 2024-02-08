package ms

import (
	"aio-server/gql/payloads"
	"aio-server/models"
	"context"

	"gorm.io/gorm"
)

type SnippetsResolver struct {
	Db  *gorm.DB
	Ctx *context.Context
	C   *models.SnippetsCollection
}

func (sr *SnippetsResolver) Collection(ctx context.Context) *[]*SnippetResolver {
	snippets := sr.C.Collection

	r := make([]*SnippetResolver, len(snippets))
	for i := range snippets {
		r[i] = &SnippetResolver{
			Db: sr.Db,
			M:  snippets[i],
		}
	}

	return &r
}

func (sr *SnippetsResolver) Metadata(ctx context.Context) *payloads.MetadataResolver {
	return &payloads.MetadataResolver{
		M: sr.C.Metadata,
	}
}
