package globalTypes

import (
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
	return helpers.GqlTimePointer(*tt.Project.ActivedAt)
}

func (tt *ProjectType) InactivedAt(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(*tt.Project.InactivedAt)
}

func (tt *ProjectType) StartedAt(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(*tt.Project.StartedAt)
}

func (tt *ProjectType) EndedAt(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(*tt.Project.EndedAt)
}

func (tt *ProjectType) CreatedAt(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(tt.Project.CreatedAt)
}

func (tt *ProjectType) UpdatedAt(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(tt.Project.UpdatedAt)
}

func (tt *ProjectType) SprintDuration(ctx context.Context) *int32 {
	return nil
}

func (pt *ProjectType) ClientId(ctx context.Context) *graphql.ID {
	return helpers.GqlIDP(pt.Project.ClientId)
}

func (pt *ProjectType) CurrentSprintId(ctx context.Context) *graphql.ID {
	return helpers.GqlIDP(pt.Project.CurrentSprintId)
}

func (pt *ProjectType) ProjectAssignees(ctx context.Context) *[]*ProjectAssigneeType {
	return nil
}

func (pt *ProjectType) ProjectIssueStatuses(ctx context.Context) *[]*ProjectIssueStatusType {
	return nil
}
