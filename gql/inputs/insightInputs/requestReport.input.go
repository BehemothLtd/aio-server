package insightInputs

import (
	"aio-server/enums"
	"strings"
)

type RequestReportInput struct {
	Query *RequestReportQueryInput
}

type RequestReportQueryInput struct {
	UserIdEq         *int32
	RequestTypeIn    *[]*string
	CreatedAtBetween *[]*string
}

func (rpi *RequestReportInput) ToQuery() RequestReportQueryInput {
	query := RequestReportQueryInput{}

	if rpi.Query != nil {

		// Handle user_id_eq params
		if rpi.Query.UserIdEq != nil {
			query.UserIdEq = rpi.Query.UserIdEq
		}
		// Handle created_at_between params
		if rpi.Query.CreatedAtBetween != nil {
			query.CreatedAtBetween = rpi.Query.CreatedAtBetween
		}

		// Handle request_type_in params
		requestTypes := rpi.Query.RequestTypeIn

		if requestTypes != nil && len(*requestTypes) > 0 {
			listRequestTypes := []*string{}

			for _, requestType := range *requestTypes {
				if requestType != nil && strings.TrimSpace(*requestType) != "" {
					_, err := enums.ParseRequestType(*requestType)

					if err != nil {
						continue
					}
					listRequestTypes = append(listRequestTypes, requestType)
				}
			}
			query.RequestTypeIn = &listRequestTypes
		}
	}

	return query
}
