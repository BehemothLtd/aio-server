package insightInputs

import (
	"aio-server/gql/inputs/globalInputs"
	"aio-server/models"
	"strings"
)

type IssueStatusesInput struct {
	Input *globalInputs.PagyInput
	Query *IssueStatusesQueryInput
}

func (isi *IssueStatusesInput) ToPaginationDataAndQuery() (IssueStatusesQueryInput, models.PaginationData) {
	paginationData := isi.Input.ToPaginationInput()
	query := IssueStatusesQueryInput{}

	if isi.Query != nil && isi.Query.TitleCont != nil && strings.TrimSpace(*isi.Query.TitleCont) != "" {
		query.TitleCont = isi.Query.TitleCont
	}

	if isi.Query != nil && isi.Query.StatusTypeEq != nil {
		query.StatusTypeEq = isi.Query.StatusTypeEq
	}
	return query, paginationData
}
