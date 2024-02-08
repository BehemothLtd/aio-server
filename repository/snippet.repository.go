package repository

import (
	"aio-server/models"
	"aio-server/pkg/helpers"
	"context"
	"strings"

	"gorm.io/gorm"
)

func NewSnippetRepository(c *context.Context, db *gorm.DB) *Repository {
	return &Repository{
		db: db,
		c:  c,
	}
}

func (r *Repository) ListSnippets(
	snippets *[]*models.Snippet, paginateData *models.PaginationData,
	query *models.SnippetsQuery,
) error {
	dbTables := r.db.Model(&models.Snippet{})

	return dbTables.Scopes(
		helpers.Paginate(dbTables.Preload("FavoritedUsers").Scopes(
			r.titleLike(query.TitleCont),
		), paginateData),
	).Order("id desc").Find(&snippets).Error
}

func (r *Repository) FindSnippetById(snippet *models.Snippet, id int32) error {
	dbTables := r.db.Model(&models.Snippet{})

	return dbTables.Preload("FavoritedUsers").First(&snippet, id).Error
}

func (r *Repository) titleLike(titleLike string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if titleLike == "" || titleLike == "null" {
			return db
		} else {
			return db.Where(gorm.Expr(`lower(snippets.title) LIKE ?`, strings.ToLower("%"+titleLike+"%")))
		}
	}
}

func (r *Repository) ofUser(userId int32) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(gorm.Expr(`user_id = ?`, userId))
	}
}

func (r *Repository) ListSnippetsOfUser(
	snippets *[]*models.Snippet, paginateData *models.PaginationData,
	user *models.User,
) error {
	dbTables := r.db.Table("snippets")

	return dbTables.Scopes(
		r.ofUser(user.Id),
		helpers.Paginate(dbTables.Scopes(), paginateData),
	).Order("id desc").Find(&snippets).Error
}

func (r *Repository) CreateSnippet(snippet *models.Snippet) error {
	dbTables := r.db.Table("snippets")

	return dbTables.Create(&snippet).Error
}

func (r *Repository) UpdateSnippet(snippet *models.Snippet) error {
	dbTables := r.db.Model(&snippet)

	return dbTables.Updates(&snippet).Error
}
