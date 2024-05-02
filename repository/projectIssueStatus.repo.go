package repository

import (
	"aio-server/models"
	"context"

	"gorm.io/gorm"
)

type ProjectIssueStatusRepository struct {
	Repository
}

func NewProjectIssueStatusRepository(c *context.Context, db *gorm.DB) *ProjectIssueStatusRepository {
	return &ProjectIssueStatusRepository{
		Repository: Repository{
			db:  db,
			ctx: c,
		},
	}
}

func (r *ProjectIssueStatusRepository) FindIdsByProjectId(projectId int32, projectIssueStatusIds []int32, result *[]int32) error {
	return r.db.Model(&models.ProjectIssueStatus{}).
		Select("id").Where("project_id = ?", projectId).
		Where("id IN (?)", projectIssueStatusIds).
		Scan(&result).
		Error
}

func (r *ProjectIssueStatusRepository) UpdateBatchOfNewPositionsForAProject(projectId int32, ids []int32) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		for index, id := range ids {
			if err := tx.Model(&models.ProjectIssueStatus{}).Where("project_id = ? AND id = ?", projectId, id).
				Update("position", index+1).
				Error; err != nil {
				// return any error will rollback
				return err
			}
		}

		// return nil will commit the whole transaction
		return nil
	})
}
