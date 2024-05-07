package insightInputs

import (
	"aio-server/enums"
	"strings"
)

type RequestReportInput struct {
	UserIdEq         *int32
	RequestTypeIn    *[]*string
	CreatedAtBetween *[]*string
}

func (rpi *RequestReportInput) ToQuery() RequestReportInput {
	query := RequestReportInput{}

	// Handle user_id_eq params
	if rpi.UserIdEq != nil {
		query.UserIdEq = rpi.UserIdEq
	}
	// Handle created_at_between params
	if rpi.CreatedAtBetween != nil {
		query.CreatedAtBetween = rpi.CreatedAtBetween
	}

	// Handle request_type_in params
	requestTypes := rpi.RequestTypeIn

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

	return query
}
