package msInputs

import (
	"aio-server/gql/inputs/globalInputs"
	"aio-server/models"
	"strings"
)

// TagsInput represents input for querying tags collection.
type TagsInput struct {
	Input *globalInputs.PagyInput
	Query *TagsQueryInput
}

// ToPaginationDataAndQuery converts TagsInput to TagsQueryInput and models.PaginationData.
func (msi *TagsInput) ToPaginationDataAndQuery() (TagsQueryInput, models.PaginationData) {
	paginationData := msi.Input.ToPaginationInput()
	query := TagsQueryInput{}

	if msi.Query != nil && msi.Query.NameCont != nil && strings.TrimSpace(*msi.Query.NameCont) != "" {
		query.NameCont = msi.Query.NameCont
	}

	return query, paginationData
}
