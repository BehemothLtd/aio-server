package validators

import (
	"aio-server/exceptions"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/repository"
)

type DeviceTypeForm struct {
	Form
	insightInputs.DeviceTypeFormInput
	DeviceType *models.DeviceType
	Repo       *repository.DeviceTypeRepository
}

func NewDeviceTypeFormValidation(
	input *insightInputs.DeviceTypeFormInput,
	repo *repository.DeviceTypeRepository,
	deviceType *models.DeviceType,
) DeviceTypeForm {
	form := DeviceTypeForm{
		Form:                Form{},
		DeviceTypeFormInput: *input,
		DeviceType:          deviceType,
		Repo:                repo,
	}
	form.assignAttributes()

	return form
}

func (form *DeviceTypeForm) Save() error {
	if err := form.validate(); err != nil {
		return err
	}

	if form.DeviceType.Id == 0 {
		return form.Repo.Create(form.DeviceType)
	}

	return form.Repo.Update(form.DeviceType)
}

func (form *DeviceTypeForm) assignAttributes() {
	form.AddAttributes(
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "name",
			},
			Value: helpers.GetStringOrDefault(form.Name),
		},
	)
}

func (form *DeviceTypeForm) validate() error {
	form.validateName().summaryErrors()

	if form.Errors != nil {
		return exceptions.NewUnprocessableContentError("", form.Errors)
	}
	return nil
}

func (form *DeviceTypeForm) validateName() *DeviceTypeForm {
	nameField := form.FindAttrByCode("name")

	nameField.ValidateRequired()

	if nameField.IsClean() {
		form.DeviceType.Name = *form.Name
	}

	return form
}
