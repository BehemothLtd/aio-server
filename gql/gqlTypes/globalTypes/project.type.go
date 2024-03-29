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

func (tt *ProjectType) Name(ctx context.Context) *string {
	return &tt.Project.Name
}

func (tt *ProjectType) Code(ctx context.Context) *string {
	return &tt.Project.Code
}

func (tt *ProjectType) ProjectType(ctx context.Context) *string {
	projectType := tt.Project.ProjectType.String()
	return &projectType
}

func (tt *ProjectType) ProjectPriority(ctx context.Context) *string {
	projectPriority := tt.Project.ProjectPriority.String()
	return &projectPriority
}

func (tt *ProjectType) State(ctx context.Context) string {
	return tt.Project.State.String()
}

func (tt *ProjectType) ActivedAt(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(tt.Project.ActivedAt)
}

func (tt *ProjectType) InactivedAt(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(tt.Project.InactivedAt)
}

func (tt *ProjectType) StartedAt(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(tt.Project.StartedAt)
}

func (tt *ProjectType) EndedAt(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(tt.Project.EndedAt)
}

func (tt *ProjectType) CreatedAt(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(&tt.Project.CreatedAt)
}

func (tt *ProjectType) UpdatedAt(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(&tt.Project.UpdatedAt)
}

func (tt *ProjectType) SprintDuration(ctx context.Context) *int32 {
	if tt.Project.ProjectType == enums.ProjectTypeKanban {
		return nil
	} else {
		return helpers.Int32Pointer(int32(*tt.Project.SprintDuration))
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

// func (pt *ProjectType) Files(ctx context.Context) *AttachmentType {
// 	if pt.Project.Logo == nil {
// 		return nil
// 	}

// 	return &AttachmentType{
// 		Attachment: pt.Project.Logo,
// 	}
// }
