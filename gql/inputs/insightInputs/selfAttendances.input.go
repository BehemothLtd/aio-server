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
		if timeValue, err := time.Parse(constants.YYYYMMDD_DateFormat, *sai.Query.CheckinAtGteq); err != nil {
			query.CheckinAtGteq = nil
		} else {
			query.CheckinAtGteq = &timeValue
		}
	}

	if sai.Query != nil && sai.Query.CheckinAtLteq != nil && strings.TrimSpace(*sai.Query.CheckinAtLteq) != "" {
		if timeValue, err := time.Parse(constants.YYYYMMDD_DateFormat, *sai.Query.CheckinAtLteq); err != nil {
			query.CheckinAtLteq = nil
		} else {
			query.CheckinAtLteq = &timeValue
		}
	}

	return query, paginationData
}
