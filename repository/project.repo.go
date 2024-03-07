package repository

import (
	"aio-server/enums"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"context"
	"strings"

	"gorm.io/gorm"
)

type ProjectRepository struct {
	Repository
}

// NewProjectRepository initializes a new ProjectRepository instance.
func NewProjectRepository(c *context.Context, db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{
		Repository: Repository{
			db:  db,
			ctx: c,
		},
	}
}

func (r *ProjectRepository) nameLike(nameLike *string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if nameLike == nil {
			return db
		} else {
			return db.Where(gorm.Expr(`lower(projects.name) LIKE ?`, strings.ToLower("%"+*nameLike+"%")))
		}
	}
}

func (r *ProjectRepository) stateEq(stateEq *string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if stateEq == nil {
			return db
		} else {
			stateInInt, err := enums.ParseUserState(*stateEq)
			if err != nil {
				return db
			}
			return db.Where(gorm.Expr(`projects.state = ?`, stateInInt))
		}
	}
}

func (r *ProjectRepository) descriptionLike(descriptionLike *string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if descriptionLike == nil {
			return db
		} else {
			return db.Where(gorm.Expr(`lower(projects.description) LIKE ?`, strings.ToLower("%"+*descriptionLike+"%")))
		}
	}
}

func (r *ProjectRepository) projectTypeEq(projectTypeLike *string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if projectTypeLike == nil {
			return db
		} else {
			stateInInt, err := enums.ParseUserState(*projectTypeLike)
			if err != nil {
				return db
			}
			return db.Where(gorm.Expr(`projects.state = ?`, stateInInt))
		}
	}
}

func (r *ProjectRepository) List(
	projects *[]*models.Project,
	paginateData *models.PaginationData,
	query insightInputs.ProjectQueryInput,
	user *models.User,
) error {
	dbTables := r.db.Model(&models.Project{})

	return dbTables.Scopes(
		helpers.Paginate(dbTables.Scopes(
			r.nameLike(query.NameCont),
			r.stateEq(query.StateCont),
			r.descriptionLike(query.DescriptionCont),
			r.projectTypeEq(query.ProjectTypeCont),
		), paginateData),
	).Order("id desc").Find(&projects).Error
}
