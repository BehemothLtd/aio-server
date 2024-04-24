package insightInputs

type LeaveDayRequestsQueryInput struct {
	RequestTypeEq  *string
	RequestStateEq *string
	UserIdEq       *int32
	FromGteq       *string
	ToLteq         *string
}
