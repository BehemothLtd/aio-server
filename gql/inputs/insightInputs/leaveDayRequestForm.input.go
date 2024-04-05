package insightInputs

type LeaveDayRequestFormInput struct {
	From        string
	To          string
	TimeOff     float64
	RequestType string
	Mentions    *[]*string
	Reason      *string
	LockVersion *int32
}
