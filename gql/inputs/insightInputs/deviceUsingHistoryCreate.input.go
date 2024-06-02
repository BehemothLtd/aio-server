package insightInputs

type DeviceUsingHistoryCreateInput struct {
	Input DeviceUsingHistoryCreateFormInput
}

type DeviceUsingHistoryCreateFormInput struct {
	UserId   *int32
	DeviceId *int32
	State    *string
}
