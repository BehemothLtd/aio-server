package gql

import (
	"aio-server/database"
	"aio-server/gql/payloads/ms"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"context"
	"errors"
	"fmt"

	graphql "github.com/graph-gophers/graphql-go"
)

func (r *Resolver) MsSnippet(ctx context.Context, args struct{ Id graphql.ID }) (*ms.SnippetResolver, error) {
	if args.Id == "" {
		return nil, errors.New("invalid Id")
	}

	snippetId, err := helpers.GqlIdToInt32(args.Id)
	if err != nil {
		return nil, err
	}

	snippet := models.Snippet{}

	repo := repository.NewSnippetRepository(&ctx, database.Db)
	repo.FindSnippetById(&snippet, snippetId)

	s := ms.SnippetResolver{
		Db:  r.Db,
		Ctx: &ctx,
		M:   &snippet,
	}

	return &s, nil
}

func (r *Resolver) MsSnippets(ctx context.Context, args struct {
	Input *PagyInput
	Query *SnippetQuery
}) (*ms.SnippetsResolver, error) {
	fmt.Printf("ARGS : %+v", args.Input.Page)

	return nil, nil
}

type PagyInput struct {
	PerPage *int32
	Page    *int32
}

type SnippetQuery struct {
	TitleCont *string
}
