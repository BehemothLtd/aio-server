package insightResolvers

import (
	"aio-server/exceptions"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"context"

	"gorm.io/gorm"
)

func (r *Resolver) IssueStatus(ctx context.Context, args insightInputs.IssueStatusInput) (*globalTypes.IssueStatusType, error) {
	if args.Id == "" {
		return nil, exceptions.NewBadRequestError("Invalid Id")
	}

	issueStatusId, err := helpers.GqlIdToInt32(args.Id)

	if err != nil {
		return nil, err
	}

	issueStatus := models.IssueStatus{Id: issueStatusId}
	repo := repository.NewIssueStatusRepository(&ctx, r.Db)
	err = repo.Find(&issueStatus)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exceptions.NewRecordNotFoundError()
		}
		return nil, err
	}

	return &globalTypes.IssueStatusType{IssueStatus: &issueStatus}, nil
}
