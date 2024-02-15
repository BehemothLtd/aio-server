package repository

import (
	"aio-server/models"
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// UserRepository handles operations related to users.
type UserRepository struct {
	Repository
}

// NewUserRepository initializes a new UserRepository instance.
func NewUserRepository(c *context.Context, db *gorm.DB) *Repository {
	return &Repository{
		db: db,
		ctx:  c,
	}
}

// FindByEmail finds a user by their email.
func (r *Repository) FindByEmail(user *models.User, email string) error {
	dbTables := r.db.Table("users")

	return dbTables.Where("email = ?", email).First(&user).Error
}

// Auth authenticates a user by their email and password.
func (r *Repository) Auth(email string, password string) (user *models.User, err error) {
	var u models.User

	userFindErr := r.FindByEmail(&u, email)

	if userFindErr != nil {
		return nil, errors.New("cant find user")
	}

	comparePwError := bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(password))

	if comparePwError != nil {
		return nil, errors.New("email or password is incorrect")
	}

	return &u, nil
}
