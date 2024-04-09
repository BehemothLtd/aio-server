package insightTypes

type SelfThisWeekIssuesStateType struct {
	Labels []string
	Series SelfThisWeekIssuesStateSeriesType
}

type SelfThisWeekIssuesStateSeriesType struct {
	Done    []int32
	NotDone []int32
}
