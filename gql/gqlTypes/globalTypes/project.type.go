package globalTypes

import (
	"aio-server/models"
	"aio-server/pkg/helpers"
	"context"

	"github.com/graph-gophers/graphql-go"
	"gorm.io/gorm"
)

type ProjectType struct {
	Ctx *context.Context
	Db  *gorm.DB

	Project *models.Project
}

func (pt *ProjectType) ID(ctx context.Context) *graphql.ID {
	return helpers.GqlIDP(pt.Project.Id)
}

func (tt *ProjectType) Name(ctx context.Context) *string {
	return &tt.Project.Name
}

func (tt *ProjectType) Code(ctx context.Context) *string {
	return &tt.Project.Code
}
