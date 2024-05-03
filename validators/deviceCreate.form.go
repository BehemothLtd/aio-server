package validators

import (
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/repository"
)

type DeviceCreateForm struct {
	Form
	insightInputs.DeviceCreateFormInput
	Device *models.Device
	Repo   *repository.DeviceRepository
}

func NewDeviceCreateFormValidator(
	input *insightInputs.DeviceCreateFormInput,
	repo *repository.DeviceRepository,
	device *models.Device,
) DeviceCreateForm {
	form := DeviceCreateForm{
		Form:                  Form{},
		DeviceCreateFormInput: *input,
		Device:                device,
		Repo:                  repo,
	}

	form.assignAttributes()

	return form
}
