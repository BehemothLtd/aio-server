package insightServices

import (
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/repository"
	"aio-server/validators"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type DeviceUsingHistoryCreateService struct {
	Ctx                *context.Context
	Db                 *gorm.DB
	Args               insightInputs.DeviceUsingHistoryCreateInput
	DeviceUsingHistory *models.DevicesUsingHistory
}

func (s *DeviceUsingHistoryCreateService) Execute() error {
	form := validators.NewDeviceUsingHistoryFormValidator(
		&s.Args.Input,
		repository.NewDeviceUsingHistoryRepository(s.Ctx, s.Db),
		s.DeviceUsingHistory,
	)

	if err := form.Save(); err != nil {
		return err
	}

	return nil
}
