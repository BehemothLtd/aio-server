package repository

import (
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"context"
	"strings"

	"gorm.io/gorm"
)

type ClientRepository struct {
	Repository
}

// NewClientRepository initializes a new ClientRepository instance.
func NewClientRepository(c *context.Context, db *gorm.DB) *ClientRepository {
	return &ClientRepository{
		Repository: Repository{
			db:  db,
			ctx: c,
		},
	}
}

// List retrieves a list of clients based on provided pagination data and query.
func (r *ClientRepository) List(
	clients *[]*models.Client,
	paginateData *models.PaginationData,
	query insightInputs.ClientsQueryInput,
) error {
	dbTables := r.db.Model(&models.Client{})

	return dbTables.Scopes(
		helpers.Paginate(dbTables.Scopes(
			r.nameLike(query.NameCont),
		), paginateData),
	).Order("id desc").Find(&clients).Error
}

// nameLike returns a function that filters client by name.
func (r *ClientRepository) nameLike(nameLike *string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if nameLike == nil {
			return db
		} else {
			return db.Where(gorm.Expr(`lower(clients.name) LIKE ?`, strings.ToLower("%"+*nameLike+"%")))
		}
	}
}

func (r *ClientRepository) Find(client *models.Client) error {
	dbTables := r.db.Model(&models.Client{})

	return dbTables.Where(&client).First(&client).Error
}

func (r *ClientRepository) Create(client *models.Client) error {
	return r.db.Model(&client).Create(&client).First(&client).Error
}

func (r *ClientRepository) Update(client *models.Client, updateClient models.Client) error {
	if err := r.db.Model(&client).Updates(&updateClient).Error; err != nil {
		return err
	}

	return r.db.Model(&client).Where("id = ?", client.Id).First(&client).Error
}

func (r *ClientRepository) Destroy(client *models.Client) error {
	return r.db.Table("clients").Delete(&client).Error
}

func (r *ClientRepository) All(clients *[]*models.Client) error {
	return r.db.Table("clients").Order("id ASC").Find(&clients).Error
}
