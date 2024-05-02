package insightResolvers

import (
	"aio-server/exceptions"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/gqlTypes/insightTypes"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/auths"
	"aio-server/pkg/constants"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"context"
	"time"
)

func (r *Resolver) WorkingTimelogByAttributes(ctx context.Context, args insightInputs.WorkingTimeLogInputByAttr) (*insightTypes.WorkingtimelogByAttributeType, error) {
	_, err := auths.AuthInsightUserFromCtx(ctx)

	if err != nil {
		return nil, exceptions.NewUnauthorizedError("")
	}

	projectId, projectIdErr := helpers.GqlIdToInt32(*args.ProjectId)
	issueId, issueIdErr := helpers.GqlIdToInt32(*args.IssueId)
	loggedAt, loggedAtErr := time.ParseInLocation(constants.DDMMYYYY_DateFormat, *args.LoggedAt, time.Local)

	if projectIdErr != nil || issueIdErr != nil || loggedAtErr != nil {
		return nil, exceptions.NewBadRequestError("Invalid input")
	}

	workingTimelog := models.WorkingTimelog{ProjectId: projectId, IssueId: issueId, LoggedAt: loggedAt}
	repo := repository.NewWorkingTimelogRepository(&ctx, r.Db.Preload("User").Preload("Project").Preload("Issue"))
	findErr := repo.FindByAttr(&workingTimelog)
	var dataExist = findErr == nil

	return &insightTypes.WorkingtimelogByAttributeType{
		WorkingTimelog: &globalTypes.WorkingTimelogType{WorkingTimelog: &workingTimelog},
		DataExisted:    dataExist,
	}, nil
}
