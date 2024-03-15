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

// TODO: handle assign attribute : from, to, time_off
func (form *LeaveDayRequestForm) assignAttributes() {
	form.AddAttributes(
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "name",
			},
			Value: helpers.GetStringOrDefault(form.Reason),
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "requestType",
			},
			Value: helpers.GetStringOrDefault(form.RequestType),
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "requestState",
			},
			Value: helpers.GetStringOrDefault(form.RequestState),
		},
	)
}

func (form *LeaveDayRequestForm) validate() error {
	form.validateReason().
		validateRequestType().
		validateRequestState().
		summaryErrors()

	if form.Errors != nil {
		return exceptions.NewUnprocessableContentError("", form.Errors)
	}

	return nil
}

func (form *LeaveDayRequestForm) validateReason() *LeaveDayRequestForm {
	reasonField := form.FindAttrByCode("reason")

	min := 3
	max := int64(constants.MaxStringLength)
	reasonField.ValidateLimit(&min, &max)

	if reasonField.IsClean() {
		form.Request.Reason = *form.Reason
	}

	return form
}

func (form *LeaveDayRequestForm) validateRequestType() *LeaveDayRequestForm {
	requestTypeField := form.FindAttrByCode("requestType")

	requestTypeField.ValidateRequired()

	if form.RequestType != nil {
		fieldValue := enums.RequestType(*form.RequestType)

		if !fieldValue.IsValid() {
			requestTypeField.AddError("is invalid")
		}

		if requestTypeField.IsClean() {
			form.Request.RequestType = fieldValue
		}
	}

	return form
}

func (form *LeaveDayRequestForm) validateRequestState() *LeaveDayRequestForm {
	requestStateField := form.FindAttrByCode("requestState")

	requestStateField.ValidateRequired()

	if form.RequestState != nil {
		fieldValue := enums.RequestStateType(*form.RequestState)

		if !fieldValue.IsValid() {
			requestStateField.AddError("is invalid")
		}

		if requestStateField.IsClean() {
			form.Request.RequestState = fieldValue
		}
	}

	return form
}
