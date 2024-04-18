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

func (r *ProjectExperienceRepository) List(
	projectExperiences *[]*models.ProjectExperience,
	paginateData *models.PaginationData,
	query insightInputs.AttendancesQueryInput,
) error {
	dbTables := r.db.Model(&models.ProjectExperience{})

	return dbTables.Scopes(
		helpers.Paginate(
			dbTables.Scopes(
				r.projectId(query.UserIdEq),
			), paginateData,
		),
	).Order("id desc").Find(&attendances).Error
}
