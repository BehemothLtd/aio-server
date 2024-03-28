package validators

import (
	"aio-server/database"
	"aio-server/exceptions"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/repository"
)

type IssueCreateIssueAssigneeForm struct {
	Form
	insightInputs.IssueAssigneeInputForIssueCreate

	Project           models.Project
	IssueAssignee     *models.IssueAssignee
	IssueAssigneeRepo *repository.IssueAssigneeRepository
}

func NewIssueCreateIssueAssigneeFormValidator(
	input *insightInputs.IssueAssigneeInputForIssueCreate,
	repo *repository.IssueAssigneeRepository,
	issueAssignee *models.IssueAssignee,
	project models.Project,
) IssueCreateIssueAssigneeForm {
	form := IssueCreateIssueAssigneeForm{
		Form:                             Form{},
		IssueAssigneeInputForIssueCreate: *input,
		IssueAssignee:                    issueAssignee,
		Project:                          project,
		IssueAssigneeRepo:                repo,
	}
	form.assignAttributes()

	return form
}

func (form *IssueCreateIssueAssigneeForm) assignAttributes() {
	form.AddAttributes(
		&IntAttribute[int32]{
			FieldAttribute: FieldAttribute{
				Code: "userId",
			},
			Value: helpers.GetInt32OrDefault(form.UserId),
		},
		&IntAttribute[int32]{
			FieldAttribute: FieldAttribute{
				Code: "developmentRoleId",
			},
			Value: helpers.GetInt32OrDefault(form.DevelopmentRoleId),
		},
	)
}

func (form *IssueCreateIssueAssigneeForm) Validate() exceptions.ResourceModificationError {
	form.validateUserId().
		validateDevelopmentRoleId().
		validateValidUserAndDevelopmentRole().
		summaryErrors()

	if form.Errors != nil {
		return form.Errors
	}

	return nil
}

func (form *IssueCreateIssueAssigneeForm) validateUserId() *IssueCreateIssueAssigneeForm {
	field := form.FindAttrByCode("userId")
	field.ValidateRequired()

	if field.IsClean() {
		form.IssueAssignee.UserId = *form.UserId
	}

	return form
}

func (form *IssueCreateIssueAssigneeForm) validateDevelopmentRoleId() *IssueCreateIssueAssigneeForm {
	field := form.FindAttrByCode("developmentRoleId")
	field.ValidateRequired()

	if field.IsClean() {
		form.IssueAssignee.DevelopmentRoleId = *form.DevelopmentRoleId
	}

	return form
}

func (form *IssueCreateIssueAssigneeForm) validateValidUserAndDevelopmentRole() *IssueCreateIssueAssigneeForm {
	userIdField := form.FindAttrByCode("userId")
	developmentRoleIdField := form.FindAttrByCode("developmentRoleId")

	if userIdField.IsClean() && developmentRoleIdField.IsClean() {
		projectAssigneeRepo := repository.NewProjectAssigneeRepository(nil, database.Db)

		projectAssignee := models.ProjectAssignee{ProjectId: form.Project.Id, UserId: *form.UserId, DevelopmentRoleId: *form.DevelopmentRoleId}

		if err := projectAssigneeRepo.Find(&projectAssignee); err != nil {
			userIdField.AddError("is invalid")
			developmentRoleIdField.AddError("is invalid")
		}
	}

	return form
}
