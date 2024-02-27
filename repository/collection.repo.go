package repository

import (
	"aio-server/models"
	"aio-server/pkg/helpers"
	"context"
	"strings"

	"gorm.io/gorm"
)

type CollectionRepository struct {
	Repository
}

// Innitializes a new CollectionRepository instance.
func NewCollectionRepository(c *context.Context, db *gorm.DB) *CollectionRepository {
	return &CollectionRepository{
		Repository: Repository{
			db:  db,
			ctx: c,
		},
	}
}

func (r *CollectionRepository) List(
	collections *[]*models.Collection,
	paginateData *models.PaginationData,
	query *models.CollectionQuery,
) error {
	dbTables := r.db.Model(&models.Collection{})

	return dbTables.Scopes(
		helpers.Paginate(dbTables.Scopes(
			r.titleLike(query.TitleCont),
		), paginateData),
	).Order("id desc").Find(&collections).Error
}

func (r *CollectionRepository) titleLike(titleLike string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if titleLike == "" || titleLike == "null" {
			return db
		} else {
			return db.Where(gorm.Expr(`lower(collections.title) LIKE ?`, strings.ToLower("%"+titleLike+"%")))
		}
	}
}
