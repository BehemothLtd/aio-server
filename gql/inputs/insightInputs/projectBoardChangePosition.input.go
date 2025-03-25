package insightInputs

type ProjectBoardChangePositionInput struct {
	Id          int32
	ProjectId   int32
	NewIndex    int32
	NewStatusId int32
}
