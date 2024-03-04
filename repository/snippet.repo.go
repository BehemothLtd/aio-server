package repository

import (
	"aio-server/enums"
	"aio-server/gql/inputs/msInputs"
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
			db:  db,
			ctx: c,
		},
	}
}

// FindById finds a snippet by its ID.
func (r *SnippetRepository) FindById(snippet *models.Snippet, id int32) error {
	dbTables := r.db.Model(&models.Snippet{})

	return dbTables.First(&snippet, id).Error
}

// FindBySlug finds a snippet by its Slug
func (r *SnippetRepository) FindBySlug(snippet *models.Snippet, slug string) error {
	dbTables := r.db.Model(&models.Snippet{})

	return dbTables.Where("slug = ?", slug).First(&snippet).Error
}

// List retrieves a list of snippets based on provided pagination data and query.
func (r *SnippetRepository) List(
	snippets *[]*models.Snippet,
	paginateData *models.PaginationData,
	query msInputs.SnippetsQueryInput,
) error {

	return r.db.Scopes(
		helpers.Paginate(r.db.Scopes(
			r.titleLike(query.TitleCont),
		), paginateData),
	).Order("id desc").Find(&snippets).Error
}

// ListByUser retrieves a list of snippets by user.
func (r *SnippetRepository) ListByUser(
	snippets *[]*models.Snippet,
	paginateData *models.PaginationData,
	query msInputs.SelfSnippetsQueryInput,
	user *models.User,
) error {
	return r.db.Scopes(
		helpers.Paginate(r.db.Scopes(
			r.ofUser(user.Id),
			r.titleLike(query.TitleCont),
			r.snippetTypeEq(query.SnippetType),
		), paginateData),
	).Order("id desc").Find(&snippets).Error
}

// ListByUser retrieves a list of snippets by user.
func (r *SnippetRepository) ListByUserPinned(
	snippets *[]*models.Snippet,
	paginateData *models.PaginationData,
	query msInputs.SnippetsQueryInput,
	user *models.User,
) error {
	snippetsDb := r.db.Model(&models.Snippet{}).Preload("FavoritedUsers").Preload("Pins")
	dbTables := snippetsDb.InnerJoins("Pins", r.db.Where(&models.Pin{UserId: user.Id}))

	return dbTables.Scopes(
		helpers.Paginate(dbTables.Scopes(
			r.titleLike(query.TitleCont),
		), paginateData),
	).Order("id desc").Find(&snippets).Error
}

// titleLike returns a function that filters snippets by title.
func (r *SnippetRepository) titleLike(titleLike *string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if titleLike == nil {
			return db
		} else {
			return db.Where(gorm.Expr(`lower(snippets.title) LIKE ?`, strings.ToLower("%"+*titleLike+"%")))
		}
	}
}

// snippetTypeEq returns a function that filters snippets by title.
func (r *SnippetRepository) snippetTypeEq(snippetTypeEq *string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if snippetTypeEq == nil {
			return db
		} else {
			snippetTypeEqInInt, _ := enums.ParseSnippetType(*snippetTypeEq)

			return db.Where(gorm.Expr(`snippets.snippet_type = ?`, snippetTypeEqInInt))
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
	snippet.LockVersion += 1

	return r.db.Table("snippets").Omit("FavoritedUsers", "FavoritesCount", "Pins").Updates(&snippet).Error
}
