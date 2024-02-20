package msInputs

import (
	"aio-server/gql/inputs/globalInputs"
	"aio-server/models"
)

// MsSnippetsInput represents input for querying snippets collection.
type SnippetsInput struct {
	Input *globalInputs.PagyInput
	Query *SnippetQueryInput
}

// ToPaginationDataAndSnippetsQuery converts MsSnippetsInput to models.SnippetsQuery and models.PaginationData.
func (msi *SnippetsInput) ToPaginationDataAndSnippetsQuery() (models.SnippetsQuery, models.PaginationData) {
	paginationData := msi.Input.ToPaginationInput()

	var titleCont string
	if msi.Query != nil && msi.Query.TitleCont != nil {
		titleCont = *msi.Query.TitleCont
	}

	return models.SnippetsQuery{TitleCont: titleCont}, paginationData
}
