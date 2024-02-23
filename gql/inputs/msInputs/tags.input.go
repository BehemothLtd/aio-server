package msInputs

import (
	"aio-server/gql/inputs/globalInputs"
	"aio-server/models"
)

// TagsInput represents input for querying tags collection.
type TagsInput struct {
	Input *globalInputs.PagyInput
	Query *TagQueryInput
}

// ToPaginationDataAndSnippetsQuery converts SnippetsInput to models.SnippetsQuery and models.PaginationData.
func (msi *TagsInput) ToPaginationDataAndSnippetsQuery() (models.TagsQuery, models.PaginationData) {
	paginationData := msi.Input.ToPaginationInput()

	var nameCont string
	if msi.Query != nil && msi.Query.NameCont != nil {
		nameCont = *msi.Query.NameCont
	}

	return models.TagsQuery{NameCont: nameCont}, paginationData
}
