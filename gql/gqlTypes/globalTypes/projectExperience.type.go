package globalTypes

import (
	"aio-server/models"
	"aio-server/pkg/helpers"
	"context"

	graphql "github.com/graph-gophers/graphql-go"
	"gorm.io/gorm"
)

type ProjectExperienceType struct {
	Ctx *context.Context
	Db  *gorm.DB

	ProjectExperience *models.ProjectExperience
}

func (pet *ProjectExperienceType) ID(ctx context.Context) *graphql.ID {
	return helpers.GqlIDP(pet.ProjectExperience.Id)
}

func (pet *ProjectExperienceType) Title(ctx context.Context) *string {
	return &pet.ProjectExperience.Title
}

func (pet *ProjectExperienceType) ProjectId(ctx context.Context) *graphql.ID {
	return helpers.GqlIDP(pet.ProjectExperience.ProjectId)
}

func (pet *ProjectExperienceType) UserId(ctx context.Context) *graphql.ID {
	return helpers.GqlIDP(pet.ProjectExperience.UserId)
}

func (pet *ProjectExperienceType) Description(ctx context.Context) *string {
	return &pet.ProjectExperience.Description
}

func (pet *ProjectExperienceType) CreatedAt(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(&pet.ProjectExperience.CreatedAt)
}

func (pet *ProjectExperienceType) UpdatedAt(ctx context.Context) *graphql.Time {
	return helpers.GqlTimePointer(&pet.ProjectExperience.UpdatedAt)
}
