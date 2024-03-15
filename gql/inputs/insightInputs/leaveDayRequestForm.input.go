package insightInputs

import "time"

type LeaveDayRequestFormInput struct {
	From         *time.Time
	To           *time.Time
	TimeOff      *float64
	RequestType  *string
	RequestState *string
	Reason       *string
	LockVersion  *int32
}
