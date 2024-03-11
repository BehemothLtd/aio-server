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

type ProjectCreateForm struct {
	Form
	insightInputs.ProjectCreateFormInput
	Project *models.Project
	Repo    *repository.ProjectRepository
}

func NewProjectCreateFormValidator(
	input *insightInputs.ProjectCreateFormInput,
	repo *repository.ProjectRepository,
	project *models.Project,
) ProjectCreateForm {
	form := ProjectCreateForm{
		Form:                   Form{},
		ProjectCreateFormInput: *input,
		Project:                project,
		Repo:                   repo,
	}
	form.assignAttributes()

	return form
}

func (form *ProjectCreateForm) Save() error {
	if err := form.validate(); err != nil {
		return err
	}

	return nil
}

func (form *ProjectCreateForm) assignAttributes() {
	form.AddAttributes(
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "name",
			},
			Value: helpers.GetStringOrDefault(form.Name),
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "code",
			},
			Value: helpers.GetStringOrDefault(form.Code),
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "description",
			},
			Value: helpers.GetStringOrDefault(form.Description),
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "projectType",
			},
			Value: helpers.GetStringOrDefault(form.ProjectType),
		},
		&IntAttribute[int]{
			FieldAttribute: FieldAttribute{
				Code: "sprintDuration",
			},
			Value: int(helpers.GetIntOrDefault(form.SprintDuration)),
		},
	)
}

func (form *ProjectCreateForm) validate() error {
	form.validateName().
		validateCode().
		validateDescription().
		validateProjectType().
		summaryErrors()

	if form.Errors != nil {
		return exceptions.NewUnprocessableContentError("", form.Errors)
	}
	return nil
}

func (form *ProjectCreateForm) validateName() *ProjectCreateForm {
	nameField := form.FindAttrByCode("name")

	nameField.ValidateRequired()

	min := 5
	max := int64(constants.MaxStringLength)
	nameField.ValidateLimit(&min, &max)

	if form.Name != nil && strings.TrimSpace(*form.Name) != "" {
		project := models.Project{Name: *form.Name}
		if err := form.Repo.Find(&project); err == nil {
			nameField.AddError("is already exists. Please use another name")
		}
	}

	return form
}

func (form *ProjectCreateForm) validateCode() *ProjectCreateForm {
	codeField := form.FindAttrByCode("code")

	codeField.ValidateRequired()

	min := 2
	max := int64(constants.MaxLongTextLength)
	codeField.ValidateLimit(&min, &max)

	if form.Code != nil && strings.TrimSpace(*form.Code) != "" {
		project := models.Project{Code: *form.Code}
		if err := form.Repo.Find(&project); err == nil {
			codeField.AddError("is already exists. Please use another code")
		}
	}

	return form
}

func (form *ProjectCreateForm) validateDescription() *ProjectCreateForm {
	descField := form.FindAttrByCode("description")

	descField.ValidateRequired()

	min := 5
	max := int64(constants.MaxLongTextLength)
	descField.ValidateLimit(&min, &max)

	return form
}

func (form *ProjectCreateForm) validateProjectType() *ProjectCreateForm {
	typeField := form.FindAttrByCode("projectType")

	typeField.ValidateRequired()

	if form.ProjectType != nil {
		fieldValue := enums.ProjectType(*form.ProjectType)

		if !fieldValue.IsValid() {
			typeField.AddError("is invalid")
		}

		if fieldValue == enums.ProjectTypeScrum {
			sprintDurationField := form.FindAttrByCode("sprintDuration")

			sprintDurationField.ValidateRequired()
		}
	}

	return form
}
