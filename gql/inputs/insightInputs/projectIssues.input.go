package insightInputs

import (
	"aio-server/gql/inputs/globalInputs"
	"aio-server/models"
	"strings"

	"github.com/graph-gophers/graphql-go"
)

type ProjectIssuesQueryInput struct {
	TitleCont   *string
	CodeCont    *string
	IssueTypeEq *string
	ProjectIdEq *int32
}

type ProjectIssuesInput struct {
	Id    graphql.ID
	Input *globalInputs.PagyInput
	Query *ProjectIssuesQueryInput
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
	}

	return query, paginationData
}
