package resolvers

import (
	"gorm.io/gorm"
)

type Resolver struct {
	Db *gorm.DB
}
