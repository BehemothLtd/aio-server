package insightinputs

import (
	"aio-server/gql/inputs/globalInputs"
)

type LeaveDayRequestsInput struct {
	Input *globalInputs.PagyInput
	Query *LeaveDayRequestsQueryInput
}
