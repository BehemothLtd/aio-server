package insightInputs

import "github.com/graph-gophers/graphql-go"

type LeaveDayRequestStateChangeInput struct {
	Id           graphql.ID
	RequestState string
}
