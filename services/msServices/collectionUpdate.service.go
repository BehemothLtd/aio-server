package msServices

import (
	"aio-server/exceptions"
	"aio-server/gql/inputs/msInputs"
	"aio-server/models"
	"aio-server/pkg/auths"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"aio-server/validators"
	"context"

	"gorm.io/gorm"
)

type CollectionUpdateService struct {
	Ctx  *context.Context
	Db   *gorm.DB
	Args msInputs.CollectionUpdateInput
}

func (cus *CollectionUpdateService) Execute() (*models.Collection, error) {
	// Authenticate the user
	user, err := auths.AuthUserFromCtx(*cus.Ctx)
	if err != nil {
		return nil, exceptions.NewUnauthorizedError("")
	}

	collectionRepo := repository.NewCollectionRepository(cus.Ctx, cus.Db)

	collectionId, err := helpers.GqlIdToInt32(cus.Args.Id)

	if err != nil {
		return nil, exceptions.NewBadRequestError("Invalid Id")
	}

	collection := models.Collection{
		Id: collectionId,
	}

	collectionErr := collectionRepo.FindByUser(&collection, user.Id)

	if collectionErr != nil {
		return nil, exceptions.NewRecordNotFoundError()
	}

	form := validators.NewCollectionFormValidator(
		&cus.Args.Input,
		repository.NewCollectionRepository(cus.Ctx, cus.Db),
		&collection,
	)

	if formErr := form.Save(); formErr != nil {
		return nil, formErr
	} else {
		return &collection, nil
	}
}
