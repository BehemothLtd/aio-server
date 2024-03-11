package globalTypes

import (
	"aio-server/models"
	"aio-server/pkg/helpers"
	"context"

	"github.com/graph-gophers/graphql-go"
	"gorm.io/gorm"
)

type ProjectAssigneeType struct {
	Ctx *context.Context
	Db  *gorm.DB

	ProjectAssignee *models.ProjectAssignee
}

func (pt *ProjectAssigneeType) ID(ctx context.Context) *graphql.ID {
	return helpers.GqlIDP(pt.ProjectAssignee.Id)
}

func (pt *ProjectAssigneeType) Active(ctx context.Context) bool {
	return pt.ProjectAssignee.Active
}

func (tt *ProjectAssigneeType) JoinDate(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(*tt.ProjectAssignee.JoinDate)
}

func (tt *ProjectAssigneeType) LeaveDate(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(*tt.ProjectAssignee.LeaveDate)
}
