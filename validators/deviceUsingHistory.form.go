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

type DeviceUsingHistoryForm struct {
	Form
	insightInputs.DeviceUsingHistoryCreateFormInput
	DeviceUsingHistory *models.DevicesUsingHistory
	Repo               *repository.DeviceUsingHistoryRepository
}

func NewDeviceUsingHistoryFormValidator(
	input *insightInputs.DeviceUsingHistoryCreateFormInput,
	repo *repository.DeviceUsingHistoryRepository,
	deviceUsingHistory *models.DevicesUsingHistory,
) DeviceUsingHistoryForm {
	form := DeviceUsingHistoryForm{
		Form:                              Form{},
		DeviceUsingHistoryCreateFormInput: *input,
		DeviceUsingHistory:                deviceUsingHistory,
		Repo:                              repo,
	}

	form.assignAttributes()

	return form
}

func (form *DeviceUsingHistoryForm) assignAttributes() {
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

func (form *DeviceUsingHistoryForm) validateUserId() *DeviceUsingHistoryForm {
	userIdField := form.FindAttrByCode("userId")

	if form.UserId != nil {
		userRepo := repository.NewUserRepository(nil, database.Db)

		if userIdField.IsClean() {
			if err := userRepo.Find(&models.User{Id: *form.UserId}); err != nil {
				userIdField.AddError("is invalid")
			}

			form.DeviceUsingHistory.UserId = *form.UserId
		}
	}

	return form
}

func (form *DeviceUsingHistoryForm) validateDeviceId() *DeviceUsingHistoryForm {
	deviceIdField := form.FindAttrByCode("deviceId")

	if form.DeviceId != nil {
		deviceRepo := repository.NewDeviceRepository(nil, database.Db)

		if deviceIdField.IsClean() {
			if err := deviceRepo.Find(&models.Device{Id: *form.DeviceId}); err != nil {
				deviceIdField.AddError("is invalid")
			}

			form.DeviceUsingHistory.DeviceId = *form.DeviceId
		}
	}

	return form
}

func (form *DeviceUsingHistoryForm) validateState() *DeviceUsingHistoryForm {
	stateField := form.FindAttrByCode("state")
	userIdField := form.FindAttrByCode("userId")

	stateField.ValidateRequired()

	if form.State != nil {
		fieldValue := enums.DeviceStateType(*form.State)

		if !fieldValue.IsValid() {
			stateField.AddError("is invalid")
		}

		if stateField.IsClean() {
			form.DeviceUsingHistory.State = fieldValue
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

func (form *DeviceUsingHistoryForm) validate() error {
	form.validateUserId().
		validateDeviceId().
		validateState().
		summaryErrors()

	if form.Errors != nil {
		return exceptions.NewUnprocessableContentError("", form.Errors)
	}

	return nil
}

func (form *DeviceUsingHistoryForm) Save() error {
	if err := form.validate(); err != nil {
		return err
	}

	return form.Repo.Create(form.DeviceUsingHistory)
}
