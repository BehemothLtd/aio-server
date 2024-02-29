package repository

import (
	"aio-server/models"
	"aio-server/pkg/helpers"
	"context"
	"strings"

	"gorm.io/gorm"
)

type UserGroupRepository struct {
	Repository
}

// NewUserGroupRepository initializes a new UserGroupRepository instance.
func NewUserGroupRepository(c *context.Context, db *gorm.DB) *UserGroupRepository {
	return &UserGroupRepository{
		Repository: Repository{
			db:  db,
			ctx: c,
		},
	}
}

// List retrieves a list of user groups based on provided pagination data and query.
func (r *UserGroupRepository) List(
	userGroups *[]*models.UserGroup,
	paginateData *models.PaginationData,
	query *models.UserGroupsQuery,
) error {
	dbTables := r.db.Model(&models.UserGroup{})

	return dbTables.Scopes(
		helpers.Paginate(dbTables.Scopes(
			r.titleLike(query.TitleCont),
		), paginateData),
	).Order("id desc").Find(&userGroups).Error
}

// titleLike returns a function that filters user groups by title.
func (r *UserGroupRepository) titleLike(titleLike string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if titleLike == "" || titleLike == "null" {
			return db
		} else {
			return db.Where(gorm.Expr(`lower(user_groups.title) LIKE ?`, strings.ToLower("%"+titleLike+"%")))
		}
	}
}

// FindById finds a userGroup by its ID.
func (r *UserGroupRepository) FindById(userGroup *models.UserGroup, id int32) error {
	dbTables := r.db.Model(&models.UserGroup{})

	return dbTables.First(&userGroup, id).Error
}
