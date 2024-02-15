package repository

import (
	"context"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
	ctx  *context.Context
}
