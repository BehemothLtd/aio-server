package payloads

import (
	"aio-server/gql/inputs"
	"aio-server/models"
	"aio-server/services"
	"context"

	"github.com/graph-gophers/graphql-go"
	"gorm.io/gorm"
)

// MsSnippetUpdateResolver resolves the updating of a snippet.
type MsSnippetUpdateResolver struct {
	Ctx   *context.Context
	Db    *gorm.DB
	Args  struct {
		Id    graphql.ID
		Input inputs.MsSnippetInput
	}
	Model *models.Snippet
}

// Resolve executes the snippet update service.
func (msc *MsSnippetUpdateResolver) Resolve() error {
	service := services.SnippetUpdateService{
		Ctx:  msc.Ctx,
		Db:   msc.Db,
		Args: msc.Args,
	}
	snippet, err := service.Execute()

	if err != nil {
		return err
	}

	msc.Model = snippet
	return nil
}

// Snippet returns the updated snippet.
func (msc *MsSnippetUpdateResolver) Snippet() *MsSnippetResolver {
	return &MsSnippetResolver{
		Snippet: msc.Model,
	}
}
