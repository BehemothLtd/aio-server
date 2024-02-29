package msInputs

import (
	"aio-server/gql/inputs/globalInputs"
	"aio-server/models"
)

// SnippetsInput represents input for querying snippets collection.
type SnippetsInput struct {
	Input *globalInputs.PagyInput
	Query *SnippetQueryInput
}

// ToPaginationDataAndQuery converts SnippetsInput to models.SnippetsQuery and models.PaginationData.
func (msi *SnippetsInput) ToPaginationDataAndQuery() (SnippetQueryInput, models.PaginationData) {
	paginationData := msi.Input.ToPaginationInput()
	query := SnippetQueryInput{}

	if msi.Query != nil && msi.Query.TitleCont != nil {
		query.TitleCont = msi.Query.TitleCont
	}

	return query, paginationData
}
