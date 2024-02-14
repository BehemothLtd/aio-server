package inputs

import "aio-server/models"

// For create/update
type MsSnippetInput struct {
	Title       *string
	Content     *string
	SnippetType *int32 // we have to use int32 OR float32 for input struct because Graphql ASKED FOR IT
}

func (msi *MsSnippetInput) ToFormInput() struct {
	Title       string
	Content     string
	SnippetType int32
} {
	result := struct {
		Title       string
		Content     string
		SnippetType int32
	}{}

	if msi != nil {
		if msi.Title == nil {
			result.Title = ""
		} else {
			result.Title = *msi.Title
		}

		if msi.Content == nil {
			result.Content = ""
		} else {
			result.Content = *msi.Content
		}

		if msi.SnippetType == nil {
			result.SnippetType = 0
		} else {
			result.SnippetType = *msi.SnippetType
		}
	}

	return result
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
