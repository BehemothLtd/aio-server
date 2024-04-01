package insightServices

import (
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/repository"
	"aio-server/validators"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type DeviceTypeCreateService struct {
	Ctx        *context.Context
	Db         *gorm.DB
	Args       insightInputs.DeviceTypeCreateInput
	DeviceType *models.DeviceType
}

func (dts *DeviceTypeCreateService) Execute() error {
	form := validators.NewDeviceTypeFormValidation(
		&dts.Args.Input,
		repository.NewDeviceTypeRepository(dts.Ctx, dts.Db),
		dts.DeviceType,
	)

	if err := form.Save(); err != nil {
		return err
	}

	return nil
}
