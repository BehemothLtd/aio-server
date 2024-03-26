package validators

import (
	"aio-server/enums"
	"aio-server/exceptions"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/constants"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"fmt"
	"strings"
)

type LeaveDayRequestForm struct {
	Form
	insightInputs.LeaveDayRequestFormInput
	Request *models.LeaveDayRequest
	Repo    *repository.LeaveDayRequestRepository
}

func NewLeaveDayrequestFormValidator(
	input *insightInputs.LeaveDayRequestFormInput,
	repo *repository.LeaveDayRequestRepository,
	request *models.LeaveDayRequest,
) LeaveDayRequestForm {
	form := LeaveDayRequestForm{
		Form:                     Form{},
		LeaveDayRequestFormInput: *input,
		Request:                  request,
		Repo:                     repo,
	}

	form.assignAttributes()

	return form
}

func (form *LeaveDayRequestForm) Save() error {
	if err := form.validate(); err != nil {
		return err
	}

	if err := form.Repo.Create(form.Request); err != nil {
		return err
	}
	return nil
}

func (form *LeaveDayRequestForm) assignAttributes() {
	form.AddAttributes(
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "reason",
			},
			Value: helpers.GetStringOrDefault(form.Reason),
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "requestType",
			},
			Value: helpers.GetStringOrDefault(&form.RequestType),
		},
		&TimeAttribute{
			FieldAttribute: FieldAttribute{
				Code: "from",
			},
			Value: helpers.GetStringOrDefault(&form.From),
		},
		&TimeAttribute{
			FieldAttribute: FieldAttribute{
				Code: "to",
			},
			Value: helpers.GetStringOrDefault(&form.To),
		},
		&FloatAttribute[float64]{
			FieldAttribute: FieldAttribute{
				Code: "timeOff",
			},
			Value: helpers.GetFloat64OrDefault(&form.TimeOff),
		},
	)
}

func (form *LeaveDayRequestForm) validate() error {
	form.validateReason().
		validateRequestType().
		validateTimeOff().
		validateFrom().
		validateTo().
		summaryErrors()

	if form.Errors != nil {
		return exceptions.NewUnprocessableContentError("", form.Errors)
	}

	return nil
}

func (form *LeaveDayRequestForm) validateReason() *LeaveDayRequestForm {
	reasonField := form.FindAttrByCode("reason")

	if strings.TrimSpace(*form.Reason) != "" {
		reasonField.ValidateMin(interface{}(int64(3)))
		reasonField.ValidateMax(interface{}(int64(constants.MaxStringLength)))
	}

	if reasonField.IsClean() {
		form.Request.Reason = form.Reason
	}

	return form
}

func (form *LeaveDayRequestForm) validateRequestType() *LeaveDayRequestForm {
	requestTypeField := form.FindAttrByCode("requestType")

	requestTypeField.ValidateRequired()

	if requestTypeField.IsClean() {
		fieldValue := enums.RequestType(form.RequestType)

		if !fieldValue.IsValid() {
			requestTypeField.AddError("is invalid")
		}

		if requestTypeField.IsClean() {
			form.Request.RequestType = fieldValue
		}
	}

	return form
}

func (form *LeaveDayRequestForm) validateTimeOff() *LeaveDayRequestForm {
	timeOff := form.FindAttrByCode("timeOff")

	timeOff.ValidateRequired()
	timeOff.ValidateMin(interface{}(float64(0)))
	form.Request.TimeOff = form.TimeOff

	return form
}

func (form *LeaveDayRequestForm) validateFrom() *LeaveDayRequestForm {
	field := form.FindAttrByCode("from")
	field.ValidateRequired()
	field.ValidateFormat(constants.DDMMYYY_HHMM_DateFormat, "%d-%m-%y %H:%M")

	beginningOfDay := helpers.BeginningOfDay(nil)

	if field.Time().Before(beginningOfDay) {
		field.AddError(fmt.Sprintf("is invalid, must be later than start of the day %+v",
			beginningOfDay.Format(constants.DDMMYYY_DateFormat)),
		)
	}

	form.Request.From = *field.Time()
	return form
}

func (form *LeaveDayRequestForm) validateTo() *LeaveDayRequestForm {
	toTime := form.FindAttrByCode("to")
	fromTime := form.FindAttrByCode("from")

	toTime.ValidateRequired()
	toTime.ValidateFormat(constants.DDMMYYY_HHMM_DateFormat, "%d-%m-%y %H:%M")

	if toTime.Time().Before(*fromTime.Time()) {
		toTime.AddError(fmt.Sprintf("is invalid, must be later than start time %+v",
			fromTime.Time().Format(constants.DDMMYYY_HHMM_DateFormat)),
		)
	}

	form.Request.To = *toTime.Time()
	return form
}
