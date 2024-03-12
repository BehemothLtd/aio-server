package repository

import (
	"context"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type Repository struct {
	db  *gorm.DB
	ctx *context.Context
}

func (r *Repository) ofParent(key string, parentID int32) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(gorm.Expr(key+" = ?", parentID))
	}
}

func (r *Repository) stringLike(tableName, columnName string, stringLike *string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if stringLike == nil {
			return db
		} else {
			return db.Where(gorm.Expr(fmt.Sprintf(`lower(%s.%s) LIKE ?`, tableName, columnName), strings.ToLower("%"+*stringLike+"%")))
		}
	}
}
