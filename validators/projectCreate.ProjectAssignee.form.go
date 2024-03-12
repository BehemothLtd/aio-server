package validators

import (
	"aio-server/exceptions"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/pkg/systems"
	"aio-server/repository"
	"time"
)

type ProjectCreateProjectAssigneeForm struct {
	Form
	insightInputs.ProjectAssigneeInputForProjectCreate

	ProjectAssignee *models.ProjectAssignee
	UserRepo        *repository.UserRepository
}

func NewProjectCreateProjectAssigneeFormValidator(
	input *insightInputs.ProjectAssigneeInputForProjectCreate,
	userRepo *repository.UserRepository,
	projectAssignee *models.ProjectAssignee,
) ProjectCreateProjectAssigneeForm {
	form := ProjectCreateProjectAssigneeForm{
		Form:                                 Form{},
		ProjectAssigneeInputForProjectCreate: *input,
		ProjectAssignee:                      projectAssignee,
		UserRepo:                             userRepo,
	}
	form.assignAttributes()

	return form
}

func (form *ProjectCreateProjectAssigneeForm) assignAttributes() {
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
			Value: helpers.GetInt32OrDefault(&form.DevelopmentRoleId),
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "joinDate",
			},
			Value: helpers.GetStringOrDefault(&form.JoinDate),
		},
	)
}

func (form *ProjectCreateProjectAssigneeForm) Validate() exceptions.ResourceModificationError {
	form.validateUserId().
		validateDevelopmentId().
		validateJoinDate().
		summaryErrors()

	if form.Errors != nil {
		return form.Errors
	}

	return nil
}

func (form *ProjectCreateProjectAssigneeForm) validateUserId() *ProjectCreateProjectAssigneeForm {
	field := form.FindAttrByCode("userId")
	field.ValidateRequired()

	if form.UserId != 0 {
		if err := form.UserRepo.Find(&models.User{Id: form.UserId}); err != nil {
			field.AddError("is invalid")
		}
	}

	return form
}

func (form *ProjectCreateProjectAssigneeForm) validateDevelopmentId() *ProjectCreateProjectAssigneeForm {
	field := form.FindAttrByCode("developmentRoleId")
	field.ValidateRequired()

	if form.DevelopmentRoleId != 0 {
		if developmentRole := systems.FindDevelopmentRoleById(form.DevelopmentRoleId); developmentRole == nil {
			field.AddError("is invalid")
		}
	}

	return form
}

func (form *ProjectCreateProjectAssigneeForm) validateJoinDate() *ProjectCreateProjectAssigneeForm {
	field := form.FindAttrByCode("joinDate")
	field.ValidateRequired()

	if form.JoinDate != "" {
		// TODO: add Time FieldType
		if date, err := time.Parse("1-2-2006", form.JoinDate); err != nil {
			field.AddError("need to be formatted as %d-%m-%y")
		} else {
			form.ProjectAssignee.JoinDate = &date
		}
	}

	return form
}
