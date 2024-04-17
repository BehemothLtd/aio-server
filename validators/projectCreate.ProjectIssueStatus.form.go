package validators

import (
	"aio-server/exceptions"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/repository"
)

type ProjectCreateProjectIssueStatusForm struct {
	Form
	insightInputs.ProjectIssueStatusInputForProjectCreate

	ProjectIssueStatus *models.ProjectIssueStatus
	IssueStatusRepo    *repository.IssueStatusRepository
}

func NewProjectCreateProjectIssueStatusFormValidator(
	input *insightInputs.ProjectIssueStatusInputForProjectCreate,
	repo *repository.IssueStatusRepository,
	projectIssueStatus *models.ProjectIssueStatus,
) ProjectCreateProjectIssueStatusForm {
	form := ProjectCreateProjectIssueStatusForm{
		Form:                                    Form{},
		ProjectIssueStatusInputForProjectCreate: *input,
		ProjectIssueStatus:                      projectIssueStatus,
		IssueStatusRepo:                         repo,
	}
	form.assignAttributes()

	return form
}

func (form *ProjectCreateProjectIssueStatusForm) assignAttributes() {
	form.AddAttributes(
		&IntAttribute[int32]{
			FieldAttribute: FieldAttribute{
				Code: "issueStatusId",
			},
			Value: helpers.GetInt32OrDefault(form.IssueStatusId),
		},
	)
}

func (form *ProjectCreateProjectIssueStatusForm) Validate() exceptions.ResourceModificationError {
	form.validateIssueStatusId().summaryErrors()

	if form.Errors != nil {
		return form.Errors
	}

	return nil
}

func (form *ProjectCreateProjectIssueStatusForm) validateIssueStatusId() *ProjectCreateProjectIssueStatusForm {
	field := form.FindAttrByCode("issueStatusId")
	field.ValidateRequired()

	if field.IsClean() {
		if err := form.IssueStatusRepo.Find(&models.IssueStatus{Id: *form.IssueStatusId}); err != nil {
			field.AddError("is invalid")
		}
	}

	if field.IsClean() {
		form.ProjectIssueStatus.IssueStatusId = *form.IssueStatusId
	}

	return form
}
