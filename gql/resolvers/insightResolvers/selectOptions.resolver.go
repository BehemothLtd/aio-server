package insightResolvers

import (
	"aio-server/exceptions"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/gqlTypes/insightTypes"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/pkg/auths"
	"aio-server/repository"
	"context"
)

func (r *Resolver) SelectOptions(ctx context.Context, args []insightInputs.SelectOptionInput) ([]insightTypes.SelectOptionsType, error) {
	if _, err := auths.AuthUserFromCtx(ctx); err != nil {
		return nil, exceptions.NewUnauthorizedError("")
	}

	repo := repository.NewSelectOptionRepository(&ctx, r.Db)
	var result []insightTypes.SelectOptionsType

	for _, item := range args {
		switch item.Key {
		case "users":
			users, err := repo.FetchUsers()
			if err != nil {
				return nil, err
			}
			var selectOptions []globalTypes.SelectOptionType
			for _, user := range users {
				selectOptions = append(selectOptions, globalTypes.SelectOptionType{
					SelectOption: &globalTypes.OptionType{
						Label: &user.Name,
						Value: user.Id,
					},
				})
			}

			result = append(result, insightTypes.SelectOptionsType{
				Users: selectOptions,
			})
		}
	}

	return result, nil
}
