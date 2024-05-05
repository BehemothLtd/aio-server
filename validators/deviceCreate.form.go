package validators

import (
	"aio-server/database"
	"aio-server/enums"
	"aio-server/exceptions"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/constants"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"strings"
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

func (form *DeviceCreateForm) assignAttributes() {
	form.AddAttributes(
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "name",
			},
			Value: helpers.GetStringOrDefault(form.Name),
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "code",
			},
			Value: helpers.GetStringOrDefault(form.Code),
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "state",
			},
			Value: helpers.GetStringOrDefault(form.State),
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
			Value: helpers.GetInt32OrDefault(form.DeviceTypeId),
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

func (form *DeviceCreateForm) validateName() *DeviceCreateForm {
	nameField := form.FindAttrByCode("name")

	nameField.ValidateRequired()

	if nameField.IsClean() {
		form.Device.Name = *form.Name
	}

	return form
}

func (form *DeviceCreateForm) validateCode() *DeviceCreateForm {
	codeField := form.FindAttrByCode("code")

	codeField.ValidateRequired()

	if form.Code != nil && strings.TrimSpace(*form.Code) != "" {
		device := models.Device{Code: *form.Code}

		if err := form.Repo.Find(&device); err == nil {
			codeField.AddError("is already exists. Please use another code")
		}

		if codeField.IsClean() {
			form.Device.Code = *form.Code
		}
	}

	return form
}

func (form *DeviceCreateForm) validateDescription() *DeviceCreateForm {
	descField := form.FindAttrByCode("description")

	descField.ValidateMax(interface{}(int64(constants.MaxLongTextLength)))

	if descField.IsClean() && form.Description != nil {
		form.Device.Description = *form.Description
	}

	return form
}

func (form *DeviceCreateForm) validateSeller() *DeviceCreateForm {
	sellerField := form.FindAttrByCode("seller")

	sellerField.ValidateMax(interface{}(int64(constants.MaxLongTextLength)))

	if sellerField.IsClean() && form.Seller != nil {
		form.Device.Seller = *form.Seller
	}

	return form
}

func (form *DeviceCreateForm) validateState() *DeviceCreateForm {
	stateField := form.FindAttrByCode("state")
	userIdField := form.FindAttrByCode("userId")

	stateField.ValidateRequired()

	if form.State != nil {
		fieldValue := enums.DeviceStateType(*form.State)

		if !fieldValue.IsValid() {
			stateField.AddError("is invalid")
		}

		if stateField.IsClean() {
			form.Device.State = fieldValue
		}

		if fieldValue == enums.DeviceStateTypeFixing || fieldValue == enums.DeviceStateTypeUsing {
			userIdField.ValidateRequired()
		} else {
			if form.UserId != nil {
				userIdField.AddError("need to be empty")
			}
		}
	}

	return form
}

func (form *DeviceCreateForm) validateUserId() *DeviceCreateForm {
	userIdField := form.FindAttrByCode("userId")

	if form.UserId != nil {
		userRepo := repository.NewUserRepository(nil, database.Db)

		if userIdField.IsClean() {
			if err := userRepo.Find(&models.User{Id: *form.UserId}); err != nil {
				userIdField.AddError("is invalid")
			}

			form.Device.UserId = *form.UserId
		}
	}

	return form
}

func (form *DeviceCreateForm) validateDeviceTypeId() *DeviceCreateForm {
	deviceTypeIdField := form.FindAttrByCode("deviceTypeId")

	deviceTypeIdField.ValidateRequired()

	deviceTypeRepo := repository.NewDeviceTypeRepository(nil, database.Db)

	if deviceTypeIdField.IsClean() {
		if err := deviceTypeRepo.Find(&models.DeviceType{Id: *form.DeviceTypeId}); err != nil {
			deviceTypeIdField.AddError("is invalid")
		}

		form.Device.DeviceTypeId = *form.DeviceTypeId
	}

	return form
}

func (form *DeviceCreateForm) validate() error {
	form.validateName().
		validateCode().
		validateState().
		validateUserId().
		validateDeviceTypeId().
		validateDescription().
		validateSeller().
		summaryErrors()

	if form.Errors != nil {
		return exceptions.NewUnprocessableContentError("", form.Errors)
	}

	return nil
}

func (form *DeviceCreateForm) Save() error {
	if err := form.validate(); err != nil {
		return err
	}

	if err := form.Repo.Create(form.Device); err != nil {
		return err
	}

	return nil
}
