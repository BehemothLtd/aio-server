package validators

import (
	"aio-server/enums"
	"aio-server/exceptions"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/constants"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"strings"
)

type IssueStatusCreateForm struct {
	Form
	insightInputs.IssueStatusCreateFormInput
	IssueStatus *models.IssueStatus
	Repo        *repository.IssueStatusRepository
}

func NewIssueStatusCreateFormValidator(
	input *insightInputs.IssueStatusCreateFormInput,
	repo *repository.IssueStatusRepository,
	issueStatus *models.IssueStatus,
) IssueStatusCreateForm {
	form := IssueStatusCreateForm{
		Form:                       Form{},
		IssueStatusCreateFormInput: *input,
		IssueStatus:                issueStatus,
		Repo:                       repo,
	}
	form.assignAttributes()

	return form
}

func (form *IssueStatusCreateForm) Save() error {
	if err := form.validate(); err != nil {
		return err
	}

	if err := form.Repo.Create(form.IssueStatus); err != nil {
		return err
	}

	return nil
}

func (form *IssueStatusCreateForm) assignAttributes() {
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
	)
}

func (form *IssueStatusCreateForm) validate() error {
	form.validateColor().
		validateTitle().
		validateStatusType().
		summaryErrors()

	if form.Errors != nil {
		return exceptions.NewUnprocessableContentError("", form.Errors)
	}
	return nil
}

func (form *IssueStatusCreateForm) validateColor() *IssueStatusCreateForm {
	colorField := form.FindAttrByCode("color")

	colorField.ValidateRequired()

	if form.Color != nil && strings.TrimSpace(*form.Color) != "" {
		if colorField.IsClean() {
			form.IssueStatus.Color = *form.Color
		}
	}

	return form
}

func (form *IssueStatusCreateForm) validateTitle() *IssueStatusCreateForm {
	titleField := form.FindAttrByCode("title")

	titleField.ValidateRequired()

	titleField.ValidateMin(interface{}(int64(2)))
	titleField.ValidateMax(interface{}(int64(constants.MaxStringLength)))

	if form.Title != nil && strings.TrimSpace(*form.Title) != "" {

		if titleField.IsClean() {
			form.IssueStatus.Title = *form.Title
		}
	}

	return form
}

func (form *IssueStatusCreateForm) validateStatusType() *IssueStatusCreateForm {
	typeField := form.FindAttrByCode("statusType")

	typeField.ValidateRequired()

	if form.StatusType != nil {
		fieldValue := enums.IssueStatusStatusType(*form.StatusType)

		if !fieldValue.IsValid() {
			typeField.AddError("is invalid")
		}

		if typeField.IsClean() {
			form.IssueStatus.StatusType = fieldValue
		}
	}

	return form
}
