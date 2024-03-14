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
		if ldi.Query.RequestStateCont != nil && strings.TrimSpace(*ldi.Query.RequestStateCont) != "" {
			query.RequestStateCont = ldi.Query.RequestStateCont
		}
		if ldi.Query.RequestTypeCont != nil && strings.TrimSpace(*ldi.Query.RequestTypeCont) != "" {
			query.RequestTypeCont = ldi.Query.RequestTypeCont
		}
		if ldi.Query.UserIdEq != nil {
			query.UserIdEq = ldi.Query.UserIdEq
		}
	}

	return query, paginationData
}
