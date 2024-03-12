package validators

import (
	"aio-server/database"
	"aio-server/enums"
	"aio-server/exceptions"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/constants"
	"aio-server/pkg/helpers"
	"aio-server/pkg/systems"
	"aio-server/repository"
	"fmt"
	"slices"
	"strings"
	"time"
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

	if err := form.Repo.Create(form.Project); err != nil {
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
		&IntAttribute[int32]{
			FieldAttribute: FieldAttribute{
				Code: "sprintDuration",
			},
			Value: helpers.GetInt32OrDefault(form.SprintDuration),
		},
		&SliceAttribute[insightInputs.ProjectIssueStatusInputForProjectCreate]{
			FieldAttribute: FieldAttribute{
				Code: "projectIssueStatuses",
			},
			Value: &form.ProjectIssueStatuses,
		},
		&SliceAttribute[insightInputs.ProjectAssigneeInputForProjectCreate]{
			FieldAttribute: FieldAttribute{
				Code: "projectAssignees",
			},
			Value: &form.ProjectAssignees,
		},
	)
}

func (form *ProjectCreateForm) validate() error {
	form.validateName().
		validateCode().
		validateDescription().
		validateProjectType().
		validateProjectIssueStatuses().
		validateProjectAssignees().
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

		form.Project.Name = *form.Name
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

		form.Project.Code = *form.Code
	}

	return form
}

func (form *ProjectCreateForm) validateDescription() *ProjectCreateForm {
	descField := form.FindAttrByCode("description")

	descField.ValidateRequired()

	min := 5
	max := int64(constants.MaxLongTextLength)
	descField.ValidateLimit(&min, &max)

	form.Project.Description = form.Description

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

		form.Project.ProjectType = fieldValue

		sprintDurationField := form.FindAttrByCode("sprintDuration")

		if fieldValue == enums.ProjectTypeScrum {
			sprintDurationField.ValidateRequired()

			form.Project.SprintDuration = form.SprintDuration
		} else if fieldValue == enums.ProjectTypeKanban {
			if form.SprintDuration != nil {
				sprintDurationField.AddError("need to be empty")
			}
		}
	}

	return form
}

func (form *ProjectCreateForm) validateProjectIssueStatuses() *ProjectCreateForm {
	fieldKey := "projectIssueStatuses"
	projectIssueStatusesField := form.FindAttrByCode(fieldKey)

	projectIssueStatusesField.ValidateRequired()

	issueStatusRepo := repository.NewIssueStatusRepository(nil, database.Db)

	if form.ProjectIssueStatuses != nil {
		projectIssueStatuses := []*models.ProjectIssueStatus{}

		for i, projectIssueStatusInput := range form.ProjectIssueStatuses {
			issueStatusId := projectIssueStatusInput.IssueStatusId

			// Check duplicated in input
			if foundIdx := slices.IndexFunc(projectIssueStatuses, func(pis *models.ProjectIssueStatus) bool {
				return pis.IssueStatusId == issueStatusId
			}); foundIdx != -1 {
				form.AddError(fmt.Sprintf("%s.%d.%s", fieldKey, i, "issueStatusId"), []interface{}{"is duplicated"})
			} else {
				// If not duplicated then create nested form for further validation
				projectIssueStatus := models.ProjectIssueStatus{
					IssueStatusId: issueStatusId,
				}

				projectIssueStatusForm := NewProjectCreateProjectIssueFormValidator(
					&projectIssueStatusInput,
					issueStatusRepo,
					&projectIssueStatus,
				)

				if err := projectIssueStatusForm.Validate(); err != nil {
					for key, innerErr := range err {
						form.AddError(fmt.Sprintf("%s.%d.%s", fieldKey, i, key), innerErr)
					}
				} else {
					// only push to final result when nested form has no error
					projectIssueStatuses = append(projectIssueStatuses, &models.ProjectIssueStatus{
						IssueStatusId: issueStatusId,
						Position:      i + 1,
					})
				}
			}

			form.Project.ProjectIssueStatuses = projectIssueStatuses
		}
	}

	return form
}

func (form *ProjectCreateForm) validateProjectAssignees() *ProjectCreateForm {
	projectAssigneesField := form.FindAttrByCode("projectAssignees")

	projectAssigneesField.ValidateRequired()

	userRepo := repository.NewUserRepository(nil, database.Db)
	if form.ProjectAssignees != nil {
		projectAssignees := []*models.ProjectAssignee{}

		for i, projectAssigneeInput := range form.ProjectAssignees {
			userId := projectAssigneeInput.UserId
			developentRoleId := projectAssigneeInput.DevelopmentRoleId
			active := projectAssigneeInput.Active

			// validate valid User
			if err := userRepo.Find(&models.User{Id: userId}); err != nil {
				projectAssigneesField.AddError(
					map[string]interface{}{fmt.Sprintf("%d", i): map[string][]string{"userId": {"is invalid"}}},
				)
			}

			// validate valid developmentRole
			if developmentRole := systems.FindDevelopmentRoleById(developentRoleId); developmentRole == nil {
				projectAssigneesField.AddError(
					map[string]interface{}{fmt.Sprintf("%d", i): map[string][]string{"developmentRoleId": {"is invalid"}}},
				)
			}

			if foundIdx := slices.IndexFunc(projectAssignees, func(pa *models.ProjectAssignee) bool {
				return pa.UserId == userId && pa.DevelopmentRoleId == developentRoleId
			}); foundIdx != -1 {
				projectAssigneesField.AddError(
					map[string]interface{}{fmt.Sprintf("%d", i): map[string][]string{"userId": {"is already has this role"}}},
				)
			}

			if len(projectAssigneesField.GetErrors()) == 0 {
				if joinDate, err := time.Parse("1-2-2006", projectAssigneeInput.JoinDate); err != nil {
					projectAssigneesField.AddError(
						map[string]interface{}{fmt.Sprintf("%d", i): map[string][]string{"joinDate": {"need to be formatted as %d-%m-%y"}}},
					)
				} else {
					projectAssignees = append(projectAssignees, &models.ProjectAssignee{
						UserId:            userId,
						DevelopmentRoleId: developentRoleId,
						Active:            active,
						JoinDate:          &joinDate,
					})
				}

			}
		}
		form.Project.ProjectAssignees = projectAssignees
	}

	return form
}
