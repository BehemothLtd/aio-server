package validators

import (
	"aio-server/enums"
	"aio-server/exceptions"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/constants"
	"aio-server/pkg/helpers"
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
	form.AddAttributes(
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "name",
			},
			Value: helpers.GetStringOrDefault(&form.Name),
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "code",
			},
			Value: helpers.GetStringOrDefault(&form.Code),
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "state",
			},
			Value: helpers.GetStringOrDefault(&form.State),
		},
		&IntAttribute[int32]{
			FieldAttribute: FieldAttribute{
				Code: "userId",
			},
			Value: helpers.GetInt32OrDefault(form.UserId),
		},
		&IntAttribute[int32]{
			FieldAttribute: FieldAttribute{
				Code: "deviceTypeId",
			},
			Value: helpers.GetInt32OrDefault(&form.Device.DeviceTypeId),
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "description",
			},
			Value: helpers.GetStringOrDefault(form.Description),
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "seller",
			},
			Value: helpers.GetStringOrDefault(form.Seller),
		},
	)
}

func (form *DeviceForm) validateName() *DeviceForm {
	nameField := form.FindAttrByCode("name")

	nameField.ValidateRequired()

	nameField.ValidateMin(interface{}(int64(2)))
	nameField.ValidateMax(interface{}(int64(constants.MaxStringLength)))

	if nameField.IsClean() {
		form.Device.Name = form.Name
	}

	return form
}

func (form *DeviceForm) validateCode() *DeviceForm {
	codeField := form.FindAttrByCode("code")

	codeField.ValidateRequired()

	codeField.ValidateMin(interface{}(int64(2)))
	codeField.ValidateMax(interface{}(int64(constants.MaxStringLength)))

	if codeField.IsClean() {
		form.Device.Code = form.Code
	}

	return form
}

func (form *DeviceForm) validateState() *DeviceForm {
	state := form.FindAttrByCode("state")
	state.ValidateRequired()

	if state.IsClean() {
		if stateEnum, err := enums.ParseDeviceStateType(form.State); err != nil {
			state.AddError("is invalid")
		} else {
			form.Device.State = stateEnum
		}
	}

	return form
}

func (form *DeviceForm) validateUserId() *DeviceForm {
	field := form.FindAttrByCode("userId")
	field.ValidateRequired()

	if field.IsClean() {
		form.Device.UserId = *form.UserId
	}

	return form
}

func (form *DeviceForm) validateDeviceTypeId() *DeviceForm {
	field := form.FindAttrByCode("deviceTypeId")
	field.ValidateRequired()

	if field.IsClean() {
		form.Device.DeviceTypeId = form.DeviceTypeId
	}

	return form
}

func (form *DeviceForm) validate() error {
	form.validateName().
		validateCode().
		validateState().
		validateUserId().
		validateDeviceTypeId().
		summaryErrors()

	if form.Errors != nil {
		return exceptions.NewUnprocessableContentError("", form.Errors)
	}

	return nil
}

func (form *DeviceForm) Save() error {
	if err := form.validate(); err != nil {
		return err
	}

	return form.Repo.Create(form.Device)
}
