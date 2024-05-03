package insightInputs

type DeviceCreateInput struct {
	Input DeviceCreateFormInput
}

type DeviceCreateFormInput struct {
	Name         string
	Code         string
	State        string
	UserId       *int32
	DeviceTypeId int32
	Description  *string
	Seller       *string
}
