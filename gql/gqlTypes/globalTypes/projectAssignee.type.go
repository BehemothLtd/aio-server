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

func (pat *ProjectAssigneeType) JoinDate(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(pat.ProjectAssignee.JoinDate)
}

func (pat *ProjectAssigneeType) LeaveDate(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(pat.ProjectAssignee.LeaveDate)
}

func (pat *ProjectAssigneeType) UserId(ctx context.Context) *graphql.ID {
	return helpers.GqlIDP(pat.ProjectAssignee.UserId)
}

func (pat *ProjectAssigneeType) DevelopmentRoleId(ctx context.Context) *graphql.ID {
	return helpers.GqlIDP(pat.ProjectAssignee.DevelopmentRoleId)
}

func (pat *ProjectAssigneeType) DevelopmentRole(ctx context.Context) *DevelopentRoleType {
	developmentRole := systems.FindDevelopmentRoleById(pat.ProjectAssignee.DevelopmentRoleId)

	return &DevelopentRoleType{
		DevelopmentRole: developmentRole,
	}
}

func (pat *ProjectAssigneeType) Title(ctx context.Context) *string {
	if developmentRole := systems.FindDevelopmentRoleById(pat.ProjectAssignee.DevelopmentRoleId); developmentRole != nil {
		return &developmentRole.Title
	} else {
		return nil
	}
}

func (pat *ProjectAssigneeType) Name(ctx context.Context) *string {
	return &pat.ProjectAssignee.User.Name
}

func (pat *ProjectAssigneeType) LockVersion(ctx context.Context) *int32 {
	return &pat.ProjectAssignee.LockVersion
}

func (pat *ProjectAssigneeType) Project(ctx context.Context) *ProjectType {
	return &ProjectType{
		Project: &pat.ProjectAssignee.Project,
	}
}

func (pat *ProjectAssigneeType) User(ctx context.Context) *UserType {
	return &UserType{
		User: &pat.ProjectAssignee.User,
	}
}
