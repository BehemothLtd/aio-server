package insightTypes

import (
	"aio-server/database"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/repository"

	"github.com/graph-gophers/graphql-go"
)

type ProjectBoardType struct {
	Project models.Project

	IssueRepo repository.IssueRepository
	Issues    []*models.Issue
}

func (board ProjectBoardType) Columns() []ProjectBoardColumnType {
	result := []ProjectBoardColumnType{}

	board.IssueRepo = *repository.NewIssueRepository(
		nil,
		database.Db.Preload("IssueAssignees.User.Avatar.AttachmentBlob").
			Preload("Creator.Avatar.AttachmentBlob").
			Preload("IssueStatus"),
	)
	board.IssueRepo.FetchProjectBoardIssues(board.Project, &board.Issues)

	for _, projectIssueStatus := range board.Project.ProjectIssueStatuses {
		result = append(result, board.ContructColumn(projectIssueStatus))
	}

	return result
}

func (board ProjectBoardType) ContructColumn(projectIssueStatus *models.ProjectIssueStatus) ProjectBoardColumnType {
	columnIssues := []*globalTypes.IssueType{}

	for _, issue := range board.Issues {
		if issue.IssueStatusId == projectIssueStatus.IssueStatusId {
			columnIssues = append(columnIssues, &globalTypes.IssueType{
				Issue: issue,
			})
		}
	}

	return ProjectBoardColumnType{
		Id: helpers.GqlIDP(projectIssueStatus.Id),
		IssueStatus: &globalTypes.IssueStatusType{
			IssueStatus: &projectIssueStatus.IssueStatus,
		},
		Issues: &columnIssues,
	}
}

type ProjectBoardColumnType struct {
	Id          *graphql.ID
	IssueStatus *globalTypes.IssueStatusType
	Issues      *[]*globalTypes.IssueType
}
