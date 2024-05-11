package insightServices

import (
	"aio-server/exceptions"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/repository"
	"aio-server/validators"
	"context"

	"gorm.io/gorm"
)

type DeviceUpdateService struct {
	Ctx    *context.Context
	Db     *gorm.DB
	Args   insightInputs.DeviceUpdateInput
	Device *models.Device
}

func (dus *DeviceUpdateService) Execute() error {
	repo := repository.NewDeviceRepository(dus.Ctx, dus.Db)

	if err := repo.Find(dus.Device); err != nil {
		return exceptions.NewRecordNotFoundError()
	}

	form := validators.NewDeviceFormValidator(
		&dus.Args.Input,
		repo,
		dus.Device,
	)

	if err := form.Save(); err != nil {
		return err
	}

	return nil
}
