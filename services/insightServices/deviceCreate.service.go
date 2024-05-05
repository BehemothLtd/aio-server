package insightServices

import (
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/repository"
	"aio-server/validators"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type DeviceCreateService struct {
	Ctx    *context.Context
	Db     *gorm.DB
	Args   insightInputs.DeviceCreateInput
	Device *models.Device
}

func (dcs *DeviceCreateService) Execute() error {
	form := validators.NewDeviceCreateFormValidator(
		&dcs.Args.Input,
		repository.NewDeviceRepository(dcs.Ctx, dcs.Db),
		dcs.Device,
	)

	if err := form.Save(); err != nil {
		return err
	}

	return nil
}
