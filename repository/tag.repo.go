package repository

import (
	"aio-server/models"
	"context"

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
