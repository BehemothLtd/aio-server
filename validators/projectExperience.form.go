package validators

import (
	"aio-server/database"
	"aio-server/exceptions"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/constants"
	"aio-server/pkg/helpers"
	"aio-server/repository"
)

type ProjectExperienceForm struct {
	Form
	insightInputs.ProjectExperienceFormInput
	ProjectExperience *models.ProjectExperience
	repo              *repository.ProjectExperienceRepository
	UpdatesForm       map[string]interface{}
}

func NewProjectExperienceFormValidator(
	input *insightInputs.ProjectExperienceFormInput,
	repo *repository.ProjectExperienceRepository,
	ProjectExperience *models.ProjectExperience,
) ProjectExperienceForm {
	form := ProjectExperienceForm{
		Form:                       Form{},
		ProjectExperienceFormInput: *input,
		ProjectExperience:          ProjectExperience,
		UpdatesForm:                map[string]interface{}{},
		repo:                       repo,
	}
	form.assignAttributes()

	return form
}

func (form *ProjectExperienceForm) assignAttributes() {
	form.AddAttributes(
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "Title",
			},
			Value: helpers.GetStringOrDefault(&form.Title),
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "Description",
			},
			Value: helpers.GetStringOrDefault(&form.Description),
		},
		&IntAttribute[int32]{
			FieldAttribute: FieldAttribute{
				Code: "ProjectId",
			},
			Value: helpers.GetInt32OrDefault(&form.ProjectId),
		},
	)
}

func (form *ProjectExperienceForm) Save() error {
	if err := form.validate(); err != nil {
		return err
	}

	if form.ProjectExperience.Id == 0 {
		if err := form.repo.Create(form.ProjectExperience); err != nil {
			return exceptions.NewUnprocessableContentError("", exceptions.ResourceModificationError{
				"base": {err.Error()},
			})
		}
	} else {
		if err := form.repo.Update(form.ProjectExperience, form.UpdatesForm); err != nil {
			return exceptions.NewUnprocessableContentError("", exceptions.ResourceModificationError{
				"base": {err.Error()},
			})
		}
	}

	return nil
}

func (form *ProjectExperienceForm) validate() error {
	form.validateTitle().validateDescription().validateProjectId()
	form.summaryErrors()

	if form.Errors != nil {
		return exceptions.NewUnprocessableContentError("", form.Errors)
	}
	return nil
}

func (form *ProjectExperienceForm) validateTitle() *ProjectExperienceForm {
	fieldCode := "Title"

	titleField := form.FindAttrByCode(fieldCode)

	titleField.ValidateRequired()

	titleField.ValidateMin(interface{}(int64(2)))
	titleField.ValidateMax(interface{}(int64(constants.MaxStringLength)))

	if titleField.IsClean() {
		if form.ProjectExperience.Id == 0 {
			form.ProjectExperience.Title = form.Title
		} else {
			form.UpdatesForm[fieldCode] = form.Title
		}
	}
	return form
}

func (form *ProjectExperienceForm) validateDescription() *ProjectExperienceForm {
	fieldCode := "Description"
	field := form.FindAttrByCode(fieldCode)

	field.ValidateRequired()
	field.ValidateMax(interface{}(int64(constants.MaxLongTextLength)))

	if field.IsClean() {
		if form.ProjectExperience.Id == 0 {
			form.ProjectExperience.Description = form.Description
		} else {
			form.UpdatesForm[fieldCode] = form.Description
		}
	}

	return form
}

func (form *ProjectExperienceForm) validateProjectId() *ProjectExperienceForm {
	fieldCode := "ProjectId"
	projectId := form.FindAttrByCode(fieldCode)
	projectId.ValidateRequired()

	projectRepo := repository.NewProjectRepository(nil, database.Db)
	if err := projectRepo.Find(&models.Project{Id: form.ProjectId}); err != nil {
		projectId.AddError("is invalid")
	}
	if projectId.IsClean() {
		if form.ProjectExperience.Id == 0 {
			form.ProjectExperience.ProjectId = form.ProjectId
		} else {
			form.UpdatesForm[fieldCode] = form.ProjectId
		}
	}
	return form
}
