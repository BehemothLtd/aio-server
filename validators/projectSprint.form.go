package validators

import (
	"aio-server/exceptions"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/constants"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"fmt"
	"strings"
)

type ProjectSprintForm struct {
	Form
	insightInputs.ProjectSprintFormInput
	ProjectSprint       *models.ProjectSprint
	Project             models.Project
	UpdateProjectSprint models.ProjectSprint
	Repo                *repository.ProjectSprintRepository
}

func NewProjectSprintFormValidator(
	input *insightInputs.ProjectSprintFormInput,
	repo *repository.ProjectSprintRepository,
	projectSprint *models.ProjectSprint,
	project models.Project,
) ProjectSprintForm {
	form := ProjectSprintForm{
		Form:                   Form{},
		ProjectSprintFormInput: *input,
		ProjectSprint:          projectSprint,
		Project:                project,
		UpdateProjectSprint:    models.ProjectSprint{},
		Repo:                   repo,
	}

	form.assignAttributes()

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
		&TimeAttribute{
			FieldAttribute: FieldAttribute{
				Code: "startDate",
			},
			Value: helpers.GetStringOrDefault(form.StartDate),
		},
	)
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
	form.validateTitle().
		validateStartDate().
		summaryErrors()

	if form.Errors != nil {
		return exceptions.NewUnprocessableContentError("", form.Errors)
	}
	return nil
}

func (form *ProjectSprintForm) validateTitle() *ProjectSprintForm {
	if form.Title != nil {
		title := form.FindAttrByCode("title")
		title.ValidateMax(interface{}(int64(constants.MaxStringLength)))

		if title.IsClean() {
			projectSprint := models.ProjectSprint{Title: *form.Title, ProjectId: form.Project.Id}

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
	}

	return form
}

func (form *ProjectSprintForm) validateStartDate() *ProjectSprintForm {
	startDate := form.FindAttrByCode("startDate")
	startDate.ValidateRequired()
	startDate.ValidateFormat(constants.DDMMYYYY_DateFormat, constants.HUMAN_DDMMYYYY_DateFormat)

	if startDate.IsClean() {
		endDate := startDate.Time().AddDate(0, 0, int(*form.Project.SprintDuration*7))

		var projectSprint models.ProjectSprint

		if form.ProjectSprint.Id != 0 {
			projectSprint = models.ProjectSprint{StartDate: *startDate.Time(), EndDate: &endDate, ProjectId: form.Project.Id, Id: form.ProjectSprint.Id}
		} else {
			projectSprint = models.ProjectSprint{StartDate: *startDate.Time(), EndDate: &endDate, ProjectId: form.Project.Id}
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

			if form.ProjectSprint.Title == "" || strings.TrimSpace(form.ProjectSprint.Title) == "" {
				form.ProjectSprint.Title = fmt.Sprintf("Sprint: %s ~ %s", startDate.Time().Format(constants.DDMMYYYY_DateSplashFormat), endDate.Format(constants.DDMMYYYY_DateSplashFormat))
			}
		}
	}

	return form
}
