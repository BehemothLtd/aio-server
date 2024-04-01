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

type IssueStatusForm struct {
	Form
	insightInputs.IssueStatusFormInput
	IssueStatus *models.IssueStatus
	Repo        *repository.IssueStatusRepository
}

func NewIssueStatusFormValidator(
	input *insightInputs.IssueStatusFormInput,
	repo *repository.IssueStatusRepository,
	issueStatus *models.IssueStatus,
) IssueStatusForm {
	form := IssueStatusForm{
		Form:                 Form{},
		IssueStatusFormInput: *input,
		IssueStatus:          issueStatus,
		Repo:                 repo,
	}
	form.assignAttributes()

	return form
}

func (form *IssueStatusForm) Save() error {
	if err := form.validate(); err != nil {
		return err
	}

	if form.IssueStatus.Id == 0 {
		return form.Repo.Create(form.IssueStatus)
	}

	return form.Repo.Update(form.IssueStatus)
}

func (form *IssueStatusForm) assignAttributes() {
	form.AddAttributes(
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "color",
			},
			Value: helpers.GetStringOrDefault(form.Color),
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "title",
			},
			Value: helpers.GetStringOrDefault(form.Title),
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "statusType",
			},
			Value: helpers.GetStringOrDefault(form.StatusType),
		},
		&IntAttribute[int32]{
			FieldAttribute: FieldAttribute{
				Code: "lockVersion",
			},
			Value: helpers.GetInt32OrDefault(form.LockVersion),
		},
	)
}

func (form *IssueStatusForm) validate() error {
	form.validateColor().
		validateTitle().
		validateStatusType()

	if form.IssueStatus.Id != 0 {
		form.validateLockVersion()
	}

	form.summaryErrors()

	if form.Errors != nil {
		return exceptions.NewUnprocessableContentError("", form.Errors)
	}
	return nil
}

func (form *IssueStatusForm) validateColor() *IssueStatusForm {
	colorField := form.FindAttrByCode("color")

	colorField.ValidateRequired()

	if colorField.IsClean() {
		form.IssueStatus.Color = *form.Color
	}

	return form
}

func (form *IssueStatusForm) validateTitle() *IssueStatusForm {
	titleField := form.FindAttrByCode("title")

	titleField.ValidateRequired()

	titleField.ValidateMin(interface{}(int64(2)))
	titleField.ValidateMax(interface{}(int64(constants.MaxStringLength)))

	if titleField.IsClean() {
		form.IssueStatus.Title = *form.Title
	}

	return form
}

func (form *IssueStatusForm) validateStatusType() *IssueStatusForm {
	typeField := form.FindAttrByCode("statusType")

	typeField.ValidateRequired()

	if typeField.IsClean() {
		if statusType, err := enums.ParseIssueStatusStatusType(*form.StatusType); err != nil {
			typeField.AddError("is invalid issue status type")
		} else {
			form.IssueStatus.StatusType = statusType
		}
	}

	return form
}

func (form *IssueStatusForm) validateLockVersion() *IssueStatusForm {
	currentLockVersion := form.IssueStatus.LockVersion

	field := form.FindAttrByCode("lockVersion")

	field.ValidateRequired()

	if field.IsClean() {
		field.ValidateMin(interface{}(int64(currentLockVersion)))
	}

	return form
}
