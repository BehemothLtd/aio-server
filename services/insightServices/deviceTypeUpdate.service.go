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

type DeviceTypeUpdateService struct {
	Ctx        *context.Context
	Db         *gorm.DB
	Args       insightInputs.DeviceTypeUpdateInput
	DeviceType *models.DeviceType
}

func (dtus *DeviceTypeUpdateService) Execute() error {
	repo := repository.NewDeviceTypeRepository(dtus.Ctx, dtus.Db)

	if err := repo.Find(dtus.DeviceType); err != nil {
		return exceptions.NewRecordNotFoundError()
	}

	form := validators.NewDeviceTypeFormValidation(
		&dtus.Args.Input,
		repo,
		dtus.DeviceType,
	)

	if err := form.Save(); err != nil {
		return err
	}

	return nil
}
