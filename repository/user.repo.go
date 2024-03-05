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
func NewUserRepository(c *context.Context, db *gorm.DB) *UserRepository {
	return &UserRepository{
		Repository: Repository{
			db:  db,
			ctx: c,
		},
	}
}

// Find finds a user by their attribute.
func (r *UserRepository) Find(user *models.User) error {
	dbTables := r.db.Table("users")

	return dbTables.Where(&user).First(&user).Error
}

// FindWithAvatar finds an user includes his avatar data
func (r *UserRepository) FindWithAvatar(user *models.User) error {
	dbTables := r.db.Table("users").Preload("Avatar.AttachmentBlob")

	return dbTables.Where("id = ?", user.Id).First(&user).Error
}

// FindByEmail finds a user by their email.
func (r *UserRepository) FindByEmail(user *models.User, email string) error {
	dbTables := r.db.Table("users")

	return dbTables.Where("email = ?", email).First(&user).Error
}

// Auth authenticates a user by their email and password.
func (r *UserRepository) Auth(email string, password string) (user *models.User, err error) {
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

// Update updates an user by its assigned attributes
func (r *UserRepository) Update(user *models.User, fields []string) error {
	// TODO: handle NULL value save into DB
	return r.db.Model(&user).Select(fields).Session(&gorm.Session{FullSaveAssociations: true}).Updates(&user).Error
}
