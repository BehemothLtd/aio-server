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
	field := form.FindAttrByCode("UserId")
	field.ValidateRequired()

	if field.IsClean() {
		userRepo := repository.NewUserRepository(nil, database.Db)

		if err := userRepo.Find(&models.User{Id: *form.UserId}); err != nil {
			field.AddError("Invalid User")
		}

		if field.IsClean() {
			form.ProjectAssignee.UserId = *form.UserId
		}
	}

	return form
}

func (form *ProjectAssigneeForm) validateDevelopmentId() *ProjectAssigneeForm {
	field := form.FindAttrByCode("DevelopmentRoleId")
	field.ValidateRequired()

	if field.IsClean() {
		if developmentRole := systems.FindDevelopmentRoleById(*form.DevelopmentRoleId); developmentRole == nil {
			field.AddError("is invalid")
		}

		if field.IsClean() {
			form.ProjectAssignee.DevelopmentRoleId = *form.DevelopmentRoleId
		}
	}

	return form
}

func (form *ProjectAssigneeForm) validateActive() *ProjectAssigneeForm {
	field := form.FindAttrByCode("Active")
	field.ValidateRequired()

	if form.Active == nil {
		field.AddError("is required")
	} else {
		form.ProjectAssignee.Active = *form.Active
	}

	return form
}

func (form *ProjectAssigneeForm) validateJoinDate() *ProjectAssigneeForm {
	field := form.FindAttrByCode("JoinDate")
	field.ValidateRequired()

	if field.IsClean() {
		field.ValidateFormat(constants.DDMMYYYY_DateFormat, constants.HUMAN_DDMMYYYY_DateFormat)

		if field.IsClean() {
			form.ProjectAssignee.JoinDate = field.Time()
		}
	}

	return form
}

func (form *ProjectAssigneeForm) validateLeaveDate() *ProjectAssigneeForm {
	field := form.FindAttrByCode("LeaveDate")

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
			form.ProjectAssignee.LeaveDate = field.Time()
		}
	}

	return form
}

func (form *ProjectAssigneeForm) validateDuplicate() *ProjectAssigneeForm {
	userIdField := form.FindAttrByCode("UserId")
	developmentIdField := form.FindAttrByCode("DevelopmentRoleId")

	if userIdField.IsClean() && developmentIdField.IsClean() {
		presentedProjectAssignee := models.ProjectAssignee{
			ProjectId:         form.Project.Id,
			UserId:            *form.UserId,
			DevelopmentRoleId: *form.DevelopmentRoleId,
		}

		if err := form.Repo.Find(&presentedProjectAssignee); err == nil {
			if form.ProjectAssignee.Id == presentedProjectAssignee.Id {
				return form
			}

			userIdField.AddError("already has this role")
		}
	}

	return form
}

func (form *ProjectAssigneeForm) validateLockVersion() *ProjectAssigneeForm {
	if form.ProjectAssignee.Id == 0 {
		return form
	}

	field := IntAttribute[int32]{
		FieldAttribute: FieldAttribute{
			Code: "lockVersion",
		},
		Value: helpers.GetInt32OrDefault(form.LockVersion),
	}
	form.AddAttributes(&field)

	field.ValidateRequired()
	field.ValidateMin(interface{}(int64(form.ProjectAssignee.LockVersion)))

	return form
}

func (form *ProjectAssigneeForm) Save() error {
	if err := form.validate(); err != nil {
		return err
	}

	if form.ProjectAssignee.Id != 0 {
		if err := form.Repo.Update(form.ProjectAssignee); err != nil {
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
