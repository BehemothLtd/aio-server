package msInputs

import (
	"aio-server/gql/inputs/globalInputs"
	"aio-server/models"
	"strings"
)

type CollectionsInput struct {
	Input *globalInputs.PagyInput
	Query *CollectionQueryInput
}

func (msi *CollectionsInput) ToPaginationDataAndCollectionQuery() (CollectionQueryInput, models.PaginationData) {
	paginationData := msi.Input.ToPaginationInput()
	query := CollectionQueryInput{}

	if msi.Query != nil && msi.Query.TitleCont != nil && strings.TrimSpace(*msi.Query.TitleCont) != "" {
		query.TitleCont = msi.Query.TitleCont
	}

	return query, paginationData
}
