package globalTypes

import (
	"aio-server/models"
	"aio-server/pkg/constants"
	"aio-server/pkg/helpers"
	"context"

	"github.com/graph-gophers/graphql-go"

	"gorm.io/gorm"
)

type WorkingTimelogType struct {
	Ctx *context.Context
	DB  *gorm.DB

	WorkingTimelog *models.WorkingTimelog
}

func (wtt *WorkingTimelogType) ID(ctx context.Context) *graphql.ID {
	return helpers.GqlIDP(wtt.WorkingTimelog.Id)
}

func (wtt *WorkingTimelogType) Description(ctx context.Context) string {
	return wtt.WorkingTimelog.Description
	// return nil
}

func (wtt *WorkingTimelogType) Minutes(ctx context.Context) int32 {
	return int32(wtt.WorkingTimelog.Minutes)
	// return nil
}

func (wtt *WorkingTimelogType) LoggedAt(ctx context.Context) string {
	loggedAt := helpers.GqlTimePointer(&wtt.WorkingTimelog.LoggedAt)

	return loggedAt.Format(constants.DDMMYYY_DateFormat)
}

func (wtt *WorkingTimelogType) CreatedAt(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(&wtt.WorkingTimelog.CreatedAt)
}

func (wtt *WorkingTimelogType) UpdatedAt(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(&wtt.WorkingTimelog.UpdatedAt)
}

func (wtt *WorkingTimelogType) User(ctx context.Context) *UserType {
	result := UserType{User: &wtt.WorkingTimelog.User}

	return &result
}

func (wtt *WorkingTimelogType) Project(ctx context.Context) *ProjectType {
	result := ProjectType{Project: &wtt.WorkingTimelog.Project}

	return &result
}

func (wtt *WorkingTimelogType) Issue(ctx context.Context) *IssueType {
	result := IssueType{Issue: &wtt.WorkingTimelog.Issue}

	return &result
}
