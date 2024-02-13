package payloads

import (
	"aio-server/services"
	"context"

	graphql "github.com/graph-gophers/graphql-go"
	"gorm.io/gorm"
)

type MsSnippetFavoriteResolver struct {
	Ctx  *context.Context
	Db   *gorm.DB
	Args struct{ Id graphql.ID }

	Favorited bool
}

func (sfr *MsSnippetFavoriteResolver) Resolve() error {
	service := services.SnippetFavoriteService{
		Ctx:  sfr.Ctx,
		Db:   sfr.Db,
		Args: sfr.Args,
	}

	err := service.Execute()

	if err != nil {
		return err
	}

	sfr.Favorited = *service.Result

	return nil
}
