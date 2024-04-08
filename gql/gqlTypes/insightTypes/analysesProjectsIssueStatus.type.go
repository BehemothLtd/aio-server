package insightTypes

type AnalysesProjectsIssueStatus struct {
	Categories []string
	Series     []AnalysesProjectsIssueStatusCounting
}

type AnalysesProjectsIssueStatusCounting struct {
	Data []int32
	Name string
}
