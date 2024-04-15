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

type AttendanceForm struct {
	Form
	insightInputs.AttendanceFormInput
	Attendance       *models.Attendance
	Repo             *repository.AttendanceRepository
	UpdateAttendance models.Attendance
}

func NewAttendanceFormValidator(
	input *insightInputs.AttendanceFormInput,
	repo *repository.AttendanceRepository,
	Attendance *models.Attendance,
) AttendanceForm {
	form := AttendanceForm{
		Form:                Form{},
		AttendanceFormInput: *input,
		Attendance:          Attendance,
		UpdateAttendance:    models.Attendance{},
		Repo:                repo,
	}

	form.assignAttributes()

	return form

}

func (form *AttendanceForm) assignAttributes() {
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

func (form *AttendanceForm) Save() error {
	if err := form.validate(); err != nil {
		return err
	}

	if form.Attendance.Id == 0 {
		if err := form.Repo.Create(form.Attendance); err != nil {
			return exceptions.NewUnprocessableContentError("", exceptions.ResourceModificationError{
				"base": {err.Error()},
			})
		}
	} else {
		if err := form.Repo.Update(form.Attendance, form.UpdateAttendance); err != nil {
			return exceptions.NewUnprocessableContentError("", exceptions.ResourceModificationError{
				"base": {err.Error()},
			})
		}
	}

	return nil
}

func (form *AttendanceForm) validate() error {
	form.validateCheckinAt().validateCheckOutAt().validateUserId()

	form.summaryErrors()

	if form.Errors != nil {
		return exceptions.NewUnprocessableContentError("", form.Errors)
	}
	return nil
}

func (form *AttendanceForm) validateCheckinAt() *AttendanceForm {
	checkinAt := form.FindAttrByCode("checkinAt")
	checkinAt.ValidateRequired()
	checkinAt.ValidateFormat(constants.DDMMYYY_HHMM_DateFormat, constants.HUMAN_DDMMYYY_HHMM_DateFormat)

	repo := repository.NewAttendanceRepository(nil, database.Db)
	var count int64
	var checkinAtTime = *checkinAt.Time()

	if checkinAt.IsClean() {
		var err error
		if form.Attendance.Id == 0 {
			err = repo.CountAtDateOfUser(&count, form.UserId, checkinAtTime, nil)
			form.Attendance.CheckinAt = checkinAt.Time()
		} else {
			err = repo.CountAtDateOfUser(&count, form.UserId, checkinAtTime, &form.Attendance.Id)
			form.UpdateAttendance.CheckinAt = checkinAt.Time()
		}

		if err != nil {
			checkinAt.AddError(err)
		}
		// handle validate error
		if count > 0 {
			checkinAt.AddError("user already checkin at this day")
		}
	}
	return form
}

func (form *AttendanceForm) validateCheckOutAt() *AttendanceForm {
	checkoutAt := form.FindAttrByCode("checkoutAt")
	checkoutAt.ValidateRequired()
	checkoutAt.ValidateFormat(constants.DDMMYYY_HHMM_DateFormat, constants.HUMAN_DDMMYYY_HHMM_DateFormat)

	checkinAt := form.FindAttrByCode("checkinAt")

	var checkinAtTime = *checkinAt.Time()
	checkoutAt.ValidateMin(interface{}(checkinAtTime))

	endOfDay := helpers.EndOfDay(&checkinAtTime)
	checkoutAt.ValidateMax(interface{}(endOfDay))

	if checkoutAt.IsClean() {
		if form.Attendance.Id == 0 {
			form.Attendance.CheckoutAt = checkoutAt.Time()
		} else {
			form.UpdateAttendance.CheckoutAt = checkoutAt.Time()
		}
	}

	return form
}

func (form *AttendanceForm) validateUserId() *AttendanceForm {
	field := form.FindAttrByCode("userId")
	field.ValidateRequired()

	userRepo := repository.NewUserRepository(nil, database.Db)
	if err := userRepo.Find(&models.User{Id: form.UserId}); err != nil {
		field.AddError("Invalid User")
	}

	if field.IsClean() {
		if form.Attendance.Id == 0 {
			form.Attendance.UserId = form.UserId
		} else {
			form.UpdateAttendance.UserId = form.UserId
		}
	}

	return form
}
