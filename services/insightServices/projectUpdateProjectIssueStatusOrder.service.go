package insightServices

import (
	"aio-server/exceptions"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"slices"

	"github.com/graph-gophers/graphql-go"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type ProjectUpdateProjectIssueStatusOrderService struct {
	Ctx  *context.Context
	Db   *gorm.DB
	Args struct {
		Id    graphql.ID
		Input []int32
	}
	Project *models.Project
}

func (pupisos *ProjectUpdateProjectIssueStatusOrderService) Execute() error {
	if pupisos.Args.Id == "" {
		return exceptions.NewBadRequestError("Invalid Id")
	}

	projectId, err := helpers.GqlIdToInt32(pupisos.Args.Id)

	if err != nil {
		return exceptions.NewBadRequestError("Invalid Id")
	}

	project := models.Project{Id: projectId}

	projectRepo := repository.NewProjectRepository(pupisos.Ctx, pupisos.Db.Preload("ProjectIssueStatuses"))
	if err := projectRepo.Find(&project); err != nil {
		return exceptions.NewRecordNotFoundError()
	}

	inputIds := slices.Compact(pupisos.Args.Input)

	if len(project.ProjectIssueStatuses) != len(inputIds) {
		return exceptions.NewBadRequestError("Invalid input")
	}

	currentIds := []int32{}
	projectIssueStatusRepo := repository.NewProjectIssueStatusRepository(pupisos.Ctx, pupisos.Db)

	if err := projectIssueStatusRepo.FindIdsByProjectId(projectId, inputIds, &currentIds); err != nil {
		return exceptions.NewBadRequestError("Invalid input's id")
	}

	if len(currentIds) != len(inputIds) {
		return exceptions.NewBadRequestError("Invalid input's id")
	}

	if err := projectIssueStatusRepo.UpdateBatchOfNewPositionsForAProject(projectId, inputIds); err != nil {
		return exceptions.NewUnprocessableContentError("Cant update statuses position", nil)
	}

	return nil
}
