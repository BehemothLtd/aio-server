package insightTypes

type SelectOptionsType struct {
	IssueStatusOptions []IssueStatusSelectOption
}

type IssueStatusSelectOption struct {
	CommonSelectOption
	Color string
}

type CommonSelectOption struct {
	Label string
	Value string
}
