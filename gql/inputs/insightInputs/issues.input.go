package insightInputs

import (
	"aio-server/gql/inputs/globalInputs"
	"aio-server/models"
	"aio-server/pkg/constants"
	"aio-server/pkg/helpers"
	"strings"
	"time"
)

type IssuesFrontQueryInput struct {
	TitleCont      *string
	CodeCont       *string
	IssueTypeEq    *string
	ProjectIdEq    *int32
	DeadLineAtGteq *string
	DeadLineAtLteq *string
}

type IssuesQueryInput struct {
	TitleCont      *string
	CodeCont       *string
	IssueTypeEq    *string
	ProjectIdEq    *int32
	DeadLineAtGteq *time.Time
	DeadLineAtLteq *time.Time
}

type IssuesInput struct {
	Input *globalInputs.PagyInput
	Query *IssuesFrontQueryInput
}

func (input *IssuesInput) ToPaginationDataAndQuery() (IssuesQueryInput, models.PaginationData) {
	paginationData := input.Input.ToPaginationInput()

	query := IssuesQueryInput{}

	if input.Query != nil {
		if input.Query.TitleCont != nil && strings.TrimSpace(*input.Query.TitleCont) != "" {
			query.TitleCont = input.Query.TitleCont
		}

		if input.Query.CodeCont != nil && strings.TrimSpace(*input.Query.CodeCont) != "" {
			query.CodeCont = input.Query.CodeCont
		}

		if input.Query.ProjectIdEq != nil {
			query.ProjectIdEq = input.Query.ProjectIdEq
		}

		if input.Query.IssueTypeEq != nil && strings.TrimSpace(*input.Query.IssueTypeEq) != "" {
			query.IssueTypeEq = input.Query.IssueTypeEq
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

	}

	return query, paginationData
}
