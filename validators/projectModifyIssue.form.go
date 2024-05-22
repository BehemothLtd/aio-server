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
	updates map[string]interface{}
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
		updates:                     map[string]interface{}{},
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
		if err := form.Repo.Update(form.Issue, form.updates); err != nil {
			return err
		}
	}

	return nil
}

func (form *ProjectModifyIssueForm) assignAttributes() {
	form.AddAttributes(
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "Title",
			},
			Value: helpers.GetStringOrDefault(form.Title),
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "Description",
			},
			Value: helpers.GetStringOrDefault(form.Description),
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "IssueType",
			},
			Value: helpers.GetStringOrDefault(form.IssueType),
		},
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "Priority",
			},
			Value: helpers.GetStringOrDefault(form.Priority),
		},
		&BoolAttribute{
			FieldAttribute: FieldAttribute{
				Code: "Archived",
			},
			Value: helpers.GetBoolOrDefault(form.Archived),
		},
		&TimeAttribute{
			FieldAttribute: FieldAttribute{
				Code: "Deadline",
			},
			Value: helpers.GetStringOrDefault(form.Deadline),
		},
		&TimeAttribute{
			FieldAttribute: FieldAttribute{
				Code: "StartDate",
			},
			Value: helpers.GetStringOrDefault(form.StartDate),
		},
		&IntAttribute[int32]{
			FieldAttribute: FieldAttribute{
				Code: "IssueStatusId",
			},
			Value: helpers.GetInt32OrDefault(form.IssueStatusId),
		},
		&IntAttribute[int32]{
			FieldAttribute: FieldAttribute{
				Code: "ParentId",
			},
			Value: helpers.GetInt32OrDefault(form.ParentId),
		},
		&IntAttribute[int32]{
			FieldAttribute: FieldAttribute{
				Code: "ProjectSprintId",
			},
			Value: helpers.GetInt32OrDefault(form.ProjectSprintId),
		},
		&SliceAttribute[insightInputs.IssueAssigneeInputForIssueCreate]{
			FieldAttribute: FieldAttribute{
				Code: "IssueAssignees",
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
	code := "Title"
	field := form.FindAttrByCode(code)
	field.ValidateRequired()

	field.ValidateMin(interface{}(int64(5)))
	field.ValidateMax(interface{}(int64(constants.MaxStringLength)))

	if field.IsClean() {
		form.Issue.Title = *form.Title
		form.updates[code] = *form.Title
	}

	return form
}

func (form *ProjectModifyIssueForm) validateDescription() *ProjectModifyIssueForm {
	code := "Description"
	field := form.FindAttrByCode(code)
	field.ValidateRequired()

	field.ValidateMax(interface{}(int64(constants.MaxLongTextLength)))

	if field.IsClean() {
		form.Issue.Description = *form.Description
		form.updates[code] = *form.Description
	}

	return form
}

func (form *ProjectModifyIssueForm) validateIssueType() *ProjectModifyIssueForm {
	code := "IssueType"
	field := form.FindAttrByCode(code)
	field.ValidateRequired()

	if field.IsClean() {
		if issueTypeEnum, err := enums.ParseIssueType(*form.IssueType); err != nil {
			field.AddError("is invalid type")
		} else {
			form.Issue.IssueType = issueTypeEnum
			form.updates[code] = issueTypeEnum
		}
	}

	return form
}

func (form *ProjectModifyIssueForm) validatePriority() *ProjectModifyIssueForm {
	code := "Priority"
	field := form.FindAttrByCode(code)
	field.ValidateRequired()

	if field.IsClean() {
		if priority, err := enums.ParseIssuePriority(*form.Priority); err != nil {
			field.AddError("is invalid priority")
		} else {
			form.Issue.Priority = priority
			form.updates[code] = priority
		}
	}

	return form
}

func (form *ProjectModifyIssueForm) validateArchived() *ProjectModifyIssueForm {
	archived := helpers.GetBoolOrDefault(form.Archived)
	form.Issue.Archived = archived
	form.updates["Archived"] = archived

	return form
}

func (form *ProjectModifyIssueForm) validateDeadlineAndStartDate() *ProjectModifyIssueForm {
	deadlineField := form.FindAttrByCode("Deadline")
	startDateField := form.FindAttrByCode("StartDate")

	if form.Deadline != nil && strings.TrimSpace(*form.Deadline) != "" {
		deadlineField.ValidateFormat(constants.DDMMYYYY_DateSplashFormat, constants.HUMAN_DDMMYYYY_DateSplashFormat)
	}

	if form.StartDate != nil && strings.TrimSpace(*form.StartDate) != "" {
		startDateField.ValidateFormat(constants.DDMMYYYY_DateSplashFormat, constants.HUMAN_DDMMYYYY_DateSplashFormat)
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
		form.Issue.Deadline = &deadline
		form.updates["Deadline"] = &deadline
	}

	if startDateField.IsClean() {
		form.Issue.StartDate = &startDate
		form.updates["StartDate"] = &startDate
	}

	return form
}

func (form *ProjectModifyIssueForm) validateIssueStatusId() *ProjectModifyIssueForm {
	code := "IssueStatusId"
	field := form.FindAttrByCode(code)
	field.ValidateRequired()

	if field.IsClean() {
		if foundIdx := slices.IndexFunc(form.Project.ProjectIssueStatuses, func(pis *models.ProjectIssueStatus) bool { return pis.IssueStatusId == *form.IssueStatusId }); foundIdx == -1 {
			field.AddError("is invalid status")
		} else {
			form.Issue.IssueStatusId = *form.IssueStatusId
			form.updates[code] = *form.IssueStatusId
		}
	}

	return form
}

func (form *ProjectModifyIssueForm) validateParentId() *ProjectModifyIssueForm {
	code := "ParentId"
	field := form.FindAttrByCode(code)

	if form.ParentId != nil {
		if foundIdx := slices.IndexFunc(form.Project.Issues, func(issue models.Issue) bool { return issue.Id == *form.ParentId }); foundIdx == -1 {
			field.AddError("is invalid Parent Issue")
		} else {
			form.Issue.ParentId = form.ParentId
			form.updates[code] = form.ParentId
		}
	}

	return form
}

func (form *ProjectModifyIssueForm) validateProjectSprintId() *ProjectModifyIssueForm {
	code := "ProjectSprintId"
	field := form.FindAttrByCode(code)

	if form.ProjectSprintId != nil {
		if foundIdx := slices.IndexFunc(form.Project.ProjectSprints, func(ps models.ProjectSprint) bool { return ps.Id == *form.ProjectSprintId }); foundIdx == -1 {
			field.AddError("is invalid Sprint")
		} else {
			form.Issue.ProjectSprintId = form.ProjectSprintId
			form.updates[code] = form.ProjectSprintId
		}
	}

	return form
}

func (form *ProjectModifyIssueForm) validateIssueAssignees() *ProjectModifyIssueForm {
	fieldKey := "IssueAssignees"
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
				form.AddErrorDirectlyToField(form.NestedFieldKey(fieldKey, i, "UserId"), []interface{}{"is already has same role"})
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
