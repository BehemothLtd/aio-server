package validators

import (
	"aio-server/exceptions"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/repository"
)

type ProjectCreateProjectIssueForm struct {
	Form
	insightInputs.ProjectIssueStatusInputForProjectCreate

	ProjectIssueStatus *models.ProjectIssueStatus
	IssueStatusRepo    *repository.IssueStatusRepository
}

func NewProjectCreateProjectIssueFormValidator(
	input *insightInputs.ProjectIssueStatusInputForProjectCreate,
	repo *repository.IssueStatusRepository,
	projectIssueStatus *models.ProjectIssueStatus,
) ProjectCreateProjectIssueForm {
	form := ProjectCreateProjectIssueForm{
		Form:                                    Form{},
		ProjectIssueStatusInputForProjectCreate: *input,
		ProjectIssueStatus:                      projectIssueStatus,
		IssueStatusRepo:                         repo,
	}
	form.assignAttributes()

	return form
}

func (form *ProjectCreateProjectIssueForm) assignAttributes() {
	form.AddAttributes(
		&IntAttribute[int32]{
			FieldAttribute: FieldAttribute{
				Code: "issueStatusId",
			},
			Value: helpers.GetInt32OrDefault(&form.IssueStatusId),
		},
	)
}

func (form *ProjectCreateProjectIssueForm) Validate() exceptions.ResourceModificationError {
	form.validateIssueStatusId().summaryErrors()

	if form.Errors != nil {
		return form.Errors
	}

	return nil
}

func (form *ProjectCreateProjectIssueForm) validateIssueStatusId() *ProjectCreateProjectIssueForm {
	field := form.FindAttrByCode("issueStatusId")
	field.ValidateRequired()

	if form.IssueStatusId != 0 {
		if err := form.IssueStatusRepo.Find(&models.IssueStatus{Id: form.IssueStatusId}); err != nil {
			field.AddError("is invalid")
		}
	}

	if field.IsClean() {
		form.ProjectIssueStatus.IssueStatusId = form.IssueStatusId
	}

	return form
}
