package globalTypes

import (
	"aio-server/models"
	"aio-server/pkg/helpers"
	"context"

	graphql "github.com/graph-gophers/graphql-go"
)

type IssueAssigneeType struct {
	IssueAssignee *models.IssueAssignee
}

func (iat *IssueAssigneeType) ID(ctx context.Context) *graphql.ID {
	return helpers.GqlIDP(iat.IssueAssignee.Id)
}

func (iat *IssueAssigneeType) UserId(ctx context.Context) int32 {
	return iat.IssueAssignee.UserId
}

func (iat *IssueAssigneeType) DevelopmentRoleId(ctx context.Context) int32 {
	return iat.IssueAssignee.DevelopmentRoleId
}
