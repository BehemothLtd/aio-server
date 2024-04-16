package msInputs

import (
	"aio-server/gql/inputs/globalInputs"
	"aio-server/models"
	"strings"
)

// SnippetsInput represents input for querying snippets collection.
type SnippetsInput struct {
	Input *globalInputs.PagyInput
	Query *SnippetsQueryInput
}

// ToPaginationDataAndQuery converts SnippetsInput to models.SnippetsQuery and models.PaginationData.
func (msi *SnippetsInput) ToPaginationDataAndQuery() (SnippetsQueryInput, models.PaginationData) {
	paginationData := msi.Input.ToPaginationInput()
	query := SnippetsQueryInput{}

	if msi.Query != nil && msi.Query.TitleCont != nil && strings.TrimSpace(*msi.Query.TitleCont) != "" {
		query.TitleCont = msi.Query.TitleCont
	}

	if msi.Query != nil && msi.Query.SnippetTypeEq != nil && strings.TrimSpace(*msi.Query.SnippetTypeEq) != "" {
		query.SnippetTypeEq = msi.Query.SnippetTypeEq
	}

	return query, paginationData
}
