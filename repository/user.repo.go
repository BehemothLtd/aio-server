package repository

import (
	"aio-server/enums"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"context"
	"errors"
	"strings"

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
	stateActive := enums.UserStateActive
	u := models.User{Email: email, State: stateActive}

	userFindErr := r.Find(&u)

	if userFindErr != nil {
		return nil, errors.New("cant find user")
	}

	comparePwError := bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(password))

	if comparePwError != nil {
		return nil, errors.New("email or password is incorrect")
	}

	return &u, nil
}

func (r *Repository) List(
	users *[]*models.User,
	paginateData *models.PaginationData,
	query insightInputs.UserQueryInput,
) error {
	dbTables := r.db.Model(&models.User{})

	return dbTables.Scopes(
		helpers.Paginate(dbTables.Scopes(
			r.nameLike(query.NameCont),
			r.fullNameLike(query.FullNameCont),
			r.emailLike(query.EmailCont),
			r.slackIdLike(query.SlackIdCont),
			r.stateEq(query.StateEq),
		), paginateData),
	).Order("id desc").Find(&users).Error
}

func (r *Repository) nameLike(nameLike *string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if nameLike == nil {
			return db
		} else {
			return db.Where(gorm.Expr(`lower(users.name) LIKE ?`, strings.ToLower("%"+*nameLike+"%")))
		}
	}
}

func (r *Repository) fullNameLike(fullNameLike *string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if fullNameLike == nil {
			return db
		} else {
			return db.Where(gorm.Expr(`lower(users.full_name) LIKE ?`, strings.ToLower("%"+*fullNameLike+"%")))
		}
	}
}

func (r *Repository) emailLike(emailLike *string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if emailLike == nil {
			return db
		} else {
			return db.Where(gorm.Expr(`lower(users.email) LIKE ?`, strings.ToLower("%"+*emailLike+"%")))
		}
	}
}

func (r *Repository) slackIdLike(slackIdLike *string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if slackIdLike == nil {
			return db
		} else {
			return db.Where(gorm.Expr(`lower(users.slack_id) LIKE ?`, strings.ToLower("%"+*slackIdLike+"%")))
		}
	}
}

func (r *Repository) stateEq(stateEq *string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if stateEq == nil {
			return db
		} else {
			stateInInt, err := enums.ParseUserState(*stateEq)
			if err != nil {
				return db
			}
			return db.Where(gorm.Expr(`users.state = ?`, stateInInt))
		}
	}
}

// Update updates an user by its assigned attributes
func (r *UserRepository) Update(user *models.User, fields []string) error {
	// TODO: handle NULL value save into DB
	return r.db.Model(&user).Select(fields).Session(&gorm.Session{FullSaveAssociations: true}).Updates(&user).Error
}

func (r *UserRepository) All(users *[]*models.User) error {
	return r.db.Table("users").Order("id ASC").Find(&users).Error
}
