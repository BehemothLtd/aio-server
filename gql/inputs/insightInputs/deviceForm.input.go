package insightInputs

type DeviceFormInput struct {
	Name         string
	Code         string
	State        string
	UserId       *int32
	DeviceTypeId int32
	Description  *string
	Seller       *string
}
