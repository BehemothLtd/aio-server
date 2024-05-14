package insightInputs

import (
	"aio-server/gql/inputs/globalInputs"
	"aio-server/models"
	"aio-server/pkg/constants"
	"aio-server/pkg/helpers"
	"strings"
	"time"

	"github.com/graph-gophers/graphql-go"
)

type ProjectIssuesFrontQueryInput struct {
	TitleCont         *string
	CodeCont          *string
	IssueTypeEq       *string
	ProjectIdEq       *int32
	ProjectSprintIdEq *string
	DeadLineAtGteq    *string
	DeadLineAtLteq    *string
	UserIdIn          *[]int32
}

type ProjectIssuesQueryInput struct {
	TitleCont         *string
	CodeCont          *string
	IssueTypeEq       *string
	ProjectIdEq       *int32
	ProjectSprintIdEq *string
	DeadLineAtGteq    *time.Time
	DeadLineAtLteq    *time.Time
	UserIdIn          *[]int32
}

type ProjectIssuesInput struct {
	Id    graphql.ID
	Input *globalInputs.PagyInput
	Query *ProjectIssuesFrontQueryInput
}

func (input *ProjectIssuesInput) ToPaginationDataAndQuery() (ProjectIssuesQueryInput, models.PaginationData) {
	paginationData := input.Input.ToPaginationInput()

	query := ProjectIssuesQueryInput{}

	if input.Query != nil {
		if input.Query.TitleCont != nil && strings.TrimSpace(*input.Query.TitleCont) != "" {
			query.TitleCont = input.Query.TitleCont
		}

		if input.Query.CodeCont != nil && strings.TrimSpace(*input.Query.CodeCont) != "" {
			query.CodeCont = input.Query.CodeCont
		}

		if input.Query.IssueTypeEq != nil && strings.TrimSpace(*input.Query.IssueTypeEq) != "" {
			query.IssueTypeEq = input.Query.IssueTypeEq
		}

		if input.Query.ProjectSprintIdEq != nil {
			query.ProjectSprintIdEq = input.Query.ProjectSprintIdEq
		}

		if input.Query.DeadLineAtGteq != nil && strings.TrimSpace(*input.Query.DeadLineAtGteq) != "" {
			if timeValue, err := time.ParseInLocation(constants.YYYYMMDD_DateFormat, *input.Query.DeadLineAtGteq, time.Local); err != nil {
				query.DeadLineAtGteq = nil
			} else {
				beginOfDay := helpers.BeginningOfDay(&timeValue)
				query.DeadLineAtGteq = &beginOfDay
			}
		}

		if input.Query.DeadLineAtLteq != nil && strings.TrimSpace(*input.Query.DeadLineAtLteq) != "" {
			if timeValue, err := time.ParseInLocation(constants.YYYYMMDD_DateFormat, *input.Query.DeadLineAtLteq, time.Local); err != nil {
				query.DeadLineAtLteq = nil
			} else {
				endOfDay := helpers.EndOfDay(&timeValue)
				query.DeadLineAtLteq = &endOfDay
			}
		}

		if input.Query.UserIdIn != nil && len(*input.Query.UserIdIn) > 0 {
			query.UserIdIn = input.Query.UserIdIn
		}
	}

	return query, paginationData
}
