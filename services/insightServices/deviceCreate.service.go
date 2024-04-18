package insightServices

import (
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/repository"
	"aio-server/validators"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type DeviceService struct {
	Ctx    *context.Context
	Db     *gorm.DB
	Args   insightInputs.DeviceCreateInput
	Device *models.Device
}

func (ds *DeviceService) Execute() error {
	form := validators.NewDeviceFormValidator(
		&ds.Args.Input,
		repository.NewDeviceRepository(ds.Ctx, ds.Db),
		ds.Device,
	)

	if err := form.Save(); err != nil {
		return err
	}
	return nil
}
