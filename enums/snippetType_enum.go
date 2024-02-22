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
	// SnippetTypePublic is a SnippetType of type public.
	SnippetTypePublic SnippetType = "public"
	// SnippetTypePrivate is a SnippetType of type private.
	SnippetTypePrivate SnippetType = "private"
)

var ErrInvalidSnippetType = errors.New("not a valid SnippetType")

// String implements the Stringer interface.
func (x SnippetType) String() string {
	return string(x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x SnippetType) IsValid() bool {
	_, err := ParseSnippetType(string(x))
	return err == nil
}

var _SnippetTypeValue = map[string]SnippetType{
	"public":  SnippetTypePublic,
	"private": SnippetTypePrivate,
}

// ParseSnippetType attempts to convert a string to a SnippetType.
func ParseSnippetType(name string) (SnippetType, error) {
	if x, ok := _SnippetTypeValue[name]; ok {
		return x, nil
	}
	return SnippetType(""), fmt.Errorf("%s is %w", name, ErrInvalidSnippetType)
}

// MarshalText implements the text marshaller method.
func (x SnippetType) MarshalText() ([]byte, error) {
	return []byte(string(x)), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *SnippetType) UnmarshalText(text []byte) error {
	tmp, err := ParseSnippetType(string(text))
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

var errSnippetTypeNilPtr = errors.New("value pointer is nil") // one per type for package clashes

var sqlIntSnippetTypeMap = map[int64]SnippetType{
	1: SnippetTypePublic,
	2: SnippetTypePrivate,
}

var sqlIntSnippetTypeValue = map[SnippetType]int64{
	SnippetTypePublic:  1,
	SnippetTypePrivate: 2,
}

func lookupSqlIntSnippetType(val int64) (SnippetType, error) {
	x, ok := sqlIntSnippetTypeMap[val]
	if !ok {
		return x, fmt.Errorf("%v is not %w", val, ErrInvalidSnippetType)
	}
	return x, nil
}

// Scan implements the Scanner interface.
func (x *SnippetType) Scan(value interface{}) (err error) {
	if value == nil {
		*x = SnippetType("")
		return
	}

	// A wider range of scannable types.
	// driver.Value values at the top of the list for expediency
	switch v := value.(type) {
	case int64:
		*x, err = lookupSqlIntSnippetType(v)
	case string:
		*x, err = ParseSnippetType(v)
	case []byte:
		if val, verr := strconv.ParseInt(string(v), 10, 64); verr == nil {
			*x, err = lookupSqlIntSnippetType(val)
		} else {
			// try parsing the value as a string
			*x, err = ParseSnippetType(string(v))
		}
	case SnippetType:
		*x = v
	case int:
		*x, err = lookupSqlIntSnippetType(int64(v))
	case *SnippetType:
		if v == nil {
			return errSnippetTypeNilPtr
		}
		*x = *v
	case uint:
		*x, err = lookupSqlIntSnippetType(int64(v))
	case uint64:
		*x, err = lookupSqlIntSnippetType(int64(v))
	case *int:
		if v == nil {
			return errSnippetTypeNilPtr
		}
		*x, err = lookupSqlIntSnippetType(int64(*v))
	case *int64:
		if v == nil {
			return errSnippetTypeNilPtr
		}
		*x, err = lookupSqlIntSnippetType(int64(*v))
	case float64: // json marshals everything as a float64 if it's a number
		*x, err = lookupSqlIntSnippetType(int64(v))
	case *float64: // json marshals everything as a float64 if it's a number
		if v == nil {
			return errSnippetTypeNilPtr
		}
		*x, err = lookupSqlIntSnippetType(int64(*v))
	case *uint:
		if v == nil {
			return errSnippetTypeNilPtr
		}
		*x, err = lookupSqlIntSnippetType(int64(*v))
	case *uint64:
		if v == nil {
			return errSnippetTypeNilPtr
		}
		*x, err = lookupSqlIntSnippetType(int64(*v))
	case *string:
		if v == nil {
			return errSnippetTypeNilPtr
		}
		*x, err = ParseSnippetType(*v)
	default:
		return errors.New("invalid type for SnippetType")
	}

	return
}

// Value implements the driver Valuer interface.
func (x SnippetType) Value() (driver.Value, error) {
	val, ok := sqlIntSnippetTypeValue[x]
	if !ok {
		return nil, ErrInvalidSnippetType
	}
	return int64(val), nil
}