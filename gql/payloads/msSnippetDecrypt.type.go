package payloads

import (
	"aio-server/gql/inputs"
	"aio-server/services"
	"context"

	"gorm.io/gorm"
)

type MsSnippetDecryptContentResolver struct {
	Ctx  *context.Context
	Db   *gorm.DB
	Args inputs.MsSnippetDecryptContentInput
}

func (msdcr *MsSnippetDecryptContentResolver) Resolve() (*string, error) {
	service := services.SnippetDecryptService{
		Ctx:  msdcr.Ctx,
		Db:   msdcr.Db,
		Args: msdcr.Args,
	}

	if decryptedContent, err := service.Execute(); err != nil {
		return nil, err
	} else {
		return decryptedContent, nil
	}
}
