package insightServices

import (
	"aio-server/gql/gqlTypes/insightTypes"
	"aio-server/models"
	"aio-server/pkg/systems"
	"aio-server/repository"
	"fmt"
	"slices"

	"gorm.io/gorm"
)

func availableKeys() []string {
	return []string{"issueStatus", "developmentRole", "user"}
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
				case "developmentRole":
					if err := fsos.handleDevelopmentRoleOptions(); err != nil {
						return nil
					}
				case "user":
					if err := fsos.handleUserOptions(); err != nil {
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

func (fsos *FetchSelectOptionsService) handleDevelopmentRoleOptions() error {
	developmentRoles := systems.GetDevelopmentRoles()

	for i := range developmentRoles {
		fsos.Result.DevelopmentRoleOptions = append(fsos.Result.DevelopmentRoleOptions, insightTypes.CommonSelectOption{
			Label: developmentRoles[i].Title,
			Value: fmt.Sprintf("%d", developmentRoles[i].Id),
		})
	}

	return nil
}

func (fsos *FetchSelectOptionsService) handleUserOptions() error {
	users := []*models.User{}
	repo := repository.NewUserRepository(nil, fsos.Db)

	if err := repo.All(&users); err != nil {
		return err
	}

	for _, user := range users {
		fsos.Result.UserOptions = append(fsos.Result.UserOptions, insightTypes.CommonSelectOption{
			Label: user.Name,
			Value: fmt.Sprintf("%d", user.Id),
		})
	}

	return nil
}
