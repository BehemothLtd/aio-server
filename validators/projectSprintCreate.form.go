package validators

import (
	"aio-server/database"
	"aio-server/exceptions"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/constants"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"strings"
)

type ProjectSprintForm struct {
	Form
	insightInputs.ProjectSprintFormInput
	ProjectSprint *models.ProjectSprint
	Repo          *repository.ProjectSprintRepository
}

func NewProjectSprintCreateFormValidator(
	input *insightInputs.ProjectSprintFormInput,
	repo *repository.ProjectSprintRepository,
	projectSprint *models.ProjectSprint,
) ProjectSprintForm {
	form := ProjectSprintForm{
		Form:                   Form{},
		ProjectSprintFormInput: *input,
		ProjectSprint:          projectSprint,
		Repo:                   repo,
	}
	form.assignAttributes()

	return form
}

func (form *ProjectSprintForm) Save() error {

	if err := form.validate(); err != nil {
		return err
	}
	if err := form.Repo.Create(form.ProjectSprint); err != nil {
		return err
	}

	return nil
}

func (form *ProjectSprintForm) validate() error {
	form.validateTitle().
		validateProjectId().
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

	if form.Title != nil && strings.TrimSpace(*form.Title) != "" {
		projectSprint := models.ProjectSprint{Title: *form.Title, ProjectId: *form.ProjectId}
		if err := form.Repo.Find(&projectSprint); err == nil {
			title.AddError("is already exists. Please use another name")
		}

		if title.IsClean() {
			form.ProjectSprint.Title = *form.Title

		}
	}

	return form
}

func (form *ProjectSprintForm) validateProjectId() *ProjectSprintForm {
	projectId := form.FindAttrByCode("projectId")

	projectId.ValidateRequired()

	projectRepo := repository.NewProjectRepository(nil, database.Db)

	if projectId.IsClean() {
		if err := projectRepo.Find(&models.Project{Id: *form.ProjectId}); err != nil {
			projectId.AddError("is invalid")
		} else {
			form.ProjectSprint.ProjectId = *form.ProjectId
		}
	}

	return form
}

func (form *ProjectSprintForm) validateStartDate() *ProjectSprintForm {
	startDate := form.FindAttrByCode("startDate")
	startDate.ValidateRequired()
	startDate.ValidateFormat("1-2-2006", "%d-%m-%y")

	project := models.Project{Id: *form.ProjectId}
	projectRepo := repository.NewProjectRepository(nil, database.Db)

	projectRepo.Find(&project)

	endDate := startDate.Time().AddDate(0, 0, int(project.SprintDuration*7))

	projectSprint := models.ProjectSprint{StartDate: *startDate.Time(), EndDate: &endDate}
	sprintError := form.Repo.CollapsedSprints(&projectSprint)

	if sprintError == nil {
		startDate.AddError("is duplicate with another sprints")
	} else {
		form.ProjectSprint.StartDate = *startDate.Time()
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
