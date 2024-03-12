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
	// UserStateActive is a UserState of type active.
	UserStateActive UserState = "active"
	// UserStateInactive is a UserState of type inactive.
	UserStateInactive UserState = "inactive"
)

var ErrInvalidUserState = errors.New("not a valid UserState")

// String implements the Stringer interface.
func (x UserState) String() string {
	return string(x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x UserState) IsValid() bool {
	_, err := ParseUserState(string(x))
	return err == nil
}

var _UserStateValue = map[string]UserState{
	"active":   UserStateActive,
	"inactive": UserStateInactive,
}

// ParseUserState attempts to convert a string to a UserState.
func ParseUserState(name string) (UserState, error) {
	if x, ok := _UserStateValue[name]; ok {
		return x, nil
	}
	return UserState(""), fmt.Errorf("%s is %w", name, ErrInvalidUserState)
}

// MarshalText implements the text marshaller method.
func (x UserState) MarshalText() ([]byte, error) {
	return []byte(string(x)), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *UserState) UnmarshalText(text []byte) error {
	tmp, err := ParseUserState(string(text))
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

var errUserStateNilPtr = errors.New("value pointer is nil") // one per type for package clashes

var sqlIntUserStateMap = map[int64]UserState{
	1: UserStateActive,
	2: UserStateInactive,
}

var sqlIntUserStateValue = map[UserState]int64{
	UserStateActive:   1,
	UserStateInactive: 2,
}

func lookupSqlIntUserState(val int64) (UserState, error) {
	x, ok := sqlIntUserStateMap[val]
	if !ok {
		return x, fmt.Errorf("%v is not %w", val, ErrInvalidUserState)
	}
	return x, nil
}

// Scan implements the Scanner interface.
func (x *UserState) Scan(value interface{}) (err error) {
	if value == nil {
		*x = UserState("")
		return
	}

	// A wider range of scannable types.
	// driver.Value values at the top of the list for expediency
	switch v := value.(type) {
	case int64:
		*x, err = lookupSqlIntUserState(v)
	case string:
		*x, err = ParseUserState(v)
	case []byte:
		if val, verr := strconv.ParseInt(string(v), 10, 64); verr == nil {
			*x, err = lookupSqlIntUserState(val)
		} else {
			// try parsing the value as a string
			*x, err = ParseUserState(string(v))
		}
	case UserState:
		*x = v
	case int:
		*x, err = lookupSqlIntUserState(int64(v))
	case *UserState:
		if v == nil {
			return errUserStateNilPtr
		}
		*x = *v
	case uint:
		*x, err = lookupSqlIntUserState(int64(v))
	case uint64:
		*x, err = lookupSqlIntUserState(int64(v))
	case *int:
		if v == nil {
			return errUserStateNilPtr
		}
		*x, err = lookupSqlIntUserState(int64(*v))
	case *int64:
		if v == nil {
			return errUserStateNilPtr
		}
		*x, err = lookupSqlIntUserState(int64(*v))
	case float64: // json marshals everything as a float64 if it's a number
		*x, err = lookupSqlIntUserState(int64(v))
	case *float64: // json marshals everything as a float64 if it's a number
		if v == nil {
			return errUserStateNilPtr
		}
		*x, err = lookupSqlIntUserState(int64(*v))
	case *uint:
		if v == nil {
			return errUserStateNilPtr
		}
		*x, err = lookupSqlIntUserState(int64(*v))
	case *uint64:
		if v == nil {
			return errUserStateNilPtr
		}
		*x, err = lookupSqlIntUserState(int64(*v))
	case *string:
		if v == nil {
			return errUserStateNilPtr
		}
		*x, err = ParseUserState(*v)
	default:
		return errors.New("invalid type for UserState")
	}

	return
}

// Value implements the driver Valuer interface.
func (x UserState) Value() (driver.Value, error) {
	val, ok := sqlIntUserStateValue[x]
	if !ok {
		return nil, ErrInvalidUserState
	}
	return int64(val), nil
}
