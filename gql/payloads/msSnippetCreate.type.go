package payloads

import (
	"aio-server/gql/inputs"
	"aio-server/models"
	"aio-server/services"
	"context"

	"gorm.io/gorm"
)

type MsSnippetCreateResolver struct {
	Ctx  *context.Context
	Db   *gorm.DB
	Args struct{ Input inputs.MsSnippetInput }

	model *models.Snippet
}

func (msc *MsSnippetCreateResolver) Resolve() error {
	service := services.SnippetCreateService{
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

func (msc *MsSnippetCreateResolver) Snippet() *MsSnippetResolver {
	return &MsSnippetResolver{
		Snippet: msc.model,
	}
}
