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

type ProjectSprintForm struct {
	Form
	insightInputs.ProjectSprintFormInput
	ProjectSprint       *models.ProjectSprint
	UpdateProjectSprint models.ProjectSprint
	Repo                *repository.ProjectSprintRepository
}

func NewProjectSprintFormValidator(
	input *insightInputs.ProjectSprintFormInput,
	repo *repository.ProjectSprintRepository,
	projectSprint *models.ProjectSprint,
) ProjectSprintForm {
	form := ProjectSprintForm{
		Form:                   Form{},
		ProjectSprintFormInput: *input,
		ProjectSprint:          projectSprint,
		UpdateProjectSprint:    models.ProjectSprint{},
		Repo:                   repo,
	}
	form.assignAttributes()

	return form
}

func (form *ProjectSprintForm) Save() error {
	if err := form.validate(); err != nil {
		return err
	}

	if form.ProjectSprint.Id == 0 {
		if err := form.Repo.Create(form.ProjectSprint); err != nil {
			return exceptions.NewUnprocessableContentError("", exceptions.ResourceModificationError{
				"base": {err.Error()},
			})
		}
	} else {

		if err := form.Repo.Update(form.ProjectSprint, form.UpdateProjectSprint); err != nil {
			return exceptions.NewUnprocessableContentError("", exceptions.ResourceModificationError{
				"base": {err.Error()},
			})
		}
	}

	return nil
}
func (form *ProjectSprintForm) validate() error {
	form.validateProjectId().
		validateTitle().
		validateStartDate().
		summaryErrors()

	if form.Errors != nil {
		return exceptions.NewUnprocessableContentError("", form.Errors)
	}
	return nil
}
func (form *ProjectSprintForm) validateTitle() *ProjectSprintForm {

	title := form.FindAttrByCode("title")

	title.ValidateMax(interface{}(int64(constants.MaxStringLength)))

	if title.IsClean() {
		projectSprint := models.ProjectSprint{Title: *form.Title, ProjectId: *form.ProjectId}

		if err := form.Repo.Find(&projectSprint); err == nil {

			if form.ProjectSprint.Id == 0 || projectSprint.Id != form.ProjectSprint.Id {
				title.AddError("is already exists. Please use another name")
			}
		}

		if title.IsClean() {
			if form.ProjectSprint.Id == 0 {
				form.ProjectSprint.Title = *form.Title
			} else {
				form.UpdateProjectSprint.Title = *form.Title
			}
		}
	}

	return form
}
func (form *ProjectSprintForm) validateProjectId() *ProjectSprintForm {
	projectId := form.FindAttrByCode("projectId")
	projectId.ValidateRequired()

	projectRepo := repository.NewProjectRepository(nil, database.Db)
	err := projectRepo.Find(&models.Project{Id: *form.ProjectId})

	if projectId.IsClean() {
		if err != nil || (form.ProjectSprint.Id != 0 && *form.ProjectId != form.ProjectSprint.ProjectId) {
			projectId.AddError("is invalid")
		} else {
			if form.ProjectSprint.Id == 0 {
				form.ProjectSprint.ProjectId = *form.ProjectId
			} else {
				form.UpdateProjectSprint.ProjectId = *form.ProjectId
			}
		}
	}
	return form
}

func (form *ProjectSprintForm) validateStartDate() *ProjectSprintForm {
	startDate := form.FindAttrByCode("startDate")
	startDate.ValidateRequired()
	startDate.ValidateFormat(constants.DDMMYYYY_DateFormat, constants.HUMAN_DDMMYYYY_DateFormat)

	project := models.Project{Id: *form.ProjectId}
	projectRepo := repository.NewProjectRepository(nil, database.Db)

	projectRepo.Find(&project)
	endDate := startDate.Time().AddDate(0, 0, int(*project.SprintDuration*7))

	var projectSprint models.ProjectSprint

	if form.ProjectSprint.Id != 0 {
		projectSprint = models.ProjectSprint{StartDate: *startDate.Time(), EndDate: &endDate, ProjectId: *form.ProjectId, Id: form.ProjectSprint.Id}
	} else {
		projectSprint = models.ProjectSprint{StartDate: *startDate.Time(), EndDate: &endDate, ProjectId: *form.ProjectId}
	}

	if err := form.Repo.FindCollapsedSprints(&projectSprint); err == nil {
		startDate.AddError("is duplicate with another sprints")
	} else {
		if form.ProjectSprint.Id != 0 {
			form.UpdateProjectSprint.StartDate = *startDate.Time()
			form.UpdateProjectSprint.EndDate = &endDate
		} else {
			form.ProjectSprint.StartDate = *startDate.Time()
			form.ProjectSprint.EndDate = &endDate
		}
	}

	return form
}

func (form *ProjectSprintForm) assignAttributes() {
	form.AddAttributes(
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "title",
			},
			Value: helpers.GetStringOrDefault(form.Title),
		},
		&IntAttribute[int32]{
			FieldAttribute: FieldAttribute{
				Code: "projectId",
			},
			Value: helpers.GetInt32OrDefault(form.ProjectId),
		},
		&TimeAttribute{
			FieldAttribute: FieldAttribute{
				Code: "startDate",
			},
			Value: helpers.GetStringOrDefault(form.StartDate),
		},
	)
}
