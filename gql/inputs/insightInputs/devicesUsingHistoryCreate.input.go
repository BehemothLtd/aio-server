package insightInputs

type DevicesUsingHistoryCreateInput struct {
	Input DevicesUsingHistoryCreateFormInput
}

type DevicesUsingHistoryCreateFormInput struct {
	UserId   *int32
	DeviceId *int32
	State    *string
}
