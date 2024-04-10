package insightInputs

type DevicesQueryInput struct {
	DeviceTypeIdIn *[]*int32
	NameCont       *string
	StateIn        *[]*string
	UserIdIn       *[]*int32
}
