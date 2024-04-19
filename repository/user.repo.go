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
	dbTables := r.db.Table("users").Preload("Avatar", "name = 'avatar'").Preload("Avatar.AttachmentBlob").Preload("ProjectAssignees")

	return dbTables.Where("id = ?", user.Id).First(&user).Error
}

func (r *Repository) FindWithProjectAssignees(user *models.User) error {
	dbTables := r.db.Table("users").Preload("ProjectAssignees")

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

func (r *UserRepository) List(
	users *[]*models.User,
	paginateData *models.PaginationData,
	query insightInputs.UserQueryInput,
) error {
	dbTables := r.db.Model(&models.User{})

	return dbTables.
		Scopes(
			helpers.Paginate(dbTables.Scopes(
				r.nameLike(query.NameCont),
				r.fullNameLike(query.FullNameCont),
				r.emailLike(query.EmailCont),
				r.slackIdLike(query.SlackIdCont),
				r.stateEq(query.StateEq),
			), paginateData),
		).
		Preload("Avatar", "name = 'avatar'").
		Preload("Avatar.AttachmentBlob").
		Order("id desc").Find(&users).Error
}

func (r *UserRepository) nameLike(nameLike *string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if nameLike == nil {
			return db
		} else {
			return db.Where(gorm.Expr(`lower(users.name) LIKE ?`, strings.ToLower("%"+*nameLike+"%")))
		}
	}
}

func (r *UserRepository) fullNameLike(fullNameLike *string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if fullNameLike == nil {
			return db
		} else {
			return db.Where(gorm.Expr(`lower(users.full_name) LIKE ?`, strings.ToLower("%"+*fullNameLike+"%")))
		}
	}
}

func (r *UserRepository) emailLike(emailLike *string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if emailLike == nil {
			return db
		} else {
			return db.Where(gorm.Expr(`lower(users.email) LIKE ?`, strings.ToLower("%"+*emailLike+"%")))
		}
	}
}

func (r *UserRepository) slackIdLike(slackIdLike *string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if slackIdLike == nil {
			return db
		} else {
			return db.Where(gorm.Expr(`lower(users.slack_id) LIKE ?`, strings.ToLower("%"+*slackIdLike+"%")))
		}
	}
}

func (r *UserRepository) stateEq(stateEq *string) func(db *gorm.DB) *gorm.DB {
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
	if user.Avatar != nil {
		err := r.db.Transaction(func(tx *gorm.DB) error {
			if err := r.db.Model(&models.User{}).Unscoped().Where("name = 'avatar'").Association("Avatar").Unscoped().Clear(); err != nil {
				return err
			}

			return r.db.Model(&user).Select(fields).Session(&gorm.Session{FullSaveAssociations: true}).Updates(&user).Error
		})

		return err
	}

	return r.db.Model(&user).Select(fields).Updates(&user).Error
}

func (r *UserRepository) All(users *[]*models.User) error {
	return r.db.Table("users").Order("id ASC").Find(&users).Error
}

func (r *UserRepository) UpdateProfile(user *models.User, updates map[string]interface{}) error {
	if err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&models.User{Id: user.Id}).Unscoped().Where("name = 'avatar'").Association("Avatar").Unscoped().Clear(); err != nil {
			return err
		}

		if err := tx.Model(&user).Select(append(helpers.GetKeys(updates), "Avatar")).Updates(updates).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return r.db.
		Model(&models.User{}).
		Where("id = ?", &user.Id).
		Preload("Avatar.AttachmentBlob").
		Preload("ProjectAssignees.User").
		Preload("ProjectAssignees.Project").
		First(&user).Error
}

// func (ur *UserRepository) Create(user *models.User, attributes map[string]interface{}) error {
// 	if err := ur.db.Transaction(func(tx *gorm.DB) error {
// 		ur.getDefaultData(attributes)

// 		if err := tx.Model(&user).
// 			Session(&gorm.Session{FullSaveAssociations: true}).
// 			Select(helpers.GetKeys(attributes)).Create(attributes).Error; err != nil {
// 			return err
// 		}

// 		attachment := attributes["Avatar"]

// 		if attachment != nil {
// 			avatar := models.Attachment{
// 				OwnerID: user.Id,
// 			}

// 			if err := tx.Model(&avatar).Create(attachment).Error; err != nil {
// 				return err
// 			}
// 		}
// 		return nil
// 	}); err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (ur *UserRepository) getDefaultData(attributes map[string]interface{}) {
// 	re := regexp.MustCompile(`(.*)@`)

// 	email := attributes["Email"].(string)
// 	matches := re.FindStringSubmatch(email)

// 	if len(matches) >= 2 {
// 		attributes["Name"] = matches[1]
// 	}

// 	timing := models.UserTiming{
// 		ActiveAt: time.Now().Format(constants.DateTimeZoneFormat),
// 	}

// 	attributes["Timing"] = &timing
// 	attributes["CreatedAt"] = time.Now().Format(constants.DataTimeMilisFormat)
// 	attributes["UpdatedAt"] = time.Now().Format(constants.DataTimeMilisFormat)
// 	attributes["LockVersion"] = 1
// }

//

func (ur *UserRepository) Create(user *models.User) error {
	return ur.db.Model(&user).Create(&user).Error
	// if err := ur.db.Transaction(func(tx *gorm.DB) error {
	// 	if err := tx.Model(&user).Create(&user).Error; err != nil {
	// 		return err
	// 	}
	// 	return nil
	// }); err != nil {
	// 	return err
	// }

	// return nil
}

func (r *UserRepository) UpdateUser(user *models.User, attributes map[string]interface{}) error {
	if err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&models.User{Id: user.Id}).Unscoped().Where("name = 'avatar'").Association("Avatar").Unscoped().Clear(); err != nil {
			return err
		}

		if err := tx.Model(&user).Select(append(helpers.GetKeys(attributes), "Avatar")).Updates(attributes).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return r.db.
		Model(&models.User{}).
		Where("id = ?", &user.Id).
		Preload("Avatar.AttachmentBlob").
		Preload("ProjectAssignees.User").
		Preload("ProjectAssignees.Project").
		First(&user).Error
}
