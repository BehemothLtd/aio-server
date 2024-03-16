package insightInputs

import (
    "aio-server/gql/inputs/globalInputs"
    "aio-server/models"
    "strings"
)

// ClientInput represents input for querying clients collection.
type ClientsInput struct {
    Input *globalInputs.PagyInput
    Query *ClientsQueryInput
}

// ToPaginationDataAndClientQuery converts ClientInput to appropriate types for pagination.
func (ci *ClientsInput) ToPaginationDataAndQuery() (ClientsQueryInput, models.PaginationData) {
    paginationData := ci.Input.ToPaginationInput()

    // Create an instance of the anonymous empty struct and take its address
    query := ClientsQueryInput{}

    if ci.Query != nil && ci.Query.NameCont != nil && strings.TrimSpace(*ci.Query.NameCont) != "" {
		query.NameCont = ci.Query.NameCont
	}

    return query, paginationData
}
