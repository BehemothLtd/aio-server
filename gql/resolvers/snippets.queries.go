package gql

import (
	"aio-server/gql/inputs"
	"aio-server/gql/payloads"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"context"
	"errors"

	graphql "github.com/graph-gophers/graphql-go"
)

func (r *Resolver) MsSnippet(ctx context.Context, args struct{ Id graphql.ID }) (*payloads.SnippetResolver, error) {
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

	s := payloads.SnippetResolver{
		Snippet: &snippet,
	}

	return &s, nil
}

func (r *Resolver) MsSnippets(ctx context.Context, args struct {
	Input *inputs.PagyInput
	Query *inputs.SnippetQueryInput
}) (*payloads.SnippetsResolver, error) {
	var snippets []*models.Snippet

	paginationInput := helpers.GeneratePaginationInput(args.Input)

	repo := repository.NewSnippetRepository(&ctx, r.Db)

	outputQuery := models.SnippetsQuery{TitleCont: ""}

	if args.Query != nil && *args.Query.TitleCont != "" {
		outputQuery.TitleCont = *args.Query.TitleCont
	}

	err := repo.ListSnippets(&snippets, &paginationInput, &outputQuery)

	resolver := payloads.SnippetsResolver{
		SnippetsCollection: &models.SnippetsCollection{
			Collection: snippets,
			Metadata:   &paginationInput.Metadata,
		},
	}

	if err != nil {
		return &resolver, err
	}

	return &resolver, nil
}
