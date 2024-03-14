package globalTypes

import (
	"aio-server/models"
	"aio-server/pkg/helpers"
	"context"

	"github.com/graph-gophers/graphql-go"
	"gorm.io/gorm"
)

type ProjectIssueStatusType struct {
	Ctx *context.Context
	Db  *gorm.DB

	ProjectIssueStatus *models.ProjectIssueStatus
}

func (pt *ProjectIssueStatusType) ID(ctx context.Context) *graphql.ID {
	return helpers.GqlIDP(pt.ProjectIssueStatus.Id)
}

func (pt *ProjectIssueStatusType) Position(ctx context.Context) int32 {
	return int32(pt.ProjectIssueStatus.Position)
}

func (pt *ProjectIssueStatusType) IssueStatusId(ctx context.Context) *graphql.ID {
	return helpers.GqlIDP(pt.ProjectIssueStatus.IssueStatusId)
}

func (pt *ProjectIssueStatusType) IssueStatus(ctx context.Context) *IssueStatusType {
	return &IssueStatusType{
		IssueStatus: &pt.ProjectIssueStatus.IssueStatus,
	}
}
