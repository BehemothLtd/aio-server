package ms

import (
	"aio-server/gql/payloads"
	"context"

	"gorm.io/gorm"
)

type SnippetsResolver struct {
	Db  *gorm.DB
	Ctx *context.Context
}

func (sr *SnippetsResolver) Collection(ctx context.Context) *[]*SnippetResolver {
	return nil
}

func (sr *SnippetsResolver) Metadata(ctx context.Context) *payloads.MetadataResolver {
	return nil
}
