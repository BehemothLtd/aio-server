package validators

import (
	"aio-server/database"
	"aio-server/enums"
	"aio-server/exceptions"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/repository"
)

type DevicesUsingHistoryForm struct {
	Form
	insightInputs.DevicesUsingHistoryCreateFormInput
	DevicesUsingHistory *models.DevicesUsingHistory
	Repo                *repository.DevicesUsingHistoryRepository
}

func NewDevicesUsingHistoryFormValidator(
	input *insightInputs.DevicesUsingHistoryCreateFormInput,
	repo *repository.DevicesUsingHistoryRepository,
	devicesUsingHistory *models.DevicesUsingHistory,
) DevicesUsingHistoryForm {
	form := DevicesUsingHistoryForm{
		Form:                               Form{},
		DevicesUsingHistoryCreateFormInput: *input,
		DevicesUsingHistory:                devicesUsingHistory,
		Repo:                               repo,
	}

	form.assignAttributes()

	return form
}

func (form *DevicesUsingHistoryForm) assignAttributes() {
	form.AddAttributes(
		&IntAttribute[int32]{
			FieldAttribute: FieldAttribute{
				Code: "userId",
			},
			Value: helpers.GetInt32OrDefault(form.UserId),
		},
		&IntAttribute[int32]{
			FieldAttribute: FieldAttribute{
				Code: "deviceId",
			},
			Value: helpers.GetInt32OrDefault(form.DeviceId),
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "state",
			},
			Value: helpers.GetStringOrDefault(form.State),
		},
	)
}

func (form *DevicesUsingHistoryForm) validateUserId() *DevicesUsingHistoryForm {
	userIdField := form.FindAttrByCode("userId")

	if form.UserId != nil {
		userRepo := repository.NewUserRepository(nil, database.Db)

		if userIdField.IsClean() {
			if err := userRepo.Find(&models.User{Id: *form.UserId}); err != nil {
				userIdField.AddError("is invalid")
			}

			form.DevicesUsingHistory.UserId = *form.UserId
		}
	}

	return form
}

func (form *DevicesUsingHistoryForm) validateDeviceId() *DevicesUsingHistoryForm {
	deviceIdField := form.FindAttrByCode("deviceId")

	if form.DeviceId != nil {
		deviceRepo := repository.NewDeviceRepository(nil, database.Db)

		if deviceIdField.IsClean() {
			if err := deviceRepo.Find(&models.Device{Id: *form.DeviceId}); err != nil {
				deviceIdField.AddError("is invalid")
			}

			form.DevicesUsingHistory.DeviceId = *form.DeviceId
		}
	}

	return form
}

func (form *DevicesUsingHistoryForm) validateState() *DevicesUsingHistoryForm {
	stateField := form.FindAttrByCode("state")
	userIdField := form.FindAttrByCode("userId")

	stateField.ValidateRequired()

	if form.State != nil {
		fieldValue := enums.DeviceStateType(*form.State)

		if !fieldValue.IsValid() {
			stateField.AddError("is invalid")
		}

		if stateField.IsClean() {
			form.DevicesUsingHistory.State = fieldValue
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

func (form *DevicesUsingHistoryForm) validate() error {
	form.validateUserId().
		validateDeviceId().
		validateState().
		summaryErrors()

	if form.Errors != nil {
		return exceptions.NewUnprocessableContentError("", form.Errors)
	}

	return nil
}

func (form *DevicesUsingHistoryForm) Save() error {
	if err := form.validate(); err != nil {
		return err
	}

	return form.Repo.Create(form.DevicesUsingHistory)
}
