package payloads

import (
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/inputs"
	"aio-server/models"
	"aio-server/services"
	"context"

	"gorm.io/gorm"
)

// MsSnippetCreateResolver resolves the creation of a snippet.
type MsSnippetCreateResolver struct {
	Ctx  *context.Context
	Db   *gorm.DB
	Args struct {
		Input inputs.MsSnippetInput
	}

	Model *models.Snippet
}

// Resolve executes the snippet creation service.
func (msc *MsSnippetCreateResolver) Resolve() error {
	service := services.SnippetCreateService{
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

// Snippet returns the created snippet.
func (msc *MsSnippetCreateResolver) Snippet() *globalTypes.SnippetType {
	return &globalTypes.SnippetType{
		Snippet: msc.Model,
	}
}
