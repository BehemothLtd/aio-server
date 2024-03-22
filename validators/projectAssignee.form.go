package validators

import (
	"aio-server/database"
	"aio-server/exceptions"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
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
				Code: "userId",
			},
			Value: helpers.GetInt32OrDefault(&form.UserId),
		},
		&IntAttribute[int32]{
			FieldAttribute: FieldAttribute{
				Code: "developmentRoleId",
			},
			Value: helpers.GetInt32OrDefault(&form.UserId),
		},
		&TimeAttribute{
			FieldAttribute: FieldAttribute{
				Code: "joinDate",
			},
			Value: helpers.GetStringOrDefault(&form.JoinDate),
		},
		&TimeAttribute{
			FieldAttribute: FieldAttribute{
				Code: "leaveDate",
			},
			Value: helpers.GetStringOrDefault(form.LeaveDate),
		},
	)

	form.ProjectAssignee.Active = form.Active
}

func (form *ProjectAssigneeForm) validate() error {
	form.validateUserId().
		validateDevelopmentId().
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
	field := form.FindAttrByCode("userId")
	field.ValidateRequired()

	userRepo := repository.NewUserRepository(nil, database.Db)
	if err := userRepo.Find(&models.User{Id: form.UserId}); err != nil {
		field.AddError("Invalid User")
	}

	if field.IsClean() {
		form.ProjectAssignee.UserId = form.UserId
	}

	return form
}

func (form *ProjectAssigneeForm) validateDevelopmentId() *ProjectAssigneeForm {
	field := form.FindAttrByCode("developmentRoleId")
	field.ValidateRequired()

	if field.IsClean() {
		if developmentRole := systems.FindDevelopmentRoleById(form.DevelopmentRoleId); developmentRole == nil {
			field.AddError("is invalid")
		}
	}

	if field.IsClean() {
		form.ProjectAssignee.DevelopmentRoleId = form.DevelopmentRoleId
	}

	return form
}

func (form *ProjectAssigneeForm) validateJoinDate() *ProjectAssigneeForm {
	field := form.FindAttrByCode("joinDate")
	field.ValidateRequired()
	field.ValidateFormat("1-2-2006", "%d-%m-%y")

	if field.IsClean() {
		form.ProjectAssignee.JoinDate = field.Time()
	}

	return form
}

func (form *ProjectAssigneeForm) validateLeaveDate() *ProjectAssigneeForm {
	field := form.FindAttrByCode("leaveDate")

	if form.LeaveDate != nil && *form.LeaveDate != "" && strings.TrimSpace(*form.LeaveDate) != "" {
		field.ValidateFormat("1-2-2006", "%d-%m-%y")

		joinDateTime := form.FindAttrByCode("joinDate").Time()

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
	userIdField := form.FindAttrByCode("userId")
	developmentIdField := form.FindAttrByCode("developmentRoleId")

	if userIdField.IsClean() && developmentIdField.IsClean() {
		presentedProjectAssignee := models.ProjectAssignee{
			ProjectId:         form.Project.Id,
			UserId:            form.UserId,
			DevelopmentRoleId: form.DevelopmentRoleId,
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
