package msInputs

import (
	"aio-server/enums"
	"aio-server/gql/inputs/globalInputs"
	"aio-server/models"
	"strings"
)

type SelfSnippetsInput struct {
	Input *globalInputs.PagyInput
	Query *SelfSnippetsQueryInput
}

// ToPaginationDataAndQuery converts SelfSnippetsInput to models.SnippetsQuery and models.PaginationData.
func (msi *SelfSnippetsInput) ToPaginationDataAndQuery() (SelfSnippetsQueryInput, models.PaginationData) {
	paginationData := msi.Input.ToPaginationInput()
	query := SelfSnippetsQueryInput{}

	if msi.Query != nil && msi.Query.TitleCont != nil && strings.TrimSpace(*msi.Query.TitleCont) != "" {
		query.TitleCont = msi.Query.TitleCont
	}

	if msi.Query != nil && msi.Query.SnippetType != nil && strings.TrimSpace(*msi.Query.SnippetType) != "" {
		_, err := enums.ParseSnippetType(*msi.Query.SnippetType)

		if err == nil {
			query.SnippetType = msi.Query.SnippetType
		}
	}

	return query, paginationData
}
