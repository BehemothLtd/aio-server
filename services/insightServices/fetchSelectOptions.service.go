package insightServices

import (
	"aio-server/gql/gqlTypes/insightTypes"
	"aio-server/models"
	"aio-server/repository"
	"fmt"
	"slices"

	"gorm.io/gorm"
)

func availableKeys() []string {
	return []string{"issueStatus"}
}

type FetchSelectOptionsService struct {
	Db   *gorm.DB
	Keys *[]string

	Result insightTypes.SelectOptionsType
}

func (fsos *FetchSelectOptionsService) Execute() error {
	var availableKeys = availableKeys()

	if fsos.Keys != nil {
		for _, key := range *fsos.Keys {
			if slices.Contains(availableKeys, key) {
				switch key {
				case "issueStatus":
					if err := fsos.handleIssueStatusOptions(); err != nil {
						return nil
					}
				}
			} else {
				return fmt.Errorf("invalid key %s", key)
			}
		}
	}

	return nil
}

func (fsos *FetchSelectOptionsService) handleIssueStatusOptions() error {
	issueStatuses := []*models.IssueStatus{}
	repo := repository.NewIssueStatusRepository(nil, fsos.Db)

	if err := repo.All(&issueStatuses); err != nil {
		return err
	}

	for _, issueStatus := range issueStatuses {
		fsos.Result.IssueStatusOptions = append(fsos.Result.IssueStatusOptions, insightTypes.IssueStatusSelectOption{
			CommonSelectOption: insightTypes.CommonSelectOption{
				Label: issueStatus.Title,
				Value: fmt.Sprintf("%d", issueStatus.Id),
			},
			Color: issueStatus.Color,
		})
	}

	return nil
}
