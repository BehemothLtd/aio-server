package insightResolvers

import (
	"aio-server/enums"
	"aio-server/exceptions"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/gqlTypes/insightTypes"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/auths"
	"aio-server/repository"
	"context"
)

// Projects resolves the query for retrieving a collection of Projects.
func (r *Resolver) ProjectSliceToTypes(projects []*models.Project) *[]*globalTypes.ProjectType {
	resolvers := make([]*globalTypes.ProjectType, len(projects))
	for i, s := range projects {
		resolvers[i] = &globalTypes.ProjectType{Project: s}
	}
	return &resolvers
}

func (r *Resolver) Projects(ctx context.Context, args insightInputs.ProjectInput) (*insightTypes.ProjectsType, error) {

	if _, err := r.Authorize(ctx, enums.PermissionTargetTypeProjects.String(), enums.PermissionActionTypeRead.String()); err != nil {
		return nil, err
	}
	user, err := auths.AuthUserFromCtx(ctx)

	if err != nil {
		return nil, exceptions.NewUnauthorizedError("")
	}

	var projects []*models.Project
	projectQuery, paginationData := args.ToPaginationDataAndProjectQuery()

	repo := repository.NewProjectRepository(&ctx, r.Db)

	errList := repo.List(&projects, &paginationData, projectQuery, &user)

	if errList != nil {
		return nil, errList
	}
	return &insightTypes.ProjectsType{
		Collection: r.ProjectSliceToTypes(projects),
		Metadata: &globalTypes.MetadataType{
			Metadata: &paginationData.Metadata,
		},
	}, nil
}
