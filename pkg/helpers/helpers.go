package helpers

import (
	"fmt"
	"strconv"
	"strings"
	"time"

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

// GetStringOrDefault returns the value of the string pointer if not nil, otherwise returns an empty string.
func GetStringOrDefault(str *string) string {
	if str == nil {
		return ""
	}
	return *str
}

// GetInt32OrDefault returns the value of the int32 pointer if not nil, otherwise returns 0.
func GetInt32OrDefault(num *int32) int32 {
	if num == nil {
		return 0
	}
	return *num
}

// Int32Pointer returns a pointer to the given int32 value.
func Int32Pointer(val int32) *int32 {
	return &val
}

// IDPointer returns a pointer to the graphql.ID value.
func IDPointer(id graphql.ID) *graphql.ID {
	return &id
}

// GqlTimePointer returns a pointer to the graphql.Time value.
func GqlTimePointer(val time.Time) *graphql.Time {
	time := graphql.Time{Time: val}

	return &time
}
