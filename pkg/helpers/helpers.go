package helpers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/google/uuid"
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

type SignedInteger interface {
	int | int8 | int16 | int32 | int64
}

func GqlIDP[T SignedInteger](id T) *graphql.ID {
	if id == 0 {
		return nil
	}

	r := graphql.ID(fmt.Sprint(id))
	return &r
}

func NewUUID() string {
	newUUID := uuid.New().String()

	idSplitted := strings.Split(newUUID, "-")
	idJoined := strings.Join(idSplitted[:], "")

	return idJoined
}
