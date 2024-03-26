package validators

import (
	"aio-server/exceptions"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/repository"
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

	// if err := form.Repo.Create(form.Project); err != nil {
	// 	return err
	// }

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
				Code: "issueStatus",
			},
			Value: helpers.GetStringOrDefault(form.IssueStatus),
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
	form.summaryErrors()

	if form.Errors != nil {
		return exceptions.NewUnprocessableContentError("", form.Errors)
	}

	return nil
}
