package insightInputs

import "github.com/graph-gophers/graphql-go"

type AttendanceInput struct {
	Id graphql.ID
}

type AttendanceFormInput struct {
	CheckinAt  string
	CheckoutAt string
	UserId     int32
}

type AttendanceCreateInput struct {
	Input AttendanceFormInput
}

type AttendanceUpdateInput struct {
	Id    graphql.ID
	Input AttendanceFormInput
}
