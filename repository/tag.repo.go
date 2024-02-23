package repository

import (
	"aio-server/models"
	"aio-server/pkg/helpers"
	"context"
	"strings"

	"gorm.io/gorm"
)

type TagRepository struct {
	Repository
}

// NewTagRepository initializes a new TagRepository instance.
func NewTagRepository(c *context.Context, db *gorm.DB) *TagRepository {
	return &TagRepository{
		Repository: Repository{
			db:  db,
			ctx: c,
		},
	}
}

// FindById finds a snippet by its ID.
func (r *TagRepository) FindById(tag *models.Tag, id int32) error {
	dbTables := r.db.Model(&models.Tag{})

	return dbTables.First(&tag, id).Error
}

// List retrieves a list of snippets based on provided pagination data and query.
func (r *TagRepository) List(
	tags *[]*models.Tag,
	paginateData *models.PaginationData,
	query *models.TagsQuery,
) error {
	dbTables := r.db.Model(&models.Tag{})

	return dbTables.Scopes(
		helpers.Paginate(dbTables.Scopes(
			r.nameLike(query.NameCont),
		), paginateData),
	).Order("id desc").Find(&tags).Error
}

// nameLike returns a function that filters snippets by title.
func (r *TagRepository) nameLike(nameLike string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if nameLike == "" || nameLike == "null" {
			return db
		} else {
			return db.Where(gorm.Expr(`lower(tags.name) LIKE ?`, strings.ToLower("%"+nameLike+"%")))
		}
	}
}

// Create creates a new tag.
func (tr *TagRepository) Create(tag *models.Tag) error {
	return tr.db.Table("tags").Create(&tag).Error
}

// Update updates an existing tag.
func (tr *TagRepository) Update(tag *models.Tag) error {
	tag.LockVersion += 1

	return tr.db.Table("tags").Updates(&tag).Error
}

// FindByEmail finds a user by its email.
func (tr *TagRepository) FindByName(tag *models.Tag, name string) error {
	dbTables := tr.db.Model(&models.Tag{})

	return dbTables.Where("name = ?", name).First(&tag).Error
}
