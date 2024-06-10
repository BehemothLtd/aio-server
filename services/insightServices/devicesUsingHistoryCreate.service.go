package insightServices

import (
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/repository"
	"aio-server/validators"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type DevicesUsingHistoryCreateService struct {
	Ctx                 *context.Context
	Db                  *gorm.DB
	Args                insightInputs.DevicesUsingHistoryCreateInput
	DevicesUsingHistory *models.DevicesUsingHistory
}

func (s *DevicesUsingHistoryCreateService) Execute() error {
	form := validators.NewDevicesUsingHistoryFormValidator(
		&s.Args.Input,
		repository.NewDevicesUsingHistoryRepository(s.Ctx, s.Db),
		s.DevicesUsingHistory,
	)

	if err := form.Save(); err != nil {
		return err
	}

	return nil
}
