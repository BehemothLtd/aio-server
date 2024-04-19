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
	Attendance  *models.Attendance
	Repo        *repository.AttendanceRepository
	UpdatesForm map[string]interface{}
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
		UpdatesForm:         map[string]interface{}{},
		Repo:                repo,
	}

	form.assignAttributes()

	return form

}

func (form *AttendanceForm) assignAttributes() {
	form.AddAttributes(
		&TimeAttribute{
			FieldAttribute: FieldAttribute{
				Code: "CheckinAt",
			},
			Value: helpers.GetStringOrDefault(&form.CheckinAt),
		},
		&TimeAttribute{
			FieldAttribute: FieldAttribute{
				Code: "CheckoutAt",
			},
			Value: helpers.GetStringOrDefault(&form.CheckoutAt),
		},
		&IntAttribute[int32]{
			FieldAttribute: FieldAttribute{
				Code: "UserId",
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
		if err := form.Repo.Update(form.Attendance, form.UpdatesForm); err != nil {
			return exceptions.NewUnprocessableContentError("", exceptions.ResourceModificationError{
				"base": {err.Error()},
			})
		}
	}

	return nil
}

func (form *AttendanceForm) validate() error {
	form.validateCheckinAt().validateUserId()
	form.summaryErrors()

	if form.Errors == nil {
		form.validateCheckOutAt()
		form.summaryErrors()
	}

	if form.Errors != nil {
		return exceptions.NewUnprocessableContentError("", form.Errors)
	}
	return nil
}

func (form *AttendanceForm) validateCheckinAt() *AttendanceForm {
	fieldCode := "CheckinAt"
	checkinAt := form.FindAttrByCode(fieldCode)
	checkinAt.ValidateRequired()
	checkinAt.ValidateFormat(constants.DDMMYYY_HHMM_DateFormat, constants.HUMAN_DDMMYYY_HHMM_DateFormat)

	repo := repository.NewAttendanceRepository(nil, database.Db)
	var count int64

	if checkinAt.IsClean() {
		var checkinAtTime = *checkinAt.Time()
		var err error
		if form.Attendance.Id == 0 {
			err = repo.CountAtDateOfUser(&count, form.UserId, checkinAtTime, nil)
			form.Attendance.CheckinAt = checkinAt.Time()
		} else {
			err = repo.CountAtDateOfUser(&count, form.UserId, checkinAtTime, &form.Attendance.Id)
			form.UpdatesForm[fieldCode] = checkinAt.Time()
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
	fieldCode := "CheckoutAt"
	checkoutAt := form.FindAttrByCode(fieldCode)

	if form.CheckoutAt == "" {
		if form.Attendance.Id == 0 {
			form.Attendance.CheckoutAt = nil
		} else {
			form.UpdatesForm[fieldCode] = nil
		}
		return form
	} else {
		checkoutAt.ValidateFormat(constants.DDMMYYY_HHMM_DateFormat, constants.HUMAN_DDMMYYY_HHMM_DateFormat)
		checkinAt := form.FindAttrByCode("CheckinAt")
		var checkinAtTime = *checkinAt.Time()

		checkoutAt.ValidateMin(interface{}(checkinAtTime))
		endOfDay := helpers.EndOfDay(&checkinAtTime)
		checkoutAt.ValidateMax(interface{}(endOfDay))

		if checkoutAt.IsClean() {
			if form.Attendance.Id == 0 {
				form.Attendance.CheckoutAt = checkoutAt.Time()
			} else {
				form.UpdatesForm[fieldCode] = checkoutAt.Time()
			}
		}
		return form
	}
}

func (form *AttendanceForm) validateUserId() *AttendanceForm {
	fieldCode := "UserId"
	field := form.FindAttrByCode(fieldCode)
	field.ValidateRequired()

	userRepo := repository.NewUserRepository(nil, database.Db)
	if err := userRepo.Find(&models.User{Id: form.UserId}); err != nil {
		field.AddError("Invalid User")
	}

	if field.IsClean() {
		if form.Attendance.Id == 0 {
			form.Attendance.UserId = form.UserId
		} else {
			form.UpdatesForm[fieldCode] = form.UserId
		}
	}

	return form
}
