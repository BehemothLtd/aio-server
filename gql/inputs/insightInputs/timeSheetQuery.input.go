package insightInputs

type TimeSheetQuery struct {
	LoggedAtBetween   *string
	ProjectIdIn       *string
	ProjectClientIdEq *string
}
