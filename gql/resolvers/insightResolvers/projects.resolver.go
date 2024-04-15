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

func (r *Resolver) Projects(ctx context.Context, args insightInputs.ProjectsInput) (*insightTypes.ProjectsType, error) {
	if _, err := r.Authorize(ctx, enums.PermissionTargetTypeProjects.String(), enums.PermissionActionTypeRead.String()); err != nil {
		return nil, err
	}

	var projects []*models.Project

	projectsQuery, paginationData := args.ToPaginationDataAndQuery()

	repo := repository.NewProjectRepository(
		&ctx,
		r.Db.
			Preload("Logo", "name = 'logo'").Preload("Logo.AttachmentBlob").
			Preload("ProjectAssignees.User.Avatar.AttachmentBlob"),
	)

	err := repo.List(&projects, &paginationData, projectsQuery)

	if err != nil {
		return nil, exceptions.NewBadRequestError(err.Error())
	}

	return &insightTypes.ProjectsType{
		Collection: r.ProjectsSliceToTypes(projects),
		Metadata: &globalTypes.MetadataType{
			Metadata: &paginationData.Metadata,
		},
	}, nil
}
