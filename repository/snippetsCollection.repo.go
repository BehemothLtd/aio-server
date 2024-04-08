package repository

import (
	"aio-server/models"
	"context"

	"gorm.io/gorm"
)

type SnippetsCollectionRepository struct {
	Repository
}

func NewSnippetsCollectionRepository(c *context.Context, db *gorm.DB) *SnippetsCollectionRepository {
	return &SnippetsCollectionRepository{
		Repository: Repository{
			db:  db,
			ctx: c,
		},
	}
}

func (str *SnippetsCollectionRepository) Find(snippetsCollection *models.SnippetsCollection) error {
	return str.db.Model(&models.SnippetsCollection{}).Where(&snippetsCollection).First(&snippetsCollection).Error
}

func (str *SnippetsCollectionRepository) Create(snippetsCollection *models.SnippetsCollection) error {
	return str.db.Table("snippets_collections").Create(&snippetsCollection).Error
}

func (str *SnippetsCollectionRepository) Delete(snippetsCollection *models.SnippetsCollection) error {
	return str.db.Table("snippets_collections").Delete(&snippetsCollection).Error
}
