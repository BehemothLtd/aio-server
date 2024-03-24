package insightInputs

type IssueStatusCreateInput struct {
	Input IssueStatusCreateFormInput
}

type IssueStatusCreateFormInput struct {
	Color      *string
	Title      *string
	StatusType *string
}
