package insightInputs

import (
	"aio-server/enums"
	"aio-server/gql/inputs/globalInputs"
	"aio-server/models"
	"strings"
)

type DevicesInput struct {
	Input *globalInputs.PagyInput
	Query *DevicesQueryInput
}

func (di *DevicesInput) ToPaginationDataAndQuery() (DevicesQueryInput, models.PaginationData) {
	paginationData := di.Input.ToPaginationInput()
	query := DevicesQueryInput{}

	if di.Query != nil {
		if di.Query.DeviceTypeIdIn != nil && len(*di.Query.DeviceTypeIdIn) > 0 {
			query.DeviceTypeIdIn = di.Query.DeviceTypeIdIn
		}

		if di.Query.NameCont != nil && strings.TrimSpace(*di.Query.NameCont) != "" {
			query.NameCont = di.Query.NameCont
		}

		if di.Query.StateIn != nil && len(*di.Query.StateIn) > 0 {
			query.StateIn = di.Query.StateIn

			for _, state := range *di.Query.StateIn {
				_, err := enums.ParseDeviceStateType(state)
				if err != nil {
					query.StateIn = nil
				}
			}
		}

		if di.Query.UserIdIn != nil && len(*di.Query.UserIdIn) > 0 {
			query.UserIdIn = di.Query.UserIdIn
		}
	}

	return query, paginationData
}
