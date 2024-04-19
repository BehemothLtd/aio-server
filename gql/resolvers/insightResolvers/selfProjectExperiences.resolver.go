package insightResolvers

import (
	"aio-server/enums"
	"aio-server/exceptions"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/gqlTypes/insightTypes"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/repository"
	"context"
)

func (r *Resolver) SelfProjectExperiences(ctx context.Context, args insightInputs.ProjectExperiencesInput) (*insightTypes.ProjectExperiencesType, error) {
	user, err := r.Authorize(ctx, enums.PermissionTargetTypeProjectExperiences.String(), enums.PermissionActionTypeRead.String())
	if err != nil {
		return nil, exceptions.NewUnauthorizedError("")
	}

	var projectExperiences []*models.ProjectExperience
	query, paginationData := args.ToPaginationAndQueryData()

	repo := repository.NewProjectExperienceRepository(&ctx, r.Db)
	if err := repo.ListByUser(&projectExperiences, &paginationData, query, *user); err != nil {
		return nil, err
	}

	return &insightTypes.ProjectExperiencesType{
		Collection: r.ProjectExperiencesSliceToTypes(projectExperiences),
		Metadata: &globalTypes.MetadataType{
			Metadata: &paginationData.Metadata,
		},
	}, nil

}
