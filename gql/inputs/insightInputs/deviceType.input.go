package insightInputs

import "github.com/graph-gophers/graphql-go"

type DeviceTypeInput struct {
	Id graphql.ID
}

type DeviceTypeCreateInput struct {
	Input DeviceTypeFormInput
}

type DeviceTypeUpdateInput struct {
	Id    graphql.ID
	Input DeviceTypeFormInput
}

type DeviceTypeFormInput struct {
	Name *string
}
