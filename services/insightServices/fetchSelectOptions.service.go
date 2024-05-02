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
	return []string{
		"issueStatus",
		"developmentRole",
		"user",
		"project",
		"client",
	}
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
				case "project":
					if err := fsos.handleProjectOptions(); err != nil {
						return nil
					}
				case "client":
					if err := fsos.handleClientOptions(); err != nil {
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
				Value: issueStatus.Id,
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
			Value: developmentRoles[i].Id,
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
			Value: user.Id,
		})
	}

	return nil
}

func (fsos *FetchSelectOptionsService) handleProjectOptions() error {
	projects := []*models.Project{}
	repo := repository.NewProjectRepository(nil, fsos.Db)

	if err := repo.All(&projects); err != nil {
		return err
	}

	for _, prj := range projects {
		fsos.Result.ProjectOptions = append(fsos.Result.ProjectOptions, insightTypes.CommonSelectOption{
			Label: prj.Name,
			Value: prj.Id,
		})
	}

	return nil
}

func (fsos *FetchSelectOptionsService) handleClientOptions() error {
	clients := []*models.Client{}
	repo := repository.NewClientRepository(nil, fsos.Db)

	if err := repo.All(&clients); err != nil {
		return err
	}

	for _, client := range clients {
		fsos.Result.ClientOptions = append(fsos.Result.ClientOptions, insightTypes.CommonSelectOption{
			Label: client.Name,
			Value: client.Id,
		})
	}

	return nil
}
