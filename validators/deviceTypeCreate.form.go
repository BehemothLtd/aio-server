package validators

import (
	"aio-server/exceptions"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"strings"
)

type DeviceTypeCreateForm struct {
	Form
	insightInputs.DeviceTypeCreateFormInput
	DeviceType *models.DeviceType
	Repo       *repository.DeviceTypeRepository
}

func NewDeviceTypeCreateFormValidation(
	input *insightInputs.DeviceTypeCreateFormInput,
	repo *repository.DeviceTypeRepository,
	deviceType *models.DeviceType,
) DeviceTypeCreateForm {
	form := DeviceTypeCreateForm{
		Form:                      Form{},
		DeviceTypeCreateFormInput: *input,
		DeviceType:                deviceType,
		Repo:                      repo,
	}
	form.assignAttributes()

	return form
}

func (form *DeviceTypeCreateForm) Save() error {
	if err := form.validate(); err != nil {
		return err
	}

	if err := form.Repo.Create(form.DeviceType); err != nil {
		return err
	}

	return nil
}

func (form *DeviceTypeCreateForm) assignAttributes() {
	form.AddAttributes(
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "name",
			},
			Value: helpers.GetStringOrDefault(form.Name),
		},
	)
}

func (form *DeviceTypeCreateForm) validate() error {
	form.validateName().summaryErrors()

	if form.Errors != nil {
		return exceptions.NewUnprocessableContentError("", form.Errors)
	}
	return nil
}

func (form *DeviceTypeCreateForm) validateName() *DeviceTypeCreateForm {
	nameField := form.FindAttrByCode("name")

	nameField.ValidateRequired()

	if form.Name != nil && strings.TrimSpace(*form.Name) != "" {
		if nameField.IsClean() {
			form.DeviceType.Name = *form.Name
		}
	}

	return form
}
