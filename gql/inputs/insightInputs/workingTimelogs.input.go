package insightInputs

import (
	"aio-server/gql/inputs/globalInputs"
	"aio-server/models"
	"strings"
)

type WorkingTimelogsInput struct {
	Input *globalInputs.PagyInput
	Query *WorkingTimelogsQueryInput
}

func (wti *WorkingTimelogsInput) ToPaginationDataAndQuery() (WorkingTimelogsQueryInput, models.PaginationData) {
	paginationData := wti.Input.ToPaginationInput()
	query := WorkingTimelogsQueryInput{}

	if wti.Query != nil && wti.Query.DescriptionCont != nil && strings.TrimSpace(*wti.Query.DescriptionCont) != "" {
		query.DescriptionCont = wti.Query.DescriptionCont
	}

	if wti.Query != nil && wti.Query.IssueCodeEq != nil {
		query.IssueCodeEq = wti.Query.IssueCodeEq
	}

	if wti.Query != nil && wti.Query.IssueTitleCont != nil && strings.TrimSpace(*wti.Query.IssueTitleCont) != "" {
		query.IssueTitleCont = wti.Query.IssueTitleCont
	}

	return query, paginationData
}
