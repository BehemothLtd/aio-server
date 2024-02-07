package gql

import (
	"aio-server/gql/payloads/ms"
	"aio-server/models"
	"context"

	graphql "github.com/graph-gophers/graphql-go"
)

func (r *Resolver) MsSnippet(ctx context.Context, args struct{ ID graphql.ID }) (*ms.SnippetResolver, error) {
	snippet := models.Snippet{Id: 1}

	r.Db.First(&snippet)

	s := ms.SnippetResolver{
		Db:  r.Db,
		Ctx: &ctx,
		M:   &snippet,
	}

	return &s, nil
}
