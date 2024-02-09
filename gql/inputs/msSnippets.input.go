package inputs

import "aio-server/models"

type MsSnippetsInput struct {
	Input *PagyInput
	Query *SnippetQueryInput
}

func (msi *MsSnippetsInput) ToPaginationDataAndSnippetsQuery() (models.SnippetsQuery, models.PaginationData) {
	paginationData := msi.Input.ToPaginationInput()

	outputQuery := models.SnippetsQuery{TitleCont: ""}

	if msi.Query != nil && msi.Query.TitleCont != nil {
		outputQuery.TitleCont = *msi.Query.TitleCont
	}

	return outputQuery, paginationData
}
