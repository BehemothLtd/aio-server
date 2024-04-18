package insightResolvers

import (
	"aio-server/enums"
	"aio-server/gql/gqlTypes/insightTypes"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/repository"
	"context"
)

func (r *Resolver) ProjectExperiences(ctx context.Context, args insightInputs.ProjectExperiencesInput) (*insightTypes.ProjectExperiencesType, error) {
	if _, err := r.Authorize(ctx, enums.PermissionTargetTypeProjectExperiences.String(), enums.PermissionActionTypeRead.String()); err != nil {
		return nil, err
	}

	var projectExperiences []*models.ProjectExperience
	query, paginationData := args.ToPaginationAndQueryData()

	repo := repository.NewProjectExperienceRepository(&ctx, r.Db)
}
