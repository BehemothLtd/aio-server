package insightInputs

import "github.com/graph-gophers/graphql-go"

type DeviceFormInput struct {
	Name         string
	Code         string
	State        string
	UserId       *graphql.ID
	DeviceTypeId graphql.ID
	Description  *string
	Seller       *string
}
