package insightInputs

import (
	"aio-server/enums"
	"aio-server/gql/inputs/globalInputs"
	"aio-server/models"
	"aio-server/pkg/constants"
	"aio-server/pkg/helpers"
	"strings"
	"time"
)

type LeaveDayRequestsInput struct {
	Input *globalInputs.PagyInput
	Query *LeaveDayRequestsFrontQueryInput
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
			if UserIdEq, err := helpers.GqlIdToInt32(*ldi.Query.UserIdEq); err != nil {
				query.UserIdEq = nil
			} else {
				query.UserIdEq = &UserIdEq
			}
		}

		if ldi.Query.FromGteq != nil && strings.TrimSpace(*ldi.Query.FromGteq) != "" {
			if timeValue, err := time.ParseInLocation(constants.DateTimeFormat, *ldi.Query.FromGteq, time.Local); err != nil {
				print(err.Error())
				query.FromGteq = nil
			} else {
				query.FromGteq = &timeValue
			}
		}

		if ldi.Query.ToLteq != nil && strings.TrimSpace(*ldi.Query.ToLteq) != "" {
			if timeValue, err := time.ParseInLocation(constants.DateTimeFormat, *ldi.Query.ToLteq, time.Local); err != nil {
				print(err.Error())

				query.ToLteq = nil
			} else {
				query.ToLteq = &timeValue
			}
		}
	}

	return query, paginationData
}
