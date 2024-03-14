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
		if ldi.Query.RequestStateEq != nil && strings.TrimSpace(*ldi.Query.RequestStateEq) != "" {
			query.RequestStateEq = ldi.Query.RequestStateEq
		}
		if ldi.Query.RequestTypeEq != nil && strings.TrimSpace(*ldi.Query.RequestTypeEq) != "" {
			query.RequestTypeEq = ldi.Query.RequestTypeEq
		}
		if ldi.Query.UserIdEq != nil {
			query.UserIdEq = ldi.Query.UserIdEq
		}
	}

	return query, paginationData
}
