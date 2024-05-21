package insightInputs

import "github.com/graph-gophers/graphql-go"

type SelectOptionsInput struct {
	Input  SelectOptionsInputType
	Params *SelectOptionsParamsType
}

type SelectOptionsInputType struct {
	Keys *[]string
}

type SelectOptionsParamsType struct {
	ProjectId *graphql.ID
}
