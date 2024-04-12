package globalTypes

import (
	"aio-server/enums"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"context"

	"github.com/graph-gophers/graphql-go"
	"gorm.io/gorm"
)

type ProjectType struct {
	Ctx *context.Context
	Db  *gorm.DB

	Project *models.Project
}

func (pt *ProjectType) ID(ctx context.Context) *graphql.ID {
	return helpers.GqlIDP(pt.Project.Id)
}

func (pt *ProjectType) Name(ctx context.Context) *string {
	return &pt.Project.Name
}

func (pt *ProjectType) Code(ctx context.Context) *string {
	return &pt.Project.Code
}

func (pt *ProjectType) Description(ctx context.Context) *string {
	return pt.Project.Description
}

func (pt *ProjectType) ProjectType(ctx context.Context) *string {
	projectType := pt.Project.ProjectType.String()
	return &projectType
}

func (pt *ProjectType) ProjectPriority(ctx context.Context) *string {
	projectPriority := pt.Project.ProjectPriority.String()
	return &projectPriority
}

func (pt *ProjectType) State(ctx context.Context) string {
	return pt.Project.State.String()
}

func (pt *ProjectType) ActivedAt(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(pt.Project.ActivedAt)
}

func (pt *ProjectType) InactivedAt(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(pt.Project.InactivedAt)
}

func (pt *ProjectType) StartedAt(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(pt.Project.StartedAt)
}

func (pt *ProjectType) EndedAt(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(pt.Project.EndedAt)
}

func (pt *ProjectType) CreatedAt(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(&pt.Project.CreatedAt)
}

func (pt *ProjectType) UpdatedAt(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(&pt.Project.UpdatedAt)
}

func (pt *ProjectType) SprintDuration(ctx context.Context) *int32 {
	if pt.Project.ProjectType == enums.ProjectTypeKanban {
		return nil
	} else {
		return helpers.Int32Pointer(int32(*pt.Project.SprintDuration))
	}
}

func (pt *ProjectType) ClientId(ctx context.Context) *graphql.ID {
	return helpers.GqlIDP(pt.Project.ClientId)
}

func (pt *ProjectType) CurrentSprintId(ctx context.Context) *graphql.ID {
	return helpers.GqlIDP(pt.Project.CurrentSprintId)
}

func (pt *ProjectType) ProjectAssignees(ctx context.Context) *[]*ProjectAssigneeType {
	result := make([]*ProjectAssigneeType, len(pt.Project.ProjectAssignees))

	for i, projectAssignee := range pt.Project.ProjectAssignees {
		result[i] = &ProjectAssigneeType{
			ProjectAssignee: projectAssignee,
		}
	}

	return &result
}

func (pt *ProjectType) ProjectIssueStatuses(ctx context.Context) *[]*ProjectIssueStatusType {
	result := make([]*ProjectIssueStatusType, len(pt.Project.ProjectIssueStatuses))

	for i, projectIssueStatus := range pt.Project.ProjectIssueStatuses {
		result[i] = &ProjectIssueStatusType{
			ProjectIssueStatus: projectIssueStatus,
		}
	}

	return &result
}

func (pt *ProjectType) LockVersion(ctx context.Context) int32 {
	return pt.Project.LockVersion
}

func (pt *ProjectType) Logo(ctx context.Context) *AttachmentType {
	if pt.Project.Logo == nil {
		return nil
	}

	return &AttachmentType{
		Attachment: pt.Project.Logo,
	}
}

func (pt *ProjectType) Files(ctx context.Context) *[]*AttachmentType {
	if len(pt.Project.Files) == 0 {
		return nil
	}

	result := make([]*AttachmentType, len(pt.Project.Files))

	for i, aptachment := range pt.Project.Files {
		result[i] = &AttachmentType{
			Attachment: aptachment,
		}
	}

	return &result
}
