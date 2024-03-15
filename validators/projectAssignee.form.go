package validators

import (
	"aio-server/database"
	"aio-server/exceptions"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/repository"
)

type ProjectAssigneeForm struct {
	Form
	insightInputs.ProjectModifyProjectAssigneeFormInput
	Project         models.Project
	ProjectAssignee *models.ProjectAssignee
	Repo            *repository.ProjectAssigneeRepository
}

func NewProjectAssigneeFormValidator(
	input insightInputs.ProjectModifyProjectAssigneeFormInput,
	repo *repository.ProjectAssigneeRepository,
	project models.Project,
	projectAssignee *models.ProjectAssignee,
) ProjectAssigneeForm {
	form := ProjectAssigneeForm{
		Form:                                  Form{},
		ProjectModifyProjectAssigneeFormInput: input,
		Project:                               project,
		ProjectAssignee:                       projectAssignee,
		Repo:                                  repo,
	}

	form.assignAttributes()

	return form
}

func (form *ProjectAssigneeForm) assignAttributes() {
	form.AddAttributes(
		&IntAttribute[int32]{
			FieldAttribute: FieldAttribute{
				Code: "userId",
			},
			Value: helpers.GetInt32OrDefault(&form.UserId),
		},
		&IntAttribute[int32]{
			FieldAttribute: FieldAttribute{
				Code: "developmentRoleId",
			},
			Value: helpers.GetInt32OrDefault(&form.UserId),
		},
	)
}

func (form *ProjectAssigneeForm) validate() error {
	form.validateUserId().summaryErrors()

	if form.Errors != nil {
		return exceptions.NewUnprocessableContentError("", form.Errors)
	}
	return nil
}

func (form *ProjectAssigneeForm) validateUserId() *ProjectAssigneeForm {
	field := form.FindAttrByCode("userId")
	field.ValidateRequired()

	userRepo := repository.NewUserRepository(nil, database.Db)
	if err := userRepo.Find(&models.User{Id: form.UserId}); err != nil {
		field.AddError("Invalid User")
	}

	if field.IsClean() {
		form.ProjectAssignee.UserId = form.UserId
	}

	return form
}

func (form *ProjectAssigneeForm) Save() error {
	if err := form.validate(); err != nil {
		return err
	}

	if err := form.Repo.Create(form.ProjectAssignee); err != nil {
		return err
	}

	return nil
}
