package insightInputs

import "github.com/graph-gophers/graphql-go"

type DeviceUpdateInput struct {
	Id    graphql.ID
	Input DeviceCreateFormInput
}
