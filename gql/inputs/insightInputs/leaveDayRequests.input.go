package insightInputs

import (
	"aio-server/gql/inputs/globalInputs"
	"aio-server/models"
	"strings"
)

type LeaveDayRequestsInput struct {
	Input *globalInputs.PagyInput
	Query *LeaveDayRequestsQueryInput
}

func (ldi *LeaveDayRequestsInput) ToPaginationDataAndQuery() (LeaveDayRequestsQueryInput, models.PaginationData) {
	paginationData := ldi.Input.ToPaginationInput()
	query := LeaveDayRequestsQueryInput{}

	if ldi.Query != nil {
		if ldi.Query.RequestState != nil && strings.TrimSpace(*ldi.Query.RequestState) != "" {
			query.RequestState = ldi.Query.RequestState
		}
		if ldi.Query.RequestType != nil && strings.TrimSpace(*ldi.Query.RequestType) != "" {
			query.RequestType = ldi.Query.RequestType
		}
		if ldi.Query.UserId != nil {
			query.RequestType = ldi.Query.RequestType
		}
	}

	return query, paginationData
}
