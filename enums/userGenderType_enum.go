// Code generated by go-enum DO NOT EDIT.
// Version: 0.6.0
// Revision: 919e61c0174b91303753ee3898569a01abb32c97
// Build Date: 2023-12-18T15:54:43Z
// Built By: goreleaser

package enums

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strconv"
)

const (
	// UserGenderTypeMale is a UserGenderType of type male.
	UserGenderTypeMale UserGenderType = "male"
	// UserGenderTypeFemale is a UserGenderType of type female.
	UserGenderTypeFemale UserGenderType = "female"
	// UserGenderTypeBisexuality is a UserGenderType of type bisexuality.
	UserGenderTypeBisexuality UserGenderType = "bisexuality"
)

var ErrInvalidUserGenderType = errors.New("not a valid UserGenderType")

// String implements the Stringer interface.
func (x UserGenderType) String() string {
	return string(x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x UserGenderType) IsValid() bool {
	_, err := ParseUserGenderType(string(x))
	return err == nil
}

var _UserGenderTypeValue = map[string]UserGenderType{
	"male":        UserGenderTypeMale,
	"female":      UserGenderTypeFemale,
	"bisexuality": UserGenderTypeBisexuality,
}

// ParseUserGenderType attempts to convert a string to a UserGenderType.
func ParseUserGenderType(name string) (UserGenderType, error) {
	if x, ok := _UserGenderTypeValue[name]; ok {
		return x, nil
	}
	return UserGenderType(""), fmt.Errorf("%s is %w", name, ErrInvalidUserGenderType)
}

// MarshalText implements the text marshaller method.
func (x UserGenderType) MarshalText() ([]byte, error) {
	return []byte(string(x)), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *UserGenderType) UnmarshalText(text []byte) error {
	tmp, err := ParseUserGenderType(string(text))
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

var errUserGenderTypeNilPtr = errors.New("value pointer is nil") // one per type for package clashes

var sqlIntUserGenderTypeMap = map[int64]UserGenderType{
	0:   UserGenderTypeMale,
	50:  UserGenderTypeFemale,
	100: UserGenderTypeBisexuality,
}

var sqlIntUserGenderTypeValue = map[UserGenderType]int64{
	UserGenderTypeMale:        0,
	UserGenderTypeFemale:      50,
	UserGenderTypeBisexuality: 100,
}

func lookupSqlIntUserGenderType(val int64) (UserGenderType, error) {
	x, ok := sqlIntUserGenderTypeMap[val]
	if !ok {
		return x, fmt.Errorf("%v is not %w", val, ErrInvalidUserGenderType)
	}
	return x, nil
}

// Scan implements the Scanner interface.
func (x *UserGenderType) Scan(value interface{}) (err error) {
	if value == nil {
		*x = UserGenderType("")
		return
	}

	// A wider range of scannable types.
	// driver.Value values at the top of the list for expediency
	switch v := value.(type) {
	case int64:
		*x, err = lookupSqlIntUserGenderType(v)
	case string:
		*x, err = ParseUserGenderType(v)
	case []byte:
		if val, verr := strconv.ParseInt(string(v), 10, 64); verr == nil {
			*x, err = lookupSqlIntUserGenderType(val)
		} else {
			// try parsing the value as a string
			*x, err = ParseUserGenderType(string(v))
		}
	case UserGenderType:
		*x = v
	case int:
		*x, err = lookupSqlIntUserGenderType(int64(v))
	case *UserGenderType:
		if v == nil {
			return errUserGenderTypeNilPtr
		}
		*x = *v
	case uint:
		*x, err = lookupSqlIntUserGenderType(int64(v))
	case uint64:
		*x, err = lookupSqlIntUserGenderType(int64(v))
	case *int:
		if v == nil {
			return errUserGenderTypeNilPtr
		}
		*x, err = lookupSqlIntUserGenderType(int64(*v))
	case *int64:
		if v == nil {
			return errUserGenderTypeNilPtr
		}
		*x, err = lookupSqlIntUserGenderType(int64(*v))
	case float64: // json marshals everything as a float64 if it's a number
		*x, err = lookupSqlIntUserGenderType(int64(v))
	case *float64: // json marshals everything as a float64 if it's a number
		if v == nil {
			return errUserGenderTypeNilPtr
		}
		*x, err = lookupSqlIntUserGenderType(int64(*v))
	case *uint:
		if v == nil {
			return errUserGenderTypeNilPtr
		}
		*x, err = lookupSqlIntUserGenderType(int64(*v))
	case *uint64:
		if v == nil {
			return errUserGenderTypeNilPtr
		}
		*x, err = lookupSqlIntUserGenderType(int64(*v))
	case *string:
		if v == nil {
			return errUserGenderTypeNilPtr
		}
		*x, err = ParseUserGenderType(*v)
	default:
		return errors.New("invalid type for UserGenderType")
	}

	return
}

// Value implements the driver Valuer interface.
func (x UserGenderType) Value() (driver.Value, error) {
	val, ok := sqlIntUserGenderTypeValue[x]
	if !ok {
		return nil, ErrInvalidUserGenderType
	}
	return int64(val), nil
}
