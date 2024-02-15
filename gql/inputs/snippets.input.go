package inputs

import "aio-server/models"

// MsSnippetsInput represents input for querying snippets collection.
type MsSnippetsInput struct {
	Input *PagyInput
	Query *SnippetQueryInput
}

// ToPaginationDataAndSnippetsQuery converts MsSnippetsInput to models.SnippetsQuery and models.PaginationData.
func (msi *MsSnippetsInput) ToPaginationDataAndSnippetsQuery() (models.SnippetsQuery, models.PaginationData) {
	paginationData := msi.Input.ToPaginationInput()

	var titleCont string
	if msi.Query != nil && msi.Query.TitleCont != nil {
		titleCont = *msi.Query.TitleCont
	}

	return models.SnippetsQuery{TitleCont: titleCont}, paginationData
}
