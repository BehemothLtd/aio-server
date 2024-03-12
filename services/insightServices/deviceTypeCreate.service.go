package insightServices

import (
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/repository"
	"aio-server/validators"
	"context"

	"gorm.io/gorm"
)

type DeviceTypeCreateService struct {
	Ctx        *context.Context
	Db         *gorm.DB
	Args       insightInputs.DeviceTypeCreateInput
	DeviceType *models.DeviceType
}

func (dtcs *DeviceTypeCreateService) Execute() (*models.DeviceType, error) {
	deviceType := models.DeviceType{}

	form := validators.NewDeviceTypeCreateFormValidator(
		&dtcs.Args.Input,
		repository.NewDeviceTypeRepository(dtcs.Ctx, dtcs.Db),
		&deviceType,
	)

	if err := form.Save(); err != nil {
		return nil, err
	}

	return &deviceType, nil
}
