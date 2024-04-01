package insightServices

import (
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/repository"
	"aio-server/validators"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type ClientCreateService struct {
	Ctx    *context.Context
	Db     *gorm.DB
	Args   insightInputs.ClientCreateInput
	Client *models.Client
}

func (ccs *ClientCreateService) Execute() error {
	form := validators.NewClientFormValidator(
		&ccs.Args.Input,
		repository.NewClientRepository(ccs.Ctx, ccs.Db),
		ccs.Client,
	)

	if err := form.Save(); err != nil {
		return err
	}

	return nil
}
