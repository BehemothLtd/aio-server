package insightTypes

type SelectOptionsType struct {
	IssueStatusOptions     []IssueStatusSelectOption
	DevelopmentRoleOptions []CommonSelectOption
	UserOptions            []CommonSelectOption
}

type IssueStatusSelectOption struct {
	CommonSelectOption
	Color string
}

type CommonSelectOption struct {
	Label string
	Value int32
}
