package insightServices

import (
	"aio-server/exceptions"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/repository"
	"aio-server/validators"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type ClientUpdateService struct {
	Ctx         *context.Context
	Db          *gorm.DB
	Args        insightInputs.ClientUpdateInput
	Client 			*models.Client
}

func (cus *ClientUpdateService) Execute() error {
	repo := repository.NewClientRepository(cus.Ctx, cus.Db)

	if err := repo.Find(cus.Client); err != nil {
		return exceptions.NewRecordNotFoundError()
	}

	form := validators.NewClientUpdateFormValidator(
		&cus.Args.Input,
		repo,
		cus.Client,
	) 

	if err := form.Save(); err != nil {
		return err
	}

	return nil
}
