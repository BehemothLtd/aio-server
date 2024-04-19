package repository

import (
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"context"

	"gorm.io/gorm"
)

type ProjectExperienceRepository struct {
	Repository
}

func NewProjectExperienceRepository(c *context.Context, db *gorm.DB) *ProjectExperienceRepository {
	return &ProjectExperienceRepository{
		Repository: Repository{
			db:  db,
			ctx: c,
		},
	}
}

func (r *ProjectExperienceRepository) ListByUser(
	projectExperiences *[]*models.ProjectExperience,
	paginateData *models.PaginationData,
	query insightInputs.ProjectExperiencesQueryInput,
	user models.User,
) error {
	dbTables := r.db.Model(&models.ProjectExperience{})

	return dbTables.Scopes(
		helpers.Paginate(
			dbTables.Scopes(
				r.OfUser(user.Id),
				r.ProjectIdEq(query.ProjectIdEq),
			), paginateData,
		),
	).Order("id desc").Find(&projectExperiences).Error
}

func (r *ProjectExperienceRepository) OfUser(userId int32) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("user_id = ?", userId)
	}
}

func (r *ProjectExperienceRepository) ProjectIdEq(ProjectIdEq *int32) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if ProjectIdEq == nil {
			return db
		} else {
			return db.Where("project_id = ?", ProjectIdEq)
		}
	}
}

func (r *ProjectExperienceRepository) Find(projectExperience *models.ProjectExperience) error {
	dbTables := r.db.Model(&models.ProjectExperience{})

	return dbTables.Where(&projectExperience).First(&projectExperience).Error
}
