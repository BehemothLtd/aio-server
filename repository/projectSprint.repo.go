package repository

import (
	"aio-server/exceptions"
	"aio-server/models"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type ProjectSprintRepository struct {
	Repository
}

func NewProjectSprintRepository(c *context.Context, db *gorm.DB) *ProjectSprintRepository {
	return &ProjectSprintRepository{
		Repository: Repository{
			db:  db,
			ctx: c,
		},
	}
}

func (r *ProjectSprintRepository) Find(projectSprint *models.ProjectSprint) error {
	dbTables := r.db.Model(&models.ProjectSprint{})

	return dbTables.Where(&projectSprint).First(&projectSprint).Error
}

func (psr *ProjectSprintRepository) Destroy(projectSprint *models.ProjectSprint) error {

	project := models.Project{Id: projectSprint.ProjectId}
	repo := NewProjectRepository(psr.ctx, psr.db)
	err := repo.Find(&project)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return exceptions.NewRecordNotFoundError()
		}
		return err
	}

	if projectSprint.ProjectId == int32(project.CurrentSprintId) {
		return exceptions.NewBadRequestError("Delete project sprint is not allowed")
	}

	//To do Destroy Issue

	if err := psr.db.Delete(&projectSprint).Error; err != nil {
		return exceptions.NewBadRequestError(fmt.Sprintf("Cant delete this project sprint %s", err.Error()))
	}
	return nil
}
