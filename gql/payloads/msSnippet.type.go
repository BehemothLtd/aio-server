package payloads

import (
	"aio-server/exceptions"
	"aio-server/gql/gqlTypes"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"context"

	graphql "github.com/graph-gophers/graphql-go"
	"gorm.io/gorm"
)

type MsSnippetResolver struct {
	Ctx  *context.Context
	Db   *gorm.DB
	Args struct{ Id graphql.ID }
}

func (msr *MsSnippetResolver) Resolve() (*gqlTypes.SnippetResolver, error) {
	if msr.Args.Id == "" {
		return nil, exceptions.NewBadRequestError("Invalid Id")
	}

	snippetId, err := helpers.GqlIdToInt32(msr.Args.Id)
	if err != nil {
		return nil, err
	}

	snippet := models.Snippet{}
	repo := repository.NewSnippetRepository(msr.Ctx, msr.Db)
	err = repo.FindById(&snippet, snippetId)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exceptions.NewRecordNotFoundError()
		}
		return nil, err
	}

	return &gqlTypes.SnippetResolver{Snippet: &snippet}, nil
}
