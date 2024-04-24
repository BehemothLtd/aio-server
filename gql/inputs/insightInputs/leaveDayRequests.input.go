package insightInputs

import (
	"aio-server/enums"
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

			_, err := enums.ParseRequestStateType(*ldi.Query.RequestStateEq)

			if err != nil {
				query.RequestStateEq = nil
			}
		}

		if ldi.Query.RequestTypeEq != nil && strings.TrimSpace(*ldi.Query.RequestTypeEq) != "" {
			query.RequestTypeEq = ldi.Query.RequestTypeEq

			_, err := enums.ParseRequestType(*ldi.Query.RequestTypeEq)
			if err != nil {
				query.RequestTypeEq = nil
			}
		}

		if ldi.Query.UserIdEq != nil {
			query.UserIdEq = ldi.Query.UserIdEq
		}

		if ldi.Query.FromGteq != nil {
			query.FromGteq = ldi.Query.FromGteq
		}

		if ldi.Query.ToLteq != nil {
			query.ToLteq = ldi.Query.ToLteq
		}
	}

	return query, paginationData
}
