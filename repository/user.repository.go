package repository

import (
	"aio-server/models"
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func NewUserRepository(c *context.Context, db *gorm.DB) *Repository {
	return &Repository{
		db: db,
		c:  c,
	}
}

func (r *Repository) FindUserByEmail(user *models.User, email string) error {
	dbTables := r.db.Table("users")

	return dbTables.Where("email = ?", email).First(&user).Error
}

func (r *Repository) AuthUser(email string, password string) (user *models.User, err error) {
	if email == "" {
		return nil, errors.New("email is required")
	}

	if password == "" {
		return nil, errors.New("password is required")
	}

	var u models.User

	userFindErr := r.FindUserByEmail(&u, email)

	if userFindErr != nil {
		return nil, userFindErr
	}

	comparePwError := bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(password))

	if comparePwError != nil {
		return nil, comparePwError
	}

	return &u, nil
}

func (r *Repository) FindUserById(user *models.User, uid any) error {
	dbTables := r.db.Table("users")

	return dbTables.First(&user, uid).Error
}
