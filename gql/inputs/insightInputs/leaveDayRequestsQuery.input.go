package insightInputs

import (
	"time"

	"github.com/graph-gophers/graphql-go"
)

type LeaveDayRequestsQueryInput struct {
	RequestTypeEq  *string
	RequestStateEq *string
	UserIdEq       *int32
	FromGteq       *time.Time
	ToLteq         *time.Time
}

type LeaveDayRequestsFrontQueryInput struct {
	RequestTypeEq  *string
	RequestStateEq *string
	UserIdEq       *graphql.ID
	FromGteq       *string
	ToLteq         *string
}
