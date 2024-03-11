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

// Querying Functions
func (cr *CollectionRepository) List(
	collections *[]*models.Collection,
	paginateData *models.PaginationData,
	query msInputs.CollectionQueryInput,
	user *models.User,
) error {
	dbTables := cr.db.Model(&models.Collection{}).Preload("Snippets")

	return dbTables.Scopes(
		helpers.Paginate(dbTables.Scopes(
			cr.ofParent("user_id", user.Id),
			cr.stringLike("collections", "title", query.TitleCont),
		), paginateData),
	).Order("id desc").Find(&collections).Error
}

func (cr *CollectionRepository) FindByUser(collection *models.Collection, userId int32) error {
	dbTables := cr.db.Model(&models.Collection{})

	return dbTables.Where(gorm.Expr(`user_id = ?`, userId)).First(&collection).Error
}

// Mutating Functions
func (cr *CollectionRepository) Create(collection *models.Collection) error {
	return cr.db.Table("collections").Create(&collection).Error
}

func (cr *CollectionRepository) Update(collection *models.Collection) error {
	return cr.db.Table("collections").Updates(&collection).Error
}

func (cr *Repository) Delete(collection *models.Collection) error {
	err := cr.db.Transaction(
		func(cx *gorm.DB) error {
			if err := cx.Delete(&collection).Error; err != nil {
				return err
			}

			if err := cx.Delete(&models.SnippetsCollection{}, "collection_id = ?", collection.Id).Error; err != nil {
				return err
			}

			return nil
		})

	return err
}
