package helpers

import (
	"fmt"
	"reflect"
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

func GqlIDValue[T SignedInteger](id T) graphql.ID {
	return graphql.ID(fmt.Sprint(id))
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

// GetStringOrDefault returns the value of the bool pointer if not nil, otherwise returns false.
func GetBoolOrDefault(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}

// GetInt32OrDefault returns the value of the int32 pointer if not nil, otherwise returns 0.
func GetInt32OrDefault(num *int32) int32 {
	if num == nil {
		return 0
	}
	return *num
}

// GetIntOrDefault returns the value of the int pointer if not nil, otherwise returns 0.
func GetIntOrDefault(num *int) int {
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
func GqlTimePointer(val *time.Time) *graphql.Time {
	if val != nil {
		time := graphql.Time{Time: *val}

		return &time
	} else {
		return nil
	}

}

// RubyTimeParser returns time.Time from string generated in Ruby
func RubyTimeParser(timeString string) (*time.Time, error) {
	layout := "2006-01-02 15:04:05 -0700"

	// Parse the string to time.Time object
	if t, err := time.ParseInLocation(layout, timeString, time.Local); err != nil {
		fmt.Println("Error parsing time:", err)
		return nil, errors.New("invalid time string")
	} else {
		return &t, nil
	}
}

func RubyTimeStringToGqlTime(timeString string) *graphql.Time {
	if time, err := RubyTimeParser(timeString); err != nil {
		return nil
	} else {
		return GqlTimePointer(time)
	}
}

// Float64Pointer returns a pointer to the given float64 value.
func Float64Pointer(val float64) *float64 {
	return &val
}

// GetFloat64OrDefault returns the value of the float64 pointer if not nil, otherwise returns 0.
func GetFloat64OrDefault(num *float64) float64 {
	if num == nil {
		return 0.0
	}
	return *num
}

// pluck extracts the field with the given name from each struct in the slice.
// It returns a slice of the extracted values as []interface{}.
func Pluck(slice interface{}, fieldName string) ([]interface{}, error) {
	sliceVal := reflect.ValueOf(slice)
	if sliceVal.Kind() != reflect.Slice {
		return nil, fmt.Errorf("expected slice type, got %s", sliceVal.Kind())
	}

	var result []interface{}
	for i := 0; i < sliceVal.Len(); i++ {
		elem := sliceVal.Index(i)
		if elem.Kind() != reflect.Struct {
			return nil, fmt.Errorf("expected slice of structs, got slice of %s", elem.Kind())
		}
		field := elem.FieldByName(fieldName)
		if !field.IsValid() {
			return nil, fmt.Errorf("no such field: %s in element %d", fieldName, i)
		}
		result = append(result, field.Interface())
	}

	return result, nil
}

// GroupByProperty groups a slice of structs by a specific property.
func GroupByProperty[T any, K comparable](items []T, getProperty func(T) K) map[K][]T {
	grouped := make(map[K][]T)

	for _, item := range items {
		key := getProperty(item)
		grouped[key] = append(grouped[key], item)
	}

	return grouped
}
