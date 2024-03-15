package insightTypes

import (
	"aio-server/models"
	"aio-server/pkg/helpers"
	"context"

	"github.com/graph-gophers/graphql-go"
	"gorm.io/gorm"
)

type IssueType struct {
	Ctx *context.Context
	DB  *gorm.DB

	Issue *models.Issue
}

func (it *IssueType) ID(ctx context.Context) *graphql.ID {
	return helpers.GqlIDP(it.Issue.Id)
}

func (it *IssueType) Code(ctx context.Context) *string {
	return &it.Issue.Code
}

func (it *IssueType) Title(ctx context.Context) *string {
	return &it.Issue.Title
}
