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

type DeviceForm struct {
	Form
	insightInputs.DeviceCreateFormInput
	Device *models.Device
	Repo   *repository.DeviceRepository
}

func NewDeviceFormValidator(
	input *insightInputs.DeviceCreateFormInput,
	repo *repository.DeviceRepository,
	device *models.Device,
) DeviceForm {
	form := DeviceForm{
		Form:                  Form{},
		DeviceCreateFormInput: *input,
		Device:                device,
		Repo:                  repo,
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

func (form *DeviceForm) validateName() *DeviceForm {
	nameField := form.FindAttrByCode("name")

	nameField.ValidateRequired()

	if nameField.IsClean() {
		form.Device.Name = *form.Name
	}

	return form
}

func (form *DeviceForm) validateCode() *DeviceForm {
	codeField := form.FindAttrByCode("code")

	codeField.ValidateRequired()

	if form.Code != nil && strings.TrimSpace(*form.Code) != "" {
		device := models.Device{Code: *form.Code}

		if err := form.Repo.Find(&device); err == nil {
			if form.Device.Id == 0 || device.Id != form.Device.Id {
				codeField.AddError("is already exists. Please use another code")
			}
		}

		if codeField.IsClean() {
			form.Device.Code = *form.Code
		}
	}

	return form
}

func (form *DeviceForm) validateDescription() *DeviceForm {
	descField := form.FindAttrByCode("description")

	descField.ValidateMax(interface{}(int64(constants.MaxLongTextLength)))

	if descField.IsClean() && form.Description != nil {
		form.Device.Description = *form.Description
	}

	return form
}

func (form *DeviceForm) validateSeller() *DeviceForm {
	sellerField := form.FindAttrByCode("seller")

	sellerField.ValidateMax(interface{}(int64(constants.MaxLongTextLength)))

	if sellerField.IsClean() && form.Seller != nil {
		form.Device.Seller = *form.Seller
	}

	return form
}

func (form *DeviceForm) validateState() *DeviceForm {
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

func (form *DeviceForm) validateUserId() *DeviceForm {
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

func (form *DeviceForm) validateDeviceTypeId() *DeviceForm {
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

func (form *DeviceForm) validate() error {
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

func (form *DeviceForm) Save() error {
	if err := form.validate(); err != nil {
		return err
	}

	if form.Device.Id == 0 {
		return form.Repo.Create(form.Device)
	}

	return form.Repo.Update(form.Device)
}
