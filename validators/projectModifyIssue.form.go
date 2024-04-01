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
	"slices"
	"strings"
)

type ProjectModifyIssueForm struct {
	Form
	insightInputs.ProjectModifyIssueFormInput
	Project models.Project
	Repo    repository.IssueRepository
	Issue   *models.Issue
}

func NewProjectModifyIssueFormValidator(
	input *insightInputs.ProjectModifyIssueFormInput,
	repo repository.IssueRepository,
	project models.Project,
	issue *models.Issue,
) ProjectModifyIssueForm {
	form := ProjectModifyIssueForm{
		Form:                        Form{},
		ProjectModifyIssueFormInput: *input,
		Project:                     project,
		Repo:                        repo,
		Issue:                       issue,
	}
	form.assignAttributes()

	return form
}

func (form *ProjectModifyIssueForm) Save() error {
	if err := form.validate(); err != nil {
		return err
	}

	if form.Issue.Id == 0 {
		if err := form.Repo.Create(form.Issue); err != nil {
			return err
		}
	} else {
		if err := form.Repo.Update(form.Issue); err != nil {
			return err
		}
	}

	return nil
}

func (form *ProjectModifyIssueForm) assignAttributes() {
	form.AddAttributes(
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "title",
			},
			Value: helpers.GetStringOrDefault(form.Title),
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "description",
			},
			Value: helpers.GetStringOrDefault(form.Description),
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "issueType",
			},
			Value: helpers.GetStringOrDefault(form.IssueType),
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "priority",
			},
			Value: helpers.GetStringOrDefault(form.Priority),
		},
		&BoolAttribute{
			FieldAttribute: FieldAttribute{
				Code: "archived",
			},
			Value: helpers.GetBoolOrDefault(form.Archived),
		},
		&TimeAttribute{
			FieldAttribute: FieldAttribute{
				Code: "deadline",
			},
			Value: helpers.GetStringOrDefault(form.Deadline),
		},
		&TimeAttribute{
			FieldAttribute: FieldAttribute{
				Code: "startDate",
			},
			Value: helpers.GetStringOrDefault(form.StartDate),
		},
		&IntAttribute[int32]{
			FieldAttribute: FieldAttribute{
				Code: "issueStatusId",
			},
			Value: helpers.GetInt32OrDefault(form.IssueStatusId),
		},
		&IntAttribute[int32]{
			FieldAttribute: FieldAttribute{
				Code: "parentId",
			},
			Value: helpers.GetInt32OrDefault(form.ParentId),
		},
		&IntAttribute[int32]{
			FieldAttribute: FieldAttribute{
				Code: "projectSprintId",
			},
			Value: helpers.GetInt32OrDefault(form.ProjectSprintId),
		},
		&SliceAttribute[insightInputs.IssueAssigneeInputForIssueCreate]{
			FieldAttribute: FieldAttribute{
				Code: "issueAssignees",
			},
			Value: form.IssueAssignees,
		},
	)
}

func (form *ProjectModifyIssueForm) validate() error {
	form.validateTitle().
		validateDescription().
		validateIssueType().
		validatePriority().
		validateArchived().
		validateDeadlineAndStartDate().
		validateIssueStatusId().
		validateParentId().
		validateProjectSprintId().
		validateIssueAssignees().
		summaryErrors()

	if form.Errors != nil {
		return exceptions.NewUnprocessableContentError("", form.Errors)
	}

	return nil
}

func (form *ProjectModifyIssueForm) validateTitle() *ProjectModifyIssueForm {
	field := form.FindAttrByCode("title")
	field.ValidateRequired()

	field.ValidateMin(interface{}(int64(5)))
	field.ValidateMax(interface{}(int64(constants.MaxStringLength)))

	if field.IsClean() {
		form.Issue.Title = *form.Title
	}

	return form
}

func (form *ProjectModifyIssueForm) validateDescription() *ProjectModifyIssueForm {
	field := form.FindAttrByCode("description")
	field.ValidateRequired()

	field.ValidateMax(interface{}(int64(constants.MaxLongTextLength)))

	if field.IsClean() {
		form.Issue.Description = *form.Description
	}

	return form
}

func (form *ProjectModifyIssueForm) validateIssueType() *ProjectModifyIssueForm {
	field := form.FindAttrByCode("issueType")
	field.ValidateRequired()

	if field.IsClean() {
		if issueTypeEnum, err := enums.ParseIssueType(*form.IssueType); err != nil {
			field.AddError("is invalid type")
		} else {
			form.Issue.IssueType = issueTypeEnum
		}
	}

	return form
}

func (form *ProjectModifyIssueForm) validatePriority() *ProjectModifyIssueForm {
	field := form.FindAttrByCode("priority")
	field.ValidateRequired()

	if field.IsClean() {
		if priority, err := enums.ParseIssuePriority(*form.Priority); err != nil {
			field.AddError("is invalid priority")
		} else {
			form.Issue.Priority = priority
		}
	}

	return form
}

func (form *ProjectModifyIssueForm) validateArchived() *ProjectModifyIssueForm {
	form.Issue.Archived = helpers.GetBoolOrDefault(form.Archived)

	return form
}

func (form *ProjectModifyIssueForm) validateDeadlineAndStartDate() *ProjectModifyIssueForm {
	deadlineField := form.FindAttrByCode("deadline")
	startDateField := form.FindAttrByCode("startDate")

	if form.Deadline != nil && strings.TrimSpace(*form.Deadline) != "" {
		deadlineField.ValidateFormat(constants.DDMMYYY_DateFormat, constants.HUMAN_DD_MM_YY_DateFormat)
	}

	if form.StartDate != nil && strings.TrimSpace(*form.StartDate) != "" {
		startDateField.ValidateFormat(constants.DDMMYYY_DateFormat, constants.HUMAN_DD_MM_YY_DateFormat)
	}

	if !deadlineField.IsClean() || !startDateField.IsClean() {
		return form
	}

	deadline := *deadlineField.Time()
	startDate := *startDateField.Time()

	if deadline.Before(startDate) {
		deadlineField.AddError("need to be after Start Date")
	}

	if deadlineField.IsClean() {
		form.Issue.Deadline = deadline
	}

	if startDateField.IsClean() {
		form.Issue.StartDate = startDate
	}

	return form
}

func (form *ProjectModifyIssueForm) validateIssueStatusId() *ProjectModifyIssueForm {
	field := form.FindAttrByCode("issueStatusId")
	field.ValidateRequired()

	if field.IsClean() {
		if foundIdx := slices.IndexFunc(form.Project.ProjectIssueStatuses, func(pis *models.ProjectIssueStatus) bool { return pis.IssueStatusId == *form.IssueStatusId }); foundIdx == -1 {
			field.AddError("is invalid status")
		} else {
			form.Issue.IssueStatusId = *form.IssueStatusId
		}
	}

	return form
}

func (form *ProjectModifyIssueForm) validateParentId() *ProjectModifyIssueForm {
	field := form.FindAttrByCode("parentId")

	if form.ParentId != nil {
		if foundIdx := slices.IndexFunc(form.Project.Issues, func(issue models.Issue) bool { return issue.Id == *form.ParentId }); foundIdx == -1 {
			field.AddError("is invalid Parent Issue")
		} else {
			form.Issue.ParentId = form.ParentId
		}
	}

	return form
}

func (form *ProjectModifyIssueForm) validateProjectSprintId() *ProjectModifyIssueForm {
	field := form.FindAttrByCode("projectSprintId")

	if form.ProjectSprintId != nil {
		if foundIdx := slices.IndexFunc(form.Project.ProjectSprints, func(ps models.ProjectSprint) bool { return ps.Id == *form.ProjectSprintId }); foundIdx == -1 {
			field.AddError("is invalid Sprint")
		} else {
			form.Issue.ProjectSprintId = form.ProjectSprintId
		}
	}

	return form
}

func (form *ProjectModifyIssueForm) validateIssueAssignees() *ProjectModifyIssueForm {
	fieldKey := "issueAssignees"
	field := form.FindAttrByCode(fieldKey)

	if form.IssueAssignees != nil && len(*form.IssueAssignees) > 0 {
		issueAssignees := []*models.IssueAssignee{}

		issueAssigneeRepo := repository.NewIssueAssigneeRepository(nil, database.Db)

		for i, issueAssigneeInput := range *form.IssueAssignees {
			userId := issueAssigneeInput.UserId
			developementRoleId := issueAssigneeInput.DevelopmentRoleId

			// Check duplicated in input
			if foundIdx := slices.IndexFunc(issueAssignees, func(ia *models.IssueAssignee) bool {
				return (userId != nil && ia.UserId == *userId && developementRoleId != nil && ia.DevelopmentRoleId == *developementRoleId)
			}); foundIdx != -1 {
				form.AddErrorDirectlyToField(form.NestedFieldKey(fieldKey, i, "userId"), []interface{}{"is already has same role"})
			} else {
				issueAssignee := models.IssueAssignee{}

				issueAssigneeForm := NewIssueCreateIssueAssigneeFormValidator(
					&issueAssigneeInput,
					issueAssigneeRepo,
					&issueAssignee,
					form.Project,
				)

				if err := issueAssigneeForm.Validate(); err != nil {
					form.AddNestedErrors(fieldKey, i, err)
				} else {
					// only push to final result when nested form has no error
					issueAssignees = append(issueAssignees, &issueAssignee)
				}
			}
		}

		if field.IsClean() {
			form.Issue.IssueAssignees = issueAssignees
		}
	}

	return form
}
