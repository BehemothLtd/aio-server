package globalTypes

import (
	"aio-server/models"
	"aio-server/pkg/helpers"
	"context"

	"github.com/graph-gophers/graphql-go"
)

type IssueStatusType struct {
	IssueStatus *models.IssueStatus
}

func (is *IssueStatusType) ID(ctx context.Context) *graphql.ID {
	return helpers.GqlIDP(is.IssueStatus.Id)
}

func (is *IssueStatusType) Color(ctx context.Context) *string {
	return &is.IssueStatus.Color
}

func (is *IssueStatusType) StatusType(ctx context.Context) *string {
	statusType := is.IssueStatus.StatusType.String()

	return &statusType
}

func (is *IssueStatusType) Title(ctx context.Context) *string {
	return &is.IssueStatus.Title
}

func (is *IssueStatusType) CreatedAt(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(&is.IssueStatus.CreatedAt)
}

func (is *IssueStatusType) UpdatedAt(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(&is.IssueStatus.UpdatedAt)
}

func (is *IssueStatusType) LockVersion(ctx context.Context) *int32 {
	return &is.IssueStatus.LockVersion
}
