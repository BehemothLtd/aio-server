package insightServices

import (
	"aio-server/gql/gqlTypes/insightTypes"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"fmt"
	"slices"

	"golang.org/x/exp/maps"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type AnalysesProjectsIssueStatusService struct {
	Ctx  context.Context
	Db   gorm.DB
	Data *insightTypes.AnalysesProjectsIssueStatus
}

func (apiss *AnalysesProjectsIssueStatusService) Execute() error {
	projectRepo := repository.NewProjectRepository(&apiss.Ctx, &apiss.Db)

	targetProjects := []models.Project{}
	if err := projectRepo.ActiveHighPriorityProjects(&targetProjects); err != nil {
		return err
	}

	if targetProjectIds, err := helpers.Pluck(targetProjects, "Id"); err != nil {
		return err
	} else {
		targetIssueStatus := []models.IssueStatus{}
		issueStatusRepo := repository.NewIssueStatusRepository(&apiss.Ctx, &apiss.Db)
		if err := issueStatusRepo.DefaultScrum(&targetIssueStatus); err != nil {
			return err
		}

		if targetIssueStatusIds, err := helpers.Pluck(targetIssueStatus, "Id"); err != nil {
			return err
		} else {
			issueRepo := repository.NewIssueRepository(&apiss.Ctx, &apiss.Db)
			issueCountingOnProjectAndState := []repository.IssueCountingOnProjectAndState{}
			if err := issueRepo.IssueCountingOnProjectAndState(
				&issueCountingOnProjectAndState, targetIssueStatusIds, targetProjectIds,
			); err != nil {
				return err
			}

			if err := apiss.getCategories(targetProjects); err != nil {
				return err
			}

			if err := apiss.getSeries(targetIssueStatus, issueCountingOnProjectAndState); err != nil {
				return err
			}
		}
	}

	return nil
}

func (apiss *AnalysesProjectsIssueStatusService) getCategories(targetProjects []models.Project) error {
	if categories, err := helpers.Pluck(targetProjects, "Name"); err != nil {
		return err
	} else {
		for _, category := range categories {
			apiss.Data.Categories = append(apiss.Data.Categories, fmt.Sprintf("%v", category))
		}
	}
	return nil
}

func (apiss *AnalysesProjectsIssueStatusService) getSeries(
	targetIssueStatus []models.IssueStatus,
	issueCountingOnProjectAndState []repository.IssueCountingOnProjectAndState,
) error {
	projectGroupedIssueCounting := helpers.GroupByProperty(issueCountingOnProjectAndState, func(icopas repository.IssueCountingOnProjectAndState) int32 {
		return icopas.ProjectId
	})

	projectGroupData := maps.Values(projectGroupedIssueCounting)

	for _, issueStatus := range targetIssueStatus {
		dataForIssueStatus := insightTypes.AnalysesProjectsIssueStatusCounting{
			Name: issueStatus.Title,
			Data: []int32{},
		}

		for _, countingData := range projectGroupData {
			foundIdx := slices.IndexFunc(countingData, func(pis repository.IssueCountingOnProjectAndState) bool { return pis.IssueStatusId == issueStatus.Id })

			if foundIdx == -1 {
				dataForIssueStatus.Data = append(dataForIssueStatus.Data, 0)
			} else {
				dataForIssueStatus.Data = append(dataForIssueStatus.Data, int32(countingData[foundIdx].Count))
			}
		}

		apiss.Data.Series = append(apiss.Data.Series, dataForIssueStatus)
	}

	return nil
}
