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

func (r *Resolver) Issues(ctx context.Context, args insightInputs.IssuesInput) (*insightTypes.IssuesType, error) {
	if _, err := r.Authorize(ctx, enums.PermissionTargetTypeProjectIssues.String(), enums.PermissionActionTypeRead.String()); err != nil {
		return nil, err
	}

	var issues []*models.Issue

	IssuesQuery, paginationData := args.ToPaginationDataAndQuery()

	repo := repository.NewIssueRepository(&ctx, r.Db)

	err := repo.List(&issues, &paginationData, IssuesQuery)

	if err != nil {
		return nil, exceptions.NewBadRequestError(err.Error())
	}

	return &insightTypes.IssuesType{
		Collection: r.IssueSliceToTypes(issues),
		Metadata: &globalTypes.MetadataType{
			Metadata: &paginationData.Metadata,
		},
	}, nil
}
