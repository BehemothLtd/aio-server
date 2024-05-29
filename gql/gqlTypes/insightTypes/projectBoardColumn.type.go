package insightTypes

import (
	"aio-server/database"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"fmt"

	"github.com/graph-gophers/graphql-go"
)

type ProjectBoardType struct {
	Project models.Project

	IssueRepo repository.IssueRepository
	Issues    []*models.Issue
}

func (board ProjectBoardType) Columns() []ProjectBoardColumnType {
	result := []ProjectBoardColumnType{}

	fmt.Printf("BOARD ISSUE STATUSES : %+v", board.Project.ProjectIssueStatuses)
	fmt.Printf("Project: %+v", board.Project)

	board.IssueRepo = *repository.NewIssueRepository(nil, database.Db)
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
