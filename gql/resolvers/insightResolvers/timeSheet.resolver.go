package insightResolvers

import (
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/gqlTypes/insightTypes"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/auths"
	"context"
	"fmt"
)

func (r *Resolver) TimeSheet(ctx context.Context, args insightInputs.TimeSheetInput) (*insightTypes.TimeSheet, error) {
	if _, err := auths.AuthInsightUserFromCtx(ctx); err != nil {
		return nil, err
	}

	var results []struct {
		ProjectID int32
		UserID    int32
		Minutes   int32
	}

	r.Db.Table("working_timelogs").
		Joins("JOIN projects ON working_timelogs.project_id = projects.id").
		Joins("JOIN users ON working_timelogs.user_id = users.id").
		Select("projects.id as project_id, users.id as user_id, sum(working_timelogs.minutes) as minutes").
		Group("projects.id, users.id").
		Scan(&results)

	var projectData []*globalTypes.ProjectType
	var userData []*globalTypes.UserType

	for _, result := range results {
		project := &models.Project{}
		if err := r.Db.First(project, result.ProjectID).Error; err != nil {
			return nil, err
		}
		projectData = append(projectData, &globalTypes.ProjectType{Project: project})

		user := &models.User{}
		if err := r.Db.First(user, result.UserID).Error; err != nil {
			return nil, err
		}
		userData = append(userData, &globalTypes.UserType{User: user})
		fmt.Println(userData)
	}

	return &insightTypes.TimeSheet{
		ProjectData: projectData[0],
		UserData:    userData,
	}, nil
}
