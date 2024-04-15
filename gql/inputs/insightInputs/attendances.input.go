package insightInputs

import (
	"aio-server/gql/inputs/globalInputs"
	"aio-server/models"
	"aio-server/pkg/constants"
	"aio-server/pkg/helpers"
	"strings"
	"time"

	"github.com/graph-gophers/graphql-go"
)

type AttendancesInput struct {
	Input *globalInputs.PagyInput
	Query *AttendancesFrontQueryInput
}

type AttendancesFrontQueryInput struct {
	CheckinAtGteq *string
	CheckinAtLteq *string
	UserIdEq      *graphql.ID
}

type AttendancesQueryInput struct {
	CheckinAtGteq *time.Time
	CheckinAtLteq *time.Time
	UserIdEq      *int32
}

func (ai *AttendancesInput) ToPaginationAndQueryData() (AttendancesQueryInput, models.PaginationData) {
	paginationData := ai.Input.ToPaginationInput()
	query := AttendancesQueryInput{}

	if ai.Query != nil {
		if ai.Query.CheckinAtGteq != nil && strings.TrimSpace(*ai.Query.CheckinAtGteq) != "" {
			if timeValue, err := time.ParseInLocation(constants.YYYYMMDD_DateFormat, *ai.Query.CheckinAtGteq, time.Local); err != nil {
				query.CheckinAtGteq = nil
			} else {
				beginOfDay := helpers.BeginningOfDay(&timeValue)
				query.CheckinAtGteq = &beginOfDay
			}
		}

		if ai.Query.CheckinAtLteq != nil && strings.TrimSpace(*ai.Query.CheckinAtLteq) != "" {
			if timeValue, err := time.ParseInLocation(constants.YYYYMMDD_DateFormat, *ai.Query.CheckinAtLteq, time.Local); err != nil {
				query.CheckinAtLteq = nil
			} else {
				endOfDay := helpers.EndOfDay(&timeValue)
				query.CheckinAtLteq = &endOfDay
			}
		}

		if ai.Query.UserIdEq != nil {
			if UserIdEq, err := helpers.GqlIdToInt32(*ai.Query.UserIdEq); err != nil {
				query.UserIdEq = nil
			} else {
				query.UserIdEq = &UserIdEq
			}
		}
	}

	return query, paginationData
}
