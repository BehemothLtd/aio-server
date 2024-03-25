package insightInputs

import "github.com/graph-gophers/graphql-go"

type LeaveDayRequestUpdateInput struct {
	Id    graphql.ID
	Input LeaveDayRequestFormInput
}
