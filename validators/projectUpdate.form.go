package validators

import (
	"aio-server/database"
	"aio-server/enums"
	"aio-server/exceptions"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/constants"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"strings"
)

type ProjectUpdateForm struct {
	Form
	insightInputs.ProjectUpdateFormInput
	Project *models.Project
	Repo    *repository.ProjectRepository
}

func NewProjectUpdateFormValidator(
	input *insightInputs.ProjectUpdateFormInput,
	repo *repository.ProjectRepository,
	project *models.Project,
) ProjectUpdateForm {
	form := ProjectUpdateForm{
		Form:                   Form{},
		ProjectUpdateFormInput: *input,
		Project:                project,
		Repo:                   repo,
	}
	form.assignAttributes()

	return form
}

func (form *ProjectUpdateForm) Save() error {
	if err := form.validate(); err != nil {
		return err
	}

	if err := form.Repo.Update(form.Project, []string{"Name", "ProjectPriority", "Description", "ClientId", "State", "ProjectType", "SprintDuration", "StartedAt", "EndedAt"}); err != nil {
		return exceptions.NewUnprocessableContentError("", exceptions.ResourceModificationError{
			"base": {err.Error()},
		})
	}

	return nil
}

func (form *ProjectUpdateForm) assignAttributes() error {
	form.AddAttributes(
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "name",
			},
			Value: helpers.GetStringOrDefault(form.Name),
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "projectPriority",
			},
			Value: helpers.GetStringOrDefault(form.ProjectPriority),
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "description",
			},
			Value: helpers.GetStringOrDefault(form.Description),
		},
		&IntAttribute[int32]{
			FieldAttribute: FieldAttribute{
				Code: "clientId",
			},
			Value: helpers.GetInt32OrDefault(form.ClientId),
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "state",
			},
			Value: helpers.GetStringOrDefault(form.State),
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
		&TimeAttribute{
			FieldAttribute: FieldAttribute{
				Code: "startedAt",
			},
			Value: helpers.GetStringOrDefault(form.StartedAt),
		},
		&TimeAttribute{
			FieldAttribute: FieldAttribute{
				Code: "endedAt",
			},
			Value: helpers.GetStringOrDefault(form.EndedAt),
		},
		&IntAttribute[int32]{
			FieldAttribute: FieldAttribute{
				Code: "lockVersion",
			},
			Value: helpers.GetInt32OrDefault(form.LockVersion),
		},
	)
	return nil
}

func (form *ProjectUpdateForm) validate() error {
	form.validateName().
		validateProjectPriority().
		validateDescription().
		validateClientId().
		validateState().
		validateProjectType().
		validateStartedAt().
		validateEndedAt().
		validateLockVersion().
		summaryErrors()

	if form.Errors != nil {
		return exceptions.NewUnprocessableContentError("", form.Errors)
	}

	return nil
}

func (form *ProjectUpdateForm) validateName() *ProjectUpdateForm {
	field := form.FindAttrByCode("name")
	field.ValidateRequired()

	field.ValidateMin(interface{}(int64(5)))
	field.ValidateMax(interface{}(int64(constants.MaxStringLength)))

	if form.Name != nil && strings.TrimSpace(*form.Name) != "" {
		project := models.Project{Name: *form.Name}
		if err := form.Repo.Find(&project); err == nil {
			if project.Id != form.Project.Id {
				field.AddError("is already exists. Please use another name")
			}
		}

		if field.IsClean() {
			form.Project.Name = *form.Name
		}
	}

	return form
}

func (form *ProjectUpdateForm) validateProjectPriority() *ProjectUpdateForm {
	field := form.FindAttrByCode("projectPriority")
	field.ValidateRequired()

	if form.ProjectPriority != nil && strings.TrimSpace(*form.ProjectPriority) != "" {
		if projectPriority, err := enums.ParseProjectPriority(*form.ProjectPriority); err != nil {
			field.AddError("is invalid")
		} else {
			form.Project.ProjectPriority = projectPriority
		}
	}

	return form
}

func (form *ProjectUpdateForm) validateDescription() *ProjectUpdateForm {
	field := form.FindAttrByCode("description")

	field.ValidateRequired()

	field.ValidateMin(interface{}(int64(5)))
	field.ValidateMax(interface{}(int64(constants.MaxLongTextLength)))

	if field.IsClean() {
		form.Project.Description = form.Description
	}

	return form
}

func (form *ProjectUpdateForm) validateClientId() *ProjectUpdateForm {
	field := form.FindAttrByCode("clientId")

	if form.ClientId != nil {
		clientRepo := repository.NewClientRepository(nil, database.Db)

		if field.IsClean() {
			if err := clientRepo.Find(&models.Client{Id: *form.ClientId}); err != nil {
				field.AddError("is invalid")
			}
		}

		if field.IsClean() {
			form.Project.ClientId = *form.ClientId
		}
	}

	return form
}

func (form *ProjectUpdateForm) validateState() *ProjectUpdateForm {
	field := form.FindAttrByCode("state")
	field.ValidateRequired()

	if field.IsClean() {
		if state, err := enums.ParseProjectState(*form.State); err != nil {
			field.AddError("is invalid")
		} else {
			if field.IsClean() {
				form.Project.State = state
			}
		}
	}

	return form
}

func (form *ProjectUpdateForm) validateProjectType() *ProjectUpdateForm {
	field := form.FindAttrByCode("projectType")

	field.ValidateRequired()

	if field.IsClean() {
		if projectType, err := enums.ParseProjectType(*form.ProjectType); err != nil {
			field.AddError("is invalid")
		} else {
			if field.IsClean() {
				form.Project.ProjectType = projectType

				sprintDurationField := form.FindAttrByCode("sprintDuration")

				if projectType == enums.ProjectTypeScrum {
					sprintDurationField.ValidateRequired()

					if sprintDurationField.IsClean() {
						form.Project.SprintDuration = form.SprintDuration
					}

				} else if projectType == enums.ProjectTypeKanban {
					if form.SprintDuration != nil {
						sprintDurationField.AddError("need to be empty")
					}
				}
			}
		}
	}

	return form
}

func (form *ProjectUpdateForm) validateStartedAt() *ProjectUpdateForm {
	field := form.FindAttrByCode("startedAt")

	field.ValidateFormat(constants.DDMMYYY_DateFormat, constants.HUMAN_DD_MM_YY_DateFormat)

	if field.IsClean() {
		form.Project.StartedAt = field.Time()
	}

	return form
}

func (form *ProjectUpdateForm) validateEndedAt() *ProjectUpdateForm {
	field := form.FindAttrByCode("endedAt")

	field.ValidateFormat(constants.DDMMYYY_DateFormat, constants.HUMAN_DD_MM_YY_DateFormat)

	if field.IsClean() {
		form.Project.EndedAt = field.Time()
	}

	return form
}

func (form *ProjectUpdateForm) validateLockVersion() *ProjectUpdateForm {
	currentLockVersion := form.Project.LockVersion

	field := form.FindAttrByCode("lockVersion")

	field.ValidateRequired()

	if field.IsClean() {
		field.ValidateMin(interface{}(int64(currentLockVersion)))
	}

	return form
}
