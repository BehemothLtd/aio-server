package repository

import (
	"aio-server/models"
	"context"

	"gorm.io/gorm"
)

// SnippetsTagRepository is db handles type for SnippetsTag table
type SnippetsTagRepository struct {
	Repository
}

// NewSnippetsTagRepository initializes a new SnippetsTagRepository instance.
func NewSnippetsTagRepository(c *context.Context, db *gorm.DB) *SnippetsTagRepository {
	return &SnippetsTagRepository{
		Repository: Repository{
			db:  db,
			ctx: c,
		},
	}
}

// FindBySnippetAndTag finds a snippetsTag by its snippet and tag.
func (str *SnippetsTagRepository) FindBySnippetAndTag(snippetsTag *models.SnippetsTag) error {
	dbTable := str.db.Model(&models.SnippetsTag{})

	return dbTable.Where("snippet_id = ? AND tag_id = ?", snippetsTag.SnippetId, snippetsTag.TagId).First(&snippetsTag).Error
}

func (str *SnippetsTagRepository) Create(snippetsTag *models.SnippetsTag) error {
	return str.db.Table("snippets_tags").Create(&snippetsTag).Error
}
