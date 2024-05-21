package insightServices

import (
	"aio-server/database"
	"aio-server/enums"
	"aio-server/exceptions"
	"aio-server/gql/gqlTypes/insightTypes"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/pkg/systems"
	"aio-server/repository"
	"errors"
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
		"deviceType",
		"issueType",
		"issuePriority",
		"projectIssue",
		"projectIssueStatus",
	}
}

type FetchSelectOptionsService struct {
	Db     *gorm.DB
	Keys   *[]string
	Params *insightInputs.SelectOptionsParamsType

	Result insightTypes.SelectOptionsType
}

func (service *FetchSelectOptionsService) Execute() error {
	var availableKeys = availableKeys()

	if service.Keys != nil {
		for _, key := range *service.Keys {
			if slices.Contains(availableKeys, key) {
				switch key {
				case "issueStatus":
					if err := service.handleIssueStatusOptions(); err != nil {
						return err
					}
				case "developmentRole":
					if err := service.handleDevelopmentRoleOptions(); err != nil {
						return err
					}
				case "user":
					if err := service.handleUserOptions(); err != nil {
						return err
					}
				case "project":
					if err := service.handleProjectOptions(); err != nil {
						return err
					}
				case "client":
					if err := service.handleClientOptions(); err != nil {
						return err
					}
				case "deviceType":
					if err := service.handleDeviceTypeOptions(); err != nil {
						return err
					}
				case "issueType":
					if err := service.handleIssueTypeOptions(); err != nil {
						return err
					}
				case "issuePriority":
					if err := service.handleIssuePriorityOptions(); err != nil {
						return err
					}
				case "projectIssue":
					if err := service.handleProjectIssueOptions(); err != nil {
						return err
					}
				case "projectIssueStatus":
					if err := service.handleProjectIssueStatusOptions(); err != nil {
						return err
					}
				}
			} else {
				return fmt.Errorf("invalid key %s", key)
			}
		}
	}

	return nil
}

func (service *FetchSelectOptionsService) handleIssueStatusOptions() error {
	issueStatuses := []*models.IssueStatus{}
	repo := repository.NewIssueStatusRepository(nil, service.Db)

	if err := repo.All(&issueStatuses); err != nil {
		return err
	}

	for _, issueStatus := range issueStatuses {
		service.Result.IssueStatusOptions = append(service.Result.IssueStatusOptions, insightTypes.IssueStatusSelectOption{
			CommonSelectOption: insightTypes.CommonSelectOption{
				Label: issueStatus.Title,
				Value: issueStatus.Id,
			},
			Color: issueStatus.Color,
		})
	}

	return nil
}

func (service *FetchSelectOptionsService) handleDeviceTypeOptions() error {
	deviceTypes := []*models.DeviceType{}
	repo := repository.NewDeviceTypeRepository(nil, service.Db)

	if err := repo.All(&deviceTypes); err != nil {
		return err
	}

	for _, deviceType := range deviceTypes {
		service.Result.DeviceTypeOptions = append(service.Result.DeviceTypeOptions, insightTypes.DeviceTypeSelectOption{
			Label: deviceType.Name,
			Value: deviceType.Id,
		})
	}

	return nil
}

func (service *FetchSelectOptionsService) handleDevelopmentRoleOptions() error {
	developmentRoles := systems.GetDevelopmentRoles()

	for i := range developmentRoles {
		service.Result.DevelopmentRoleOptions = append(service.Result.DevelopmentRoleOptions, insightTypes.CommonSelectOption{
			Label: developmentRoles[i].Title,
			Value: developmentRoles[i].Id,
		})
	}

	return nil
}

func (service *FetchSelectOptionsService) handleUserOptions() error {
	users := []*models.User{}
	repo := repository.NewUserRepository(nil, service.Db)

	if err := repo.All(&users); err != nil {
		return err
	}

	for _, user := range users {
		service.Result.UserOptions = append(service.Result.UserOptions, insightTypes.CommonSelectOption{
			Label: user.Name,
			Value: user.Id,
		})
	}

	return nil
}

func (service *FetchSelectOptionsService) handleProjectOptions() error {
	projects := []*models.Project{}
	repo := repository.NewProjectRepository(nil, service.Db)

	if err := repo.All(&projects); err != nil {
		return err
	}

	for _, prj := range projects {
		service.Result.ProjectOptions = append(service.Result.ProjectOptions, insightTypes.CommonSelectOption{
			Label: prj.Name,
			Value: prj.Id,
		})
	}

	return nil
}

func (service *FetchSelectOptionsService) handleClientOptions() error {
	clients := []*models.Client{}
	repo := repository.NewClientRepository(nil, service.Db)

	if err := repo.All(&clients); err != nil {
		return err
	}

	for _, client := range clients {
		service.Result.ClientOptions = append(service.Result.ClientOptions, insightTypes.CommonSelectOption{
			Label: client.Name,
			Value: client.Id,
		})
	}

	return nil
}

func (service *FetchSelectOptionsService) handleIssueTypeOptions() error {
	for issueType := range enums.SqlIntIssueTypeValue {
		service.Result.IssueTypeOptions = append(service.Result.IssueTypeOptions, insightTypes.StringStringSelectOption{
			Label: issueType.String(),
			Value: issueType.String(),
		})
	}

	return nil
}

func (service *FetchSelectOptionsService) handleIssuePriorityOptions() error {
	for issuePriority := range enums.SqlIntIssuePriorityValue {
		service.Result.IssuePriorityOptions = append(service.Result.IssuePriorityOptions, insightTypes.StringStringSelectOption{
			Label: issuePriority.String(),
			Value: issuePriority.String(),
		})
	}

	return nil
}

func (service *FetchSelectOptionsService) handleProjectIssueOptions() error {
	if service.Params != nil && service.Params.ProjectId != nil && *service.Params.ProjectId != "" {
		projectId, err := helpers.GqlIdToInt32(*service.Params.ProjectId)
		if err != nil || projectId == 0 {
			return exceptions.NewBadRequestError("Invalid Id")
		}

		project := models.Project{Id: projectId}
		projectRepo := repository.NewProjectRepository(nil, database.Db)
		if err := projectRepo.Find(&project); err != nil {
			return exceptions.NewBadRequestError("Invalid projectId Provided")
		}

		issues := []*models.Issue{}
		repo := repository.NewIssueRepository(nil, database.Db)
		if err := repo.AllByProjectId(&issues, projectId); err != nil {
			return exceptions.NewBadRequestError(err.Error())
		}

		for _, issue := range issues {
			service.Result.ProjectIssueOptions = append(service.Result.ProjectIssueOptions, insightTypes.CommonSelectOption{
				Label: issue.Title,
				Value: issue.Id,
			})
		}

		return nil
	}

	return errors.New("projectId is required for projectIssueOptions")
}

func (service *FetchSelectOptionsService) handleProjectIssueStatusOptions() error {
	if service.Params != nil && service.Params.ProjectId != nil && *service.Params.ProjectId != "" {
		projectId, err := helpers.GqlIdToInt32(*service.Params.ProjectId)
		if err != nil || projectId == 0 {
			return exceptions.NewBadRequestError("Invalid Id")
		}

		project := models.Project{Id: projectId}
		projectRepo := repository.NewProjectRepository(nil, database.Db)
		if err := projectRepo.Find(&project); err != nil {
			return exceptions.NewBadRequestError("Invalid projectId Provided")
		}

		issueStatuses := []*models.IssueStatus{}
		repo := repository.NewIssueStatusRepository(nil, database.Db)
		if err := repo.FetchAllOnProject(projectId, &issueStatuses); err != nil {
			return exceptions.NewBadRequestError(err.Error())
		}

		for _, issue := range issueStatuses {
			service.Result.ProjectIssueStatusOptions = append(service.Result.ProjectIssueStatusOptions, insightTypes.CommonSelectOption{
				Label: issue.Title,
				Value: issue.Id,
			})
		}

		return nil
	}

	return errors.New("projectId is required for projectIssueOptions")
}
