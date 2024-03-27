package insightResolvers

import (
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/gqlTypes/insightTypes"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"context"
	"time"
)

func (r *Resolver) ClientCreate(ctx context.Context, args insightInputs.ClientCreateInput) (*insightTypes.ClientCreatedType, error) {
	// if _, err := r.Authorize(ctx, enums.PermissionTargetTypeClients.String(), enums.PermissionActionTypeWrite.String()); err != nil {
	// 	return nil, err
	// }

	return &insightTypes.ClientCreatedType{
		Client: &globalTypes.ClientType{
			Client: &models.Client{
				Id:             1,
				Name:           "Example Client",
				ShowOnHomePage: true,
				LockVersion:    0,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
		},
	}, nil
	// service := insightServices.IssueStatusCreateService{
	// 	Ctx:         &ctx,
	// 	Db:          r.Db,
	// 	Args:        args,
	// 	IssueStatus: &IssueStatus,
	// }

	// if err := service.Execute(); err != nil {
	// 	return nil, err
	// } else {
	// 	return &insightTypes.IssueStatusCreatedType{
	// 		IssueStatus: &globalTypes.IssueStatusType{
	// 			IssueStatus: &IssueStatus,
	// 		},
	// 	}, nil
	// }
}
