package insightInputs

import (
	"aio-server/enums"
	"aio-server/gql/inputs/globalInputs"
	"aio-server/models"
	"strings"
)

type IssuesInput struct {
	Input *globalInputs.PagyInput
	Query *IssuesQueryInput
}

func (is *IssuesInput) ToPaginationDataAndQuery() (IssuesQueryInput, models.PaginationData) {
	paginationData := is.Input.ToPaginationInput()
	query := IssuesQueryInput{}

	if is.Query != nil && is.Query.TitleCont != nil && strings.TrimSpace(*is.Query.TitleCont) != "" {
		query.TitleCont = is.Query.TitleCont
	}

	if is.Query != nil && is.Query.CodeCont != nil && strings.TrimSpace(*is.Query.CodeCont) != "" {
		query.CodeCont = is.Query.CodeCont
	}

	if is.Query != nil && is.Query.IssueTypeEq != nil {
		query.IssueTypeEq = is.Query.IssueTypeEq

		_, err := enums.ParseIssueType(*is.Query.IssueTypeEq)

		if err != nil {
			query.IssueTypeEq = nil
		}
	}

	return query, paginationData
}
