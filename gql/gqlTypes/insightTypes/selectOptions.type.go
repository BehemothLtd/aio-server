package insightTypes

type SelectOptionsType struct {
	IssueStatusOptions     []IssueStatusSelectOption
	DevelopmentRoleOptions []CommonSelectOption
}

type IssueStatusSelectOption struct {
	CommonSelectOption
	Color string
}

type CommonSelectOption struct {
	Label string
	Value string
}
