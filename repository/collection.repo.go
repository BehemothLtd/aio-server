package repository

import (
	"aio-server/gql/inputs/msInputs"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"context"

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
	query msInputs.CollectionQueryInput,
	user *models.User,
) error {
	dbTables := r.db.Model(&models.Collection{})

	return dbTables.Scopes(
		helpers.Paginate(dbTables.Scopes(
			r.ofParent("user_id", user.Id),
			r.stringLike("collections", "title", *query.TitleCont),
		), paginateData),
	).Order("id desc").Find(&collections).Error
}
