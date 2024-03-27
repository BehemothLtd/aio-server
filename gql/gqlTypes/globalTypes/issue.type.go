package globalTypes

import (
	"aio-server/models"
	"aio-server/pkg/helpers"
	"context"

	"github.com/graph-gophers/graphql-go"
	"gorm.io/gorm"
)

type IssueType struct {
	Ctx *context.Context
	DB  *gorm.DB

	Issue *models.Issue
}

func (it *IssueType) ID(ctx context.Context) *graphql.ID {
	return helpers.GqlIDP(it.Issue.Id)
}

func (it *IssueType) Code(ctx context.Context) *string {
	return &it.Issue.Code
}

func (it *IssueType) Title(ctx context.Context) *string {
	return &it.Issue.Title
}

func (it *IssueType) Description(ctx context.Context) *string {
	return &it.Issue.Description
}

func (it *IssueType) Archived(ctx context.Context) *bool {
	return &it.Issue.Archived
}

func (it *IssueType) Archiveable(ctx context.Context) *bool {
	// return &it.Issue.IssueStatus
	return nil
}

func (it *IssueType) IssueStatusId(ctx context.Context) *graphql.ID {
	return nil
}

func (it *IssueType) IssueType(ctx context.Context) *string {
	issueType := it.Issue.IssueType.String()

	return &issueType
}

func (it *IssueType) Priority(ctx context.Context) *string {
	priority := it.Issue.Priority.String()

	return &priority
}

func (it *IssueType) Status(ctx context.Context) *string {
	status := it.Issue.IssueType.String()
	return &status
}

func (it *IssueType) Position(ctx context.Context) *int32 {
	return &it.Issue.Position
}

func (it *IssueType) CreatedAt(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(&it.Issue.CreatedAt)
}

func (it *IssueType) UpdatedAt(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(&it.Issue.UpdatedAt)
}

func (it *IssueType) Deadline(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(&it.Issue.Deadline)
}

func (it *IssueType) StartedAt(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(&it.Issue.StartDate)
}

func (it *IssueType) Children(ctx context.Context) *[]*IssueType {
	children := make([]*IssueType, len(it.Issue.Children))

	for i, issue := range it.Issue.Children {
		children[i] = &IssueType{
			Issue: &issue,
		}
	}

	return &children
}

func (it *IssueType) ParentId(ctx context.Context) *graphql.ID {
	if it.Issue.ParentId != nil {
		return helpers.GqlIDP(*it.Issue.ParentId)
	}

	return nil
}

func (it *IssueType) Parent(ctx context.Context) *IssueType {
	if it.Issue.ParentId != nil {
		return &IssueType{
			Issue: it.Issue.Parent,
		}
	}

	return nil
}

func (it *IssueType) ProjectId(ctx context.Context) *graphql.ID {
	return helpers.GqlIDP(it.Issue.ProjectId)
}

func (it *IssueType) Project(ctx context.Context) *ProjectType {
	return &ProjectType{
		Project: &it.Issue.Project,
	}
}

func (it *IssueType) CreatorId(ctx context.Context) *graphql.ID {
	return helpers.GqlIDP(it.Issue.CreatorId)
}

func (it *IssueType) Creator(ctx context.Context) *UserType {
	return &UserType{
		User: &it.Issue.Creator,
	}
}

func (it *IssueType) IssueAssignees(ctx context.Context) *[]*IssueAssigneeType {
	issueAssignees := make([]*IssueAssigneeType, len(it.Issue.IssueAssignees))

	for i, issueAssignee := range it.Issue.IssueAssignees {
		issueAssignees[i] = &IssueAssigneeType{
			IssueAssignee: &issueAssignee,
		}
	}

	return &issueAssignees
}

func (it *IssueType) Assignees(ctx context.Context) *[]*UserType {
	assignees := make([]*UserType, len(it.Issue.IssueAssignees))

	for i, issueAssignee := range it.Issue.IssueAssignees {
		assignees[i] = &UserType{
			User: &issueAssignee.User,
		}
	}

	return &assignees
}

func (it *IssueType) ProjectSprintId(ctx context.Context) *graphql.ID {
	if it.Issue.ProjectSprintId != nil {
		return helpers.GqlIDP(*it.Issue.ProjectSprintId)
	}

	return nil
}

func (it *IssueType) ProjectSprint(ctx context.Context) *ProjectSprintType {
	if it.Issue.ProjectSprintId != nil {
		return &ProjectSprintType{
			ProjectSprint: it.Issue.ProjectSprint,
		}
	}

	return nil
}
