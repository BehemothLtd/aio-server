package globalTypes

import (
	"aio-server/models"
	"aio-server/pkg/helpers"
	"context"

	graphql "github.com/graph-gophers/graphql-go"
	"gorm.io/gorm"
)

type ProjectSprintType struct {
	Ctx *context.Context
	Db  *gorm.DB

	ProjectSprint *models.ProjectSprint
}

func (pst *ProjectSprintType) ID(context.Context) *graphql.ID {
	return helpers.GqlIDP(pst.ProjectSprint.Id)
}

func (pst *ProjectSprintType) Title(context.Context) string {
	return pst.ProjectSprint.Title
}

func (pst *ProjectSprintType) ProjectId(context.Context) *graphql.ID {
	return helpers.GqlIDP(pst.ProjectSprint.ProjectId)
}

func (pst *ProjectSprintType) StartDate(context.Context) *graphql.Time {
	return helpers.GqlTimePointer(&pst.ProjectSprint.StartDate)
}

func (pst *ProjectSprintType) EndDate(context.Context) *graphql.Time {
	return helpers.GqlTimePointer(pst.ProjectSprint.EndDate)
}

func (pst *ProjectSprintType) UpdatedAt(context.Context) *graphql.Time {
	return helpers.GqlTimePointer(&pst.ProjectSprint.UpdatedAt)
}
func (pst *ProjectSprintType) CreatedAt(context.Context) *graphql.Time {
	return helpers.GqlTimePointer(&pst.ProjectSprint.CreatedAt)
}

func (pst *ProjectSprintType) Archived(context.Context) bool {
	return pst.ProjectSprint.Archived
}

func (pst *ProjectSprintType) LockVersion(context.Context) int32 {
	return pst.ProjectSprint.LockVersion
}

func (pst *ProjectSprintType) Active(context.Context) bool {
	return pst.ProjectSprint.Id == *pst.ProjectSprint.Project.CurrentSprintId
}
