package inputs

import "aio-server/models"

// For create/update
type MsSnippetInput struct {
	Title       *string
	Content     *string
	SnippetType *int32 // we have to use int32 OR float32 for input struct
}

// For List
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
