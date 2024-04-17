package validators

import (
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/repository"
)

type DeviceForm struct {
	Form
	insightInputs.DeviceFormInput
	Device *models.Device
	Repo   *repository.DeviceRepository
}

func NewDeviceFormValidator(
	input *insightInputs.DeviceFormInput,
	repo *repository.DeviceRepository,
	device *models.Device,
) DeviceForm {
	form := DeviceForm{
		Form:            Form{},
		DeviceFormInput: *input,
		Device:          device,
		Repo:            repo,
	}

	form.assignAttributes()

	return form
}

func (form *DeviceForm) assignAttributes() {
	form.AddAttributes()
}
