package payloads

import (
	"aio-server/services"
	"context"

	graphql "github.com/graph-gophers/graphql-go"
	"gorm.io/gorm"
)

// MsSnippetFavoriteResolver resolves the favoriting of a snippet.
type MsSnippetFavoriteResolver struct {
	Ctx       *context.Context
	Db        *gorm.DB
	Args      struct{ Id graphql.ID }
	Favorited bool
}

// Resolve executes the snippet favorite service.
func (msfr *MsSnippetFavoriteResolver) Resolve() error {
	service := services.SnippetFavoriteService{
		Ctx:  msfr.Ctx,
		Db:   msfr.Db,
		Args: msfr.Args,
	}

	err := service.Execute()
	if err != nil {
		return err
	}

	msfr.Favorited = *service.Result
	return nil
}
