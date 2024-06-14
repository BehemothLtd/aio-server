package insightInputs

import (
	"aio-server/gql/inputs/globalInputs"
	"aio-server/models"
)

type DevicesUsingHistoriesInput struct {
	Input *globalInputs.PagyInput
	Query *DevicesUsingHistoriesQueryInput
}

func (duhi *DevicesUsingHistoriesInput) ToPaginationDataAndQuery() (DevicesUsingHistoriesQueryInput, models.PaginationData) {
	paginationData := duhi.Input.ToPaginationInput()
	query := DevicesUsingHistoriesQueryInput{}

	if duhi.Query != nil {
		if duhi.Query.DeviceIdIn != nil && len(*duhi.Query.DeviceIdIn) > 0 {
			query.DeviceIdIn = duhi.Query.DeviceIdIn
		}
	}

	return query, paginationData
}
