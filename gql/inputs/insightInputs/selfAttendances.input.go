package insightInputs

import (
	"aio-server/gql/inputs/globalInputs"
	"aio-server/models"
	"aio-server/pkg/constants"
	"strings"
	"time"
)

type SelfAttendancesInput struct {
	Input *globalInputs.PagyInput
	Query *SelfAttendancesFrontQueryInput
}

type SelfAttendancesFrontQueryInput struct {
	CheckinAtGteq *string
	CheckinAtLteq *string
}

type SelfAttendancesQueryInput struct {
	CheckinAtGteq *time.Time
	CheckinAtLteq *time.Time
}

func (sai *SelfAttendancesInput) ToPaginationAndQueryData() (SelfAttendancesQueryInput, models.PaginationData) {
	paginationData := sai.Input.ToPaginationInput()
	query := SelfAttendancesQueryInput{}

	if sai.Query != nil && sai.Query.CheckinAtGteq != nil && strings.TrimSpace(*sai.Query.CheckinAtGteq) != "" {
		if timeValue, err := time.ParseInLocation(constants.YYYYMMDD_DateFormat, *sai.Query.CheckinAtGteq, time.Local); err != nil {
			query.CheckinAtGteq = nil
		} else {
			year, month, day := timeValue.Date()
			value := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
			query.CheckinAtGteq = &value
		}
	}

	if sai.Query != nil && sai.Query.CheckinAtLteq != nil && strings.TrimSpace(*sai.Query.CheckinAtLteq) != "" {
		if timeValue, err := time.ParseInLocation(constants.YYYYMMDD_DateFormat, *sai.Query.CheckinAtLteq, time.Local); err != nil {
			query.CheckinAtLteq = nil
		} else {
			year, month, day := timeValue.Date()
			value := time.Date(year, month, day, 23, 59, 59, 59, time.Local)

			query.CheckinAtLteq = &value
		}
	}

	return query, paginationData
}
