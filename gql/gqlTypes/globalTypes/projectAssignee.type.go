package globalTypes

import (
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/pkg/systems"
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
	return helpers.GqlTimePointer(tt.ProjectAssignee.JoinDate)
}

func (tt *ProjectAssigneeType) LeaveDate(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(tt.ProjectAssignee.LeaveDate)
}

func (tt *ProjectAssigneeType) UserId(ctx context.Context) *graphql.ID {
	return helpers.GqlIDP(tt.ProjectAssignee.UserId)
}

func (tt *ProjectAssigneeType) DevelopmentRoleId(ctx context.Context) *graphql.ID {
	return helpers.GqlIDP(tt.ProjectAssignee.DevelopmentRoleId)
}

func (tt *ProjectAssigneeType) DevelopmentRole(ctx context.Context) *DevelopentRoleType {
	developmentRole := systems.FindDevelopmentRoleById(tt.ProjectAssignee.DevelopmentRoleId)

	return &DevelopentRoleType{
		DevelopmentRole: developmentRole,
	}
}

func (tt *ProjectAssigneeType) Title(ctx context.Context) *string {
	if developmentRole := systems.FindDevelopmentRoleById(tt.ProjectAssignee.DevelopmentRoleId); developmentRole != nil {
		return &developmentRole.Title
	} else {
		return nil
	}
}

func (tt *ProjectAssigneeType) Name(ctx context.Context) *string {
	return &tt.ProjectAssignee.User.Name
}

func (tt *ProjectAssigneeType) LockVersion(ctx context.Context) *int32 {
	return &tt.ProjectAssignee.LockVersion
}
