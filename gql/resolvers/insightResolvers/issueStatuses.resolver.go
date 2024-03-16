package insightResolvers

import (
	"aio-server/enums"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/gqlTypes/insightTypes"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/repository"
	"context"
)

func (r *Resolver) IssueStatuses(ctx context.Context, args insightInputs.IssueStatusesInput) (*insightTypes.IssueStatusesType, error) {
	if _, err := r.Authorize(ctx, enums.PermissionTargetTypeIssueStatuses.String(), enums.PermissionActionTypeRead.String()); err != nil {
		return nil, err
	}

	var issueStatuses []*models.IssueStatus
	IssueStatusQuery, paginationData := args.ToPaginationDataAndQuery()

	repo := repository.NewIssueStatusRepository(&ctx, r.Db)

	err := repo.List(&issueStatuses, &paginationData, IssueStatusQuery)
	if err != nil {
		return nil, err
	}

	return &insightTypes.IssueStatusesType{
		Collection: r.IssueStatusSliceToTypes(issueStatuses),
		Metadata: &globalTypes.MetadataType{
			Metadata: &paginationData.Metadata,
		},
	}, nil
}
