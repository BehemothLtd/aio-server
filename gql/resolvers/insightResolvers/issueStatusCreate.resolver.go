package insightResolvers

import (
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/gqlTypes/insightTypes"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/services/insightServices"
	"context"
)

func (r *Resolver) IssueStatusCreate(ctx context.Context, args struct {
	Input insightInputs.IssueStatusFormInput
}) (*insightTypes.IssueStatusCreatedType, error) {
	// validate form
	// save
	// error handling

	issueStatus := models.IssueStatus{}
	service := insightServices.IssueStatusCreateService{
		Ctx:         &ctx,
		Db:          r.Db,
		Args:        args.Input,
		IssueStatus: &issueStatus,
	}

	if err := service.Execute(); err != nil {
		return nil, err
	} else {
		return &insightTypes.IssueStatusCreatedType{
			IssueStatus: &globalTypes.IssueStatusType{
				IssueStatus: &issueStatus,
			},
		}, nil
	}
}
