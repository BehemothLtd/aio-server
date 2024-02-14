package payloads

import (
	"aio-server/gql/inputs"
	"aio-server/models"
	"aio-server/services"
	"context"

	"github.com/graph-gophers/graphql-go"
	"gorm.io/gorm"
)

type MsSnippetUpdateResolver struct {
	Ctx  *context.Context
	Db   *gorm.DB
	Args struct {
		Id    graphql.ID
		Input inputs.MsSnippetInput
	}

	model *models.Snippet
}

func (msc *MsSnippetUpdateResolver) Resolve() error {
	service := services.SnippetUpdateService{
		Ctx:  msc.Ctx,
		Db:   msc.Db,
		Args: msc.Args,
	}
	snippet, err := service.Execute()

	if err != nil {
		return err
	} else {
		msc.model = snippet

		return nil
	}
}

func (msc *MsSnippetUpdateResolver) Snippet() *MsSnippetResolver {
	return &MsSnippetResolver{
		Snippet: msc.model,
	}
}
