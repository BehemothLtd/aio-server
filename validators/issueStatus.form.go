package validators

import (
	"aio-server/exceptions"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
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

func (form *IssueStatusForm) assignAttributes() {
	form.AddAttributes(
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "title",
			},
			Value: helpers.GetStringOrDefault(form.Title),
		},
	)
}

func (form *IssueStatusForm) validate() error {
	title := form.FindAttrByCode("title")

	title.ValidateRequired()

	form.summaryErrors()

	if form.Errors != nil {
		return exceptions.NewUnprocessableContentError("", form.Errors)
	}
	return nil
}

func (form *IssueStatusForm) Save() error {
	form.validate()

	if err := form.validate(); err != nil {
		return err
	}

	// TODO: save db

	return nil
}
