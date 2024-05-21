package validators

import (
	"aio-server/database"
	"aio-server/exceptions"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/constants"
	"aio-server/pkg/helpers"
	"aio-server/pkg/systems"
	"aio-server/repository"
	"strings"
)

type ProjectAssigneeForm struct {
	Form
	insightInputs.ProjectModifyProjectAssigneeFormInput
	Project         models.Project
	ProjectAssignee *models.ProjectAssignee
	formData        map[string]interface{}
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
		formData:                              map[string]interface{}{"ProjectId": project.Id},
		Repo:                                  repo,
	}

	form.assignAttributes()

	return form
}

func (form *ProjectAssigneeForm) assignAttributes() {
	form.AddAttributes(
		&IntAttribute[int32]{
			FieldAttribute: FieldAttribute{
				Code: "UserId",
			},
			Value: helpers.GetInt32OrDefault(form.UserId),
		},
		&IntAttribute[int32]{
			FieldAttribute: FieldAttribute{
				Code: "DevelopmentRoleId",
			},
			Value: helpers.GetInt32OrDefault(form.DevelopmentRoleId),
		},
		&TimeAttribute{
			FieldAttribute: FieldAttribute{
				Code: "JoinDate",
			},
			Value: helpers.GetStringOrDefault(form.JoinDate),
		},
		&TimeAttribute{
			FieldAttribute: FieldAttribute{
				Code: "LeaveDate",
			},
			Value: helpers.GetStringOrDefault(form.LeaveDate),
		},
		&BoolAttribute{
			FieldAttribute: FieldAttribute{
				Code: "Active",
			},
			Value: helpers.GetBoolOrDefault(form.Active),
		},
	)
}

func (form *ProjectAssigneeForm) validate() error {
	form.validateUserId().
		validateDevelopmentId().
		validateActive().
		validateJoinDate().
		validateLeaveDate().
		validateDuplicate().
		validateLockVersion().
		summaryErrors()

	if form.Errors != nil {
		return exceptions.NewUnprocessableContentError("", form.Errors)
	}
	return nil
}

func (form *ProjectAssigneeForm) validateUserId() *ProjectAssigneeForm {
	code := "UserId"
	field := form.FindAttrByCode(code)
	field.ValidateRequired()

	if field.IsClean() {
		userRepo := repository.NewUserRepository(nil, database.Db)

		if err := userRepo.Find(&models.User{Id: *form.UserId}); err != nil {
			field.AddError("Invalid User")
		}

		if field.IsClean() {
			form.formData[code] = *form.UserId

			if form.ProjectAssignee.Id == 0 {
				form.ProjectAssignee.UserId = *form.UserId
			}
		}
	}

	return form
}

func (form *ProjectAssigneeForm) validateDevelopmentId() *ProjectAssigneeForm {
	code := "DevelopmentRoleId"
	field := form.FindAttrByCode(code)
	field.ValidateRequired()

	if field.IsClean() {
		if developmentRole := systems.FindDevelopmentRoleById(*form.DevelopmentRoleId); developmentRole == nil {
			field.AddError("is invalid")
		}

		if field.IsClean() {
			form.formData[code] = *form.DevelopmentRoleId

			if form.ProjectAssignee.Id == 0 {
				form.ProjectAssignee.DevelopmentRoleId = *form.DevelopmentRoleId
			}
		}
	}

	return form
}

func (form *ProjectAssigneeForm) validateActive() *ProjectAssigneeForm {
	code := "Active"
	field := form.FindAttrByCode("Active")
	field.ValidateRequired()

	if form.Active == nil {
		field.AddError("is required")
	} else {
		form.formData[code] = *form.Active

		if form.ProjectAssignee.Id == 0 {
			form.ProjectAssignee.Active = *form.Active
		}
	}

	return form
}

func (form *ProjectAssigneeForm) validateJoinDate() *ProjectAssigneeForm {
	code := "JoinDate"
	field := form.FindAttrByCode(code)
	field.ValidateRequired()

	if field.IsClean() {
		field.ValidateFormat(constants.DDMMYYYY_DateFormat, constants.HUMAN_DDMMYYYY_DateFormat)

		if field.IsClean() {
			form.formData[code] = field.Time()

			if form.ProjectAssignee.Id == 0 {
				form.ProjectAssignee.JoinDate = field.Time()
			}
		}
	}

	return form
}

func (form *ProjectAssigneeForm) validateLeaveDate() *ProjectAssigneeForm {
	code := "LeaveDate"
	field := form.FindAttrByCode(code)

	if form.Active != nil && !*form.Active {
		field.ValidateRequired()
	}

	if form.LeaveDate != nil && *form.LeaveDate != "" && strings.TrimSpace(*form.LeaveDate) != "" {
		field.ValidateFormat(constants.DDMMYYYY_DateFormat, constants.HUMAN_DDMMYYYY_DateFormat)

		joinDateTime := form.FindAttrByCode("JoinDate").Time()

		if joinDateTime != nil {
			field.ValidateMin(interface{}(*joinDateTime))
		}

		if field.IsClean() {
			form.formData[code] = field.Time()

			if form.ProjectAssignee.Id == 0 {
				form.ProjectAssignee.LeaveDate = field.Time()
			}
		}
	}

	return form
}

func (form *ProjectAssigneeForm) validateDuplicate() *ProjectAssigneeForm {
	userIdField := form.FindAttrByCode("UserId")
	developmentIdField := form.FindAttrByCode("DevelopmentRoleId")

	if userIdField.IsClean() && developmentIdField.IsClean() {
		presentedProjectAssignee := models.ProjectAssignee{
			ProjectId: form.Project.Id,
			UserId:    *form.UserId,
		}

		if err := form.Repo.Find(&presentedProjectAssignee); err == nil {
			if form.ProjectAssignee.Id == presentedProjectAssignee.Id {
				return form
			}

			userIdField.AddError("already presented")
		}
	}

	return form
}

func (form *ProjectAssigneeForm) validateLockVersion() *ProjectAssigneeForm {
	if form.ProjectAssignee.Id == 0 {
		return form
	}

	code := "LockVersion"

	field := IntAttribute[int32]{
		FieldAttribute: FieldAttribute{
			Code: code,
		},
		Value: helpers.GetInt32OrDefault(form.LockVersion),
	}
	form.AddAttributes(&field)

	field.ValidateRequired()
	field.ValidateMin(interface{}(int64(*form.LockVersion)))

	return form
}

func (form *ProjectAssigneeForm) Save() error {
	if err := form.validate(); err != nil {
		return err
	}

	if form.ProjectAssignee.Id != 0 {
		if err := form.Repo.Update(form.ProjectAssignee, form.formData); err != nil {
			return exceptions.NewUnprocessableContentError("", exceptions.ResourceModificationError{
				"base": {err.Error()},
			})
		}
	} else {

		if err := form.Repo.Create(form.ProjectAssignee); err != nil {
			return exceptions.NewUnprocessableContentError("", exceptions.ResourceModificationError{
				"base": {err.Error()},
			})
		}
	}

	return nil
}
