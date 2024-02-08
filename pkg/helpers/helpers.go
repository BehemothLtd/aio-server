package helpers

import (
	"fmt"
	"strconv"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/pkg/errors"
)

func GqlIdToInt32(i graphql.ID) (int32, error) {
	r, err := strconv.ParseInt(string(i), 10, 32)

	if err != nil {
		return 0, errors.Wrap(err, "GqlIDToUint")
	}

	return int32(r), nil
}

// func int32P(i uint) *int32 {
// 	r := int32(i)
// 	return &r
// }

// func boolP(b bool) *bool {
// 	return &b
// }

func GqlIDP(id int32) *graphql.ID {
	if id == 0 {
		return nil
	}

	r := graphql.ID(fmt.Sprint(id))
	return &r
}
