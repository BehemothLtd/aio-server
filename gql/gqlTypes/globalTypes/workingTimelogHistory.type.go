package globalTypes

import (
	"aio-server/models"
	"aio-server/pkg/constants"
	"aio-server/pkg/helpers"
	"context"

	"github.com/graph-gophers/graphql-go"
	"gorm.io/gorm"
)

type WorkingTimelogHistoryType struct {
	Ctx                   *context.Context
	DB                    *gorm.DB
	WorkingTimelogHistory *models.WorkingTimelogHistory
}

func (wt *WorkingTimelogHistoryType) ID(ctx context.Context) *graphql.ID {
	return helpers.GqlIDP(wt.WorkingTimelogHistory.Id)
}

func (wt *WorkingTimelogHistoryType) IssueName(ctx context.Context) *string {
	return &wt.WorkingTimelogHistory.IssueName
}

func (wt *WorkingTimelogHistoryType) IssueDescription(ctx context.Context) *string {
	return &wt.WorkingTimelogHistory.IssueDescription
}

func (wt *WorkingTimelogHistoryType) IssueId(ctx context.Context) *graphql.ID {
	return helpers.GqlIDP(wt.WorkingTimelogHistory.IssueId)
}

func (wt *WorkingTimelogHistoryType) ProjectId(ctx context.Context) *graphql.ID {
	return helpers.GqlIDP(wt.WorkingTimelogHistory.ProjectId)
}

func (wt *WorkingTimelogHistoryType) Minutes(ctx context.Context) *int32 {
	return &wt.WorkingTimelogHistory.Minutes
}

func (wt *WorkingTimelogHistoryType) LoggedAt(ctx context.Context) string {
	loggedAt := helpers.GqlTimePointer(&wt.WorkingTimelogHistory.LoggedAt)
	return loggedAt.Format(constants.YYYYMMDD_DateFormat)
}
