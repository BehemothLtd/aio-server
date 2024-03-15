package insightResolvers

import (
	"aio-server/exceptions"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/models"
	"aio-server/pkg/auths"
	"aio-server/pkg/systems"
	"context"
)

// fromSnippets converts models.Snippet slice to []*UserGroupType.
func (r *Resolver) UserGroupSliceToTypes(userGroups []*models.UserGroup) *[]*globalTypes.UserGroupType {
	resolvers := make([]*globalTypes.UserGroupType, len(userGroups))
	for i, s := range userGroups {
		resolvers[i] = &globalTypes.UserGroupType{UserGroup: s}
	}
	return &resolvers
}

func (r *Resolver) Authorize(ctx context.Context, target string, action string) (*models.User, error) {
	user, err := auths.AuthInsightUserFromCtx(ctx)

	if err != nil {
		return nil, exceptions.NewUnauthorizedError("")
	}

	if !systems.AuthUserToAction(ctx, *r.Db, user, target, action) {
		return nil, exceptions.NewUnauthorizedError("You dont have authorization for this action")
	}

	return nil, nil
}

func (r *Resolver) UsersSliceToTypes(users []*models.User) *[]*globalTypes.UserType {
	resolvers := make([]*globalTypes.UserType, len(users))

	for i, u := range users {
		resolvers[i] = &globalTypes.UserType{User: u}
	}

	return &resolvers
}

func (r *Resolver) DeviceTypesSlideToType(deviceTypes []*models.DeviceType) *[]*globalTypes.DeviceTypeType {
	resolvers := make([]*globalTypes.DeviceTypeType, len(deviceTypes))
	for i, d := range deviceTypes {
		resolvers[i] = &globalTypes.DeviceTypeType{DeviceType: d}
	}
	return &resolvers
}


func (r *Resolver) IssueStatusSliceToTypes(issueStatuses []*models.IssueStatus) *[]*globalTypes.IssueStatusType {
	resolvers := make([]*globalTypes.IssueStatusType, len(issueStatuses))
	for i, s := range issueStatuses {
		resolvers[i] = &globalTypes.IssueStatusType{IssueStatus: s}
	}
  
  return &resolvers
}

func (r *Resolver) LeaveDayRequestSliceToTypes(requests []*models.LeaveDayRequest) *[]*globalTypes.LeaveDayRequestType {
	resolvers := make([]*globalTypes.LeaveDayRequestType, len(requests))

	for i, rq := range requests {
		resolvers[i] = &globalTypes.LeaveDayRequestType{LeaveDayRequest: rq}
	}

	return &resolvers
}
