package insightTypes

type SelectOptionsType struct {
	IssueStatusOptions     []IssueStatusSelectOption
	DevelopmentRoleOptions []CommonSelectOption
	UserOptions            []CommonSelectOption
	ProjectOptions         []CommonSelectOption
	ClientOptions          []CommonSelectOption
	DeviceTypeOptions      []DeviceTypeSelectOption
	IssueTypeOptions       []StringStringSelectOption
	IssuePriorityOptions   []StringStringSelectOption
}

type IssueStatusSelectOption struct {
	CommonSelectOption
	Color string
}

type CommonSelectOption struct {
	Label string
	Value int32
}

type StringStringSelectOption struct {
	Label string
	Value string
}

type DeviceTypeSelectOption struct {
	Label string
	Value int32
}
