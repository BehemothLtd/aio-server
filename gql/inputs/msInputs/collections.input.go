package msInputs

import (
	"aio-server/gql/inputs/globalInputs"
	"aio-server/models"
)

type CollectionsInput struct {
	Input *globalInputs.PagyInput
	Query *CollectionQueryInput
}

func (msi *CollectionsInput) ToPaginationDataAndCollectionQuery() (models.CollectionQuery, models.PaginationData) {
	paginationData := msi.Input.ToPaginationInput()

	var titleCont string
	if msi.Query != nil && msi.Query.TitleCont != nil {
		titleCont = *msi.Query.TitleCont
	}

	return models.CollectionQuery{TitleCont: titleCont}, paginationData
}
