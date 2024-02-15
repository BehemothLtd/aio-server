package repository

import (
	"aio-server/models"
	"aio-server/pkg/helpers"
	"context"
	"strings"

	"gorm.io/gorm"
)

type SnippetRepository struct {
	Repository
}

// NewSnippetRepository initializes a new SnippetRepository instance.
func NewSnippetRepository(c *context.Context, db *gorm.DB) *SnippetRepository {
	return &SnippetRepository{
		Repository: Repository{
			db: db,
			ctx: c,
		},
	}
}

// FindById finds a snippet by its ID.
func (r *SnippetRepository) FindById(snippet *models.Snippet, id int32) error {
	dbTables := r.db.Model(&models.Snippet{})

	return dbTables.Preload("FavoritedUsers").First(&snippet, id).Error
}

// List retrieves a list of snippets based on provided pagination data and query.
func (r *SnippetRepository) List(
	snippets *[]*models.Snippet,
	paginateData *models.PaginationData,
	query *models.SnippetsQuery,
) error {
	dbTables := r.db.Model(&models.Snippet{})

	return dbTables.Scopes(
		helpers.Paginate(dbTables.Preload("FavoritedUsers").Scopes(
			r.titleLike(query.TitleCont),
		), paginateData),
	).Order("id desc").Find(&snippets).Error
}

// ListByUser retrieves a list of snippets by user.
func (r *SnippetRepository) ListByUser(
	snippets *[]*models.Snippet,
	paginateData *models.PaginationData,
	query *models.SnippetsQuery,
	user *models.User,
) error {
	dbTables := r.db.Model(&models.Snippet{})

	return dbTables.Scopes(
		helpers.Paginate(dbTables.Preload("FavoritedUsers").Scopes(
			r.ofUser(user.Id),
			r.titleLike(query.TitleCont),
		), paginateData),
	).Order("id desc").Find(&snippets).Error
}

// titleLike returns a function that filters snippets by title.
func (r *SnippetRepository) titleLike(titleLike string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if titleLike == "" || titleLike == "null" {
			return db
		} else {
			return db.Where(gorm.Expr(`lower(snippets.title) LIKE ?`, strings.ToLower("%"+titleLike+"%")))
		}
	}
}

// ofUser returns a function that filters snippets by user ID.
func (r *SnippetRepository) ofUser(userId int32) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(gorm.Expr(`user_id = ?`, userId))
	}
}

// Create creates a new snippet.
func (r *SnippetRepository) Create(snippet *models.Snippet) error {
	return r.db.Table("snippets").Create(&snippet).Error
}

// Update updates an existing snippet.
func (r *SnippetRepository) Update(snippet *models.Snippet) error {
	return r.db.Table("snippets").Omit("FavoritedUsers", "FavoritesCount").Updates(&snippet).Error
}
