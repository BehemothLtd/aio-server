package insightServices

import (
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"

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

}
