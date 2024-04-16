package insightTypes

type SelectOptionsType struct {
	IssueStatusOptions []IssueStatusSelectOption
}

type IssueStatusSelectOption struct {
	Label string
	Value string
	Color string
}
