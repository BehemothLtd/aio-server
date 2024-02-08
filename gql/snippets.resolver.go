package gql

import (
	"aio-server/gql/inputs"
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

	repo := repository.NewSnippetRepository(&ctx, r.Db)
	repo.FindSnippetById(&snippet, snippetId)

	s := ms.SnippetResolver{
		Db:  r.Db,
		Ctx: &ctx,
		M:   &snippet,
	}

	return &s, nil
}

func (r *Resolver) MsSnippets(ctx context.Context, args struct {
	Input *inputs.PagyInput
	Query *inputs.SnippetQueryInput
}) (*ms.SnippetsResolver, error) {
	var snippets []*models.Snippet

	paginationInput := helpers.GeneratePaginationInput(args.Input)
	fmt.Printf("PAGINATION INPUT: %+v", paginationInput)

	repo := repository.NewSnippetRepository(&ctx, r.Db)
	fmt.Printf("REPO: %+v", repo)

	outputQuery := models.SnippetsQuery{TitleCont: ""}
	fmt.Printf("QUERY: %+v", outputQuery)

	if args.Query != nil && *args.Query.TitleCont != "" {
		outputQuery.TitleCont = *args.Query.TitleCont
	}

	err := repo.ListSnippets(&snippets, &paginationInput, &outputQuery)

	s := ms.SnippetsResolver{
		Db:  r.Db,
		Ctx: &ctx,
		C: &models.SnippetsCollection{
			Collection: snippets,
			Metadata:   &paginationInput.Metadata,
		},
	}

	if err != nil {
		return &s, err
	}

	return &s, nil
}
