package globalTypes

import (
	"aio-server/models"
	"aio-server/pkg/helpers"
	"context"

	graphql "github.com/graph-gophers/graphql-go"
)

type ProjectType struct {
	Project *models.Project
}

func (pt *ProjectType) ID(ctx context.Context) *graphql.ID {
	return helpers.GqlIDP(pt.Project.Id)
}

func (pt *ProjectType) Name(ctx context.Context) string {
	return pt.Project.Name
}

func (pt *ProjectType) Code(ctx context.Context) string {
	return pt.Project.Code
}

func (pt *ProjectType) ProjectType(ctx context.Context) *string {
	value := pt.Project.ProjectType.String()
	return &value
}

func (pt *ProjectType) ClientId(ctx context.Context) *graphql.ID {
	return helpers.GqlIDP(pt.Project.ClientId)
}

func (pt *ProjectType) JiraId(ctx context.Context) *graphql.ID {
	return helpers.GqlIDP(pt.Project.JiraId)
}

func (pt *ProjectType) SprintDuration(ctx context.Context) *graphql.ID {
	return helpers.GqlIDP(pt.Project.SprintDuration)
}

func (pt *ProjectType) Description(ctx context.Context) *string {
	return &pt.Project.Description
}

func (pt *ProjectType) CurrentSprintId(ctx context.Context) *graphql.ID {
	return helpers.GqlIDP(pt.Project.CurrentSprintId)
}

func (pt *ProjectType) ProjectPriority(ctx context.Context) string {
	return pt.Project.ProjectPriority.String()
}

func (pt *ProjectType) State(ctx context.Context) string {
	return pt.Project.State.String()
}

func (pt *ProjectType) ActiveAt(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(pt.Project.ActiveAt)
}

func (pt *ProjectType) InactiveAt(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(pt.Project.InactiveAt)
}

func (pt *ProjectType) StartedAt(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(pt.Project.StartedAt)
}

func (pt *ProjectType) EndedAt(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(pt.Project.EndedAt)
}

func (pt *ProjectType) LockVersion(ctx context.Context) int32 {
	return pt.Project.LockVersion
}

func (pt *ProjectType) CreatedAt(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(pt.Project.CreatedAt)
}

func (pt *ProjectType) UpdatedAt(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(pt.Project.UpdatedAt)
}
