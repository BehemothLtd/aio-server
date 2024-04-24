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
	updates map[string]interface{}
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
		updates:                map[string]interface{}{},
		Repo:                   repo,
	}
	form.assignAttributes()

	return form
}

func (form *ProjectUpdateForm) Save() error {
	if err := form.validate(); err != nil {
		return err
	}

	if err := form.Repo.UpdateBasicInfo(form.Project, form.updates); err != nil {
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
				Code: "Name",
			},
			Value: helpers.GetStringOrDefault(form.Name),
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "ProjectPriority",
			},
			Value: helpers.GetStringOrDefault(form.ProjectPriority),
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "Description",
			},
			Value: helpers.GetStringOrDefault(form.Description),
		},
		&IntAttribute[int32]{
			FieldAttribute: FieldAttribute{
				Code: "ClientId",
			},
			Value: helpers.GetInt32OrDefault(form.ClientId),
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "State",
			},
			Value: helpers.GetStringOrDefault(form.State),
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "ProjectType",
			},
			Value: helpers.GetStringOrDefault(form.ProjectType),
		},
		&IntAttribute[int32]{
			FieldAttribute: FieldAttribute{
				Code: "SprintDuration",
			},
			Value: helpers.GetInt32OrDefault(form.SprintDuration),
		},
		&TimeAttribute{
			FieldAttribute: FieldAttribute{
				Code: "StartedAt",
			},
			Value: helpers.GetStringOrDefault(form.StartedAt),
		},
		&TimeAttribute{
			FieldAttribute: FieldAttribute{
				Code: "EndedAt",
			},
			Value: helpers.GetStringOrDefault(form.EndedAt),
		},
		&IntAttribute[int32]{
			FieldAttribute: FieldAttribute{
				Code: "LockVersion",
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
	code := "Name"

	field := form.FindAttrByCode(code)
	field.ValidateRequired()

	field.ValidateMin(interface{}(int64(5)))
	field.ValidateMax(interface{}(int64(constants.MaxStringLength)))

	if field.IsClean() {
		project := models.Project{Name: *form.Name}

		if err := form.Repo.Find(&project); err == nil {

			if project.Id != form.Project.Id {
				field.AddError("is already exists. Please use another name")
			}
		}

		if field.IsClean() {
			form.updates[code] = *form.Name
		}
	}

	return form
}

func (form *ProjectUpdateForm) validateProjectPriority() *ProjectUpdateForm {
	code := "ProjectPriority"
	field := form.FindAttrByCode(code)
	field.ValidateRequired()

	if form.ProjectPriority != nil && strings.TrimSpace(*form.ProjectPriority) != "" {
		if projectPriority, err := enums.ParseProjectPriority(*form.ProjectPriority); err != nil {
			field.AddError("is invalid")
		} else {
			form.updates[code] = projectPriority
		}
	}

	return form
}

func (form *ProjectUpdateForm) validateDescription() *ProjectUpdateForm {
	code := "Description"
	field := form.FindAttrByCode(code)

	field.ValidateRequired()

	field.ValidateMin(interface{}(int64(5)))
	field.ValidateMax(interface{}(int64(constants.MaxLongTextLength)))

	if field.IsClean() {
		form.updates[code] = form.Description
	}

	return form
}

func (form *ProjectUpdateForm) validateClientId() *ProjectUpdateForm {
	code := "ClientId"
	field := form.FindAttrByCode(code)

	if form.ClientId != nil {
		clientRepo := repository.NewClientRepository(nil, database.Db)

		if field.IsClean() {
			if err := clientRepo.Find(&models.Client{Id: *form.ClientId}); err != nil {
				field.AddError("is invalid")
			}
		}

		if field.IsClean() {
			form.updates[code] = form.ClientId
		}
	}

	return form
}

func (form *ProjectUpdateForm) validateState() *ProjectUpdateForm {
	code := "State"
	field := form.FindAttrByCode(code)
	field.ValidateRequired()

	if field.IsClean() {
		if state, err := enums.ParseProjectState(*form.State); err != nil {
			field.AddError("is invalid")
		} else {
			if field.IsClean() {
				form.updates[code] = state
			}
		}
	}

	return form
}

func (form *ProjectUpdateForm) validateProjectType() *ProjectUpdateForm {
	code := "ProjectType"
	field := form.FindAttrByCode(code)

	field.ValidateRequired()

	if field.IsClean() {
		if projectType, err := enums.ParseProjectType(*form.ProjectType); err != nil {
			field.AddError("is invalid")
		} else {
			if field.IsClean() {
				form.updates[code] = projectType

				sprintDurationCode := "SprintDuration"
				sprintDurationField := form.FindAttrByCode(sprintDurationCode)

				if projectType == enums.ProjectTypeScrum {
					sprintDurationField.ValidateRequired()

					if sprintDurationField.IsClean() {
						form.updates[sprintDurationCode] = form.SprintDuration
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
	code := "StartedAt"
	field := form.FindAttrByCode(code)

	field.ValidateFormat(constants.DDMMYYYY_DateFormat, constants.HUMAN_DDMMYYYY_DateFormat)

	if field.IsClean() {
		form.updates[code] = field.Time()
	}

	return form
}

func (form *ProjectUpdateForm) validateEndedAt() *ProjectUpdateForm {
	code := "EndedAt"
	field := form.FindAttrByCode(code)

	field.ValidateFormat(constants.DDMMYYYY_DateFormat, constants.HUMAN_DDMMYYYY_DateFormat)

	if field.IsClean() {
		form.updates[code] = field.Time()
	}

	return form
}

func (form *ProjectUpdateForm) validateLockVersion() *ProjectUpdateForm {
	currentLockVersion := form.Project.LockVersion

	field := form.FindAttrByCode("LockVersion")

	field.ValidateRequired()

	if field.IsClean() {
		field.ValidateMin(interface{}(int64(currentLockVersion)))
	}

	return form
}
