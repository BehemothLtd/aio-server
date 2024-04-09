package validators

import (
	"aio-server/database"
	"aio-server/exceptions"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/constants"
	"aio-server/pkg/helpers"
	"aio-server/repository"
)

type AttendanceCreateForm struct {
	Form
	insightInputs.AttendanceFormInput
	Attendance *models.Attendance
	Repo       *repository.AttendanceRepository
}

func NewAttendanceCreateFormValidator(
	input *insightInputs.AttendanceFormInput,
	repo *repository.AttendanceRepository,
	Attendance *models.Attendance,
) AttendanceCreateForm {
	form := AttendanceCreateForm{
		Form:                Form{},
		AttendanceFormInput: *input,
		Attendance:          Attendance,
		Repo:                repo,
	}

	form.assignAttributes()

	return form

}

func (form *AttendanceCreateForm) assignAttributes() {
	form.AddAttributes(
		&TimeAttribute{
			FieldAttribute: FieldAttribute{
				Code: "checkinAt",
				Name: "checkin_at",
			},
			Value: helpers.GetStringOrDefault(&form.CheckinAt),
		},
		&TimeAttribute{
			FieldAttribute: FieldAttribute{
				Code: "checkoutAt",
				Name: "checkout_at",
			},
			Value: helpers.GetStringOrDefault(&form.CheckoutAt),
		},
		&IntAttribute[int32]{
			FieldAttribute: FieldAttribute{
				Code: "userId",
				Name: "user_id",
			},
			Value: helpers.GetInt32OrDefault(&form.UserId),
		},
	)
}

func (form *AttendanceCreateForm) Save() error {
	if err := form.validate(); err != nil {
		return err
	}

	return form.Repo.Create(form.Attendance)
}

func (form *AttendanceCreateForm) validate() error {
	form.validateCheckinAt().validateCheckOutAt().validateUserId()

	form.summaryErrors()

	if form.Errors != nil {
		return exceptions.NewUnprocessableContentError("", form.Errors)
	}
	return nil
}

func (form *AttendanceCreateForm) validateCheckinAt() *AttendanceCreateForm {
	checkinAt := form.FindAttrByCode("checkinAt")
	checkinAt.ValidateRequired()
	checkinAt.ValidateFormat(constants.DDMMYYY_HHMM_DateFormat, constants.HUMAN_DDMMYYY_HHMM_DateFormat)
	var checkinAtTime = *checkinAt.Time()

	repo := repository.NewAttendanceRepository(nil, database.Db)
	var count int64

	if checkinAt.IsClean() {
		if err := repo.CountAtDateOfUser(&count, form.UserId, checkinAtTime); err != nil {
			checkinAt.AddError(err)
		} else if count > 0 {
			checkinAt.AddError("user already checkin at this day")
		} else {
			form.Attendance.CheckinAt = checkinAt.Time()
		}
	}
	return form
}

func (form *AttendanceCreateForm) validateCheckOutAt() *AttendanceCreateForm {
	checkoutAt := form.FindAttrByCode("checkoutAt")
	checkoutAt.ValidateRequired()
	checkoutAt.ValidateFormat(constants.DDMMYYY_HHMM_DateFormat, constants.HUMAN_DDMMYYY_HHMM_DateFormat)

	checkinAt := form.FindAttrByCode("checkinAt")

	var checkinAtTime = *checkinAt.Time()
	checkoutAt.ValidateMin(interface{}(checkinAtTime))

	endOfDay := helpers.EndOfDay(&checkinAtTime)
	checkoutAt.ValidateMax(interface{}(endOfDay))

	if checkoutAt.IsClean() {
		form.Attendance.CheckoutAt = checkoutAt.Time()
	}

	return form
}

func (form *AttendanceCreateForm) validateUserId() *AttendanceCreateForm {
	field := form.FindAttrByCode("userId")
	field.ValidateRequired()

	userRepo := repository.NewUserRepository(nil, database.Db)
	if err := userRepo.Find(&models.User{Id: form.UserId}); err != nil {
		field.AddError("Invalid User")
	}

	if field.IsClean() {
		form.Attendance.UserId = form.UserId
	}

	return form
}
