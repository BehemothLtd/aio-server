package msServices

import (
	"aio-server/exceptions"
	"aio-server/gql/inputs/msInputs"
	"aio-server/models"
	"aio-server/pkg/auths"
	"aio-server/repository"
	"aio-server/validators"
	"context"

	"gorm.io/gorm"
)

type CollectionCreateService struct {
	Ctx  *context.Context
	Db   *gorm.DB
	Args msInputs.CollectionCreateInput
}

func (ccs *CollectionCreateService) Excecute() (*models.Collection, error) {
	user, err := auths.AuthUserFromCtx(*ccs.Ctx)

	if err != nil {
		return nil, exceptions.NewUnauthorizedError("")
	}

	collection := models.Collection{
		UserId: user.Id,
	}

	form := validators.NewCollectionFormValidator(
		&ccs.Args.Input,
		repository.NewCollectionRepository(ccs.Ctx, ccs.Db),
		&collection,
	)

	if formErr := form.Save(); formErr != nil {
		return nil, formErr
	} else {
		return &collection, nil
	}
}
