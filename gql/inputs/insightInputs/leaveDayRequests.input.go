package insightinputs

import (
	"aio-server/gql/inputs/globalInputs"
	"aio-server/models"
)

type LeaveDayRequestsInput struct {
	Input *globalInputs.PagyInput
	Query *LeaveDayRequestsQueryInput
}

func (ldi *LeaveDayRequestsInput) ToPaginationDataAndQuery() (LeaveDayRequestsInput, models.PaginationData) {
	paginationData := ldi.Input.ToPaginationInput()
}
