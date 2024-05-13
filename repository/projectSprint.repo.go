package repository

import (
	"aio-server/exceptions"
	"aio-server/models"
	"aio-server/pkg/constants"
	"aio-server/pkg/helpers"
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

func (r *ProjectSprintRepository) FindAllByProject(projectId int32, sprints *[]models.ProjectSprint) error {
	return r.db.Model(&models.ProjectSprint{}).
		Where("project_id = ?", projectId).
		Where("archived = false").
		Order("start_date DESC").
		Find(&sprints).Error
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

	if projectSprint.Id == *project.CurrentSprintId {
		return exceptions.NewBadRequestError("Delete project sprint is not allowed")
	}

	//To do Destroy Issue

	if err := psr.db.Delete(&projectSprint).Error; err != nil {
		return exceptions.NewBadRequestError(fmt.Sprintf("Cant delete this project sprint %s", err.Error()))
	}
	return nil
}

func (psr *ProjectSprintRepository) FindCollapsedSprints(projectSprint *models.ProjectSprint) error {
	startDate := projectSprint.StartDate.Format(constants.YYYYMMDD_DateFormat)
	endDate := projectSprint.EndDate.Format(constants.YYYYMMDD_DateFormat)

	query := psr.db.Where("project_id = ? AND ((start_date <= ? AND end_date >= ?) OR (start_date <= ? AND end_date >= ?) OR (start_date >= ? AND end_date <= ?))", projectSprint.ProjectId, startDate, startDate, endDate, endDate, startDate, endDate)

	if projectSprint.Id == 0 {
		return query.First(&models.ProjectSprint{}).Error
	} else {
		return query.Not("id = ?", projectSprint.Id).First(&models.ProjectSprint{}).Error
	}

}

func (cr *ProjectSprintRepository) Create(projectSprint *models.ProjectSprint) error {
	return cr.db.Model(&projectSprint).Preload("Project").Create(&projectSprint).First(&projectSprint).Error
}

func (psr *ProjectSprintRepository) Update(projectSprint *models.ProjectSprint, updates map[string]interface{}) error {
	if err := psr.db.Model(&projectSprint).Select(append(helpers.GetKeys(updates), "LockVersion")).Updates(updates).Error; err != nil {
		return err
	}

	return psr.db.Model(&models.ProjectSprint{}).
		Where("id = ?", projectSprint.Id).
		Preload("Project").
		First(&projectSprint).
		Error
}
