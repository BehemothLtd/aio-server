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
	// PinParentTypeSnippet is a PinParentType of type snippet.
	PinParentTypeSnippet PinParentType = "snippet"
	// PinParentTypeProject is a PinParentType of type project.
	PinParentTypeProject PinParentType = "project"
)

var ErrInvalidPinParentType = errors.New("not a valid PinParentType")

// String implements the Stringer interface.
func (x PinParentType) String() string {
	return string(x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x PinParentType) IsValid() bool {
	_, err := ParsePinParentType(string(x))
	return err == nil
}

var _PinParentTypeValue = map[string]PinParentType{
	"snippet": PinParentTypeSnippet,
	"project": PinParentTypeProject,
}

// ParsePinParentType attempts to convert a string to a PinParentType.
func ParsePinParentType(name string) (PinParentType, error) {
	if x, ok := _PinParentTypeValue[name]; ok {
		return x, nil
	}
	return PinParentType(""), fmt.Errorf("%s is %w", name, ErrInvalidPinParentType)
}

// MarshalText implements the text marshaller method.
func (x PinParentType) MarshalText() ([]byte, error) {
	return []byte(string(x)), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *PinParentType) UnmarshalText(text []byte) error {
	tmp, err := ParsePinParentType(string(text))
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

var errPinParentTypeNilPtr = errors.New("value pointer is nil") // one per type for package clashes

var sqlIntPinParentTypeMap = map[int64]PinParentType{
	1: PinParentTypeSnippet,
	2: PinParentTypeProject,
}

var sqlIntPinParentTypeValue = map[PinParentType]int64{
	PinParentTypeSnippet: 1,
	PinParentTypeProject: 2,
}

func lookupSqlIntPinParentType(val int64) (PinParentType, error) {
	x, ok := sqlIntPinParentTypeMap[val]
	if !ok {
		return x, fmt.Errorf("%v is not %w", val, ErrInvalidPinParentType)
	}
	return x, nil
}

// Scan implements the Scanner interface.
func (x *PinParentType) Scan(value interface{}) (err error) {
	if value == nil {
		*x = PinParentType("")
		return
	}

	// A wider range of scannable types.
	// driver.Value values at the top of the list for expediency
	switch v := value.(type) {
	case int64:
		*x, err = lookupSqlIntPinParentType(v)
	case string:
		*x, err = ParsePinParentType(v)
	case []byte:
		if val, verr := strconv.ParseInt(string(v), 10, 64); verr == nil {
			*x, err = lookupSqlIntPinParentType(val)
		} else {
			// try parsing the value as a string
			*x, err = ParsePinParentType(string(v))
		}
	case PinParentType:
		*x = v
	case int:
		*x, err = lookupSqlIntPinParentType(int64(v))
	case *PinParentType:
		if v == nil {
			return errPinParentTypeNilPtr
		}
		*x = *v
	case uint:
		*x, err = lookupSqlIntPinParentType(int64(v))
	case uint64:
		*x, err = lookupSqlIntPinParentType(int64(v))
	case *int:
		if v == nil {
			return errPinParentTypeNilPtr
		}
		*x, err = lookupSqlIntPinParentType(int64(*v))
	case *int64:
		if v == nil {
			return errPinParentTypeNilPtr
		}
		*x, err = lookupSqlIntPinParentType(int64(*v))
	case float64: // json marshals everything as a float64 if it's a number
		*x, err = lookupSqlIntPinParentType(int64(v))
	case *float64: // json marshals everything as a float64 if it's a number
		if v == nil {
			return errPinParentTypeNilPtr
		}
		*x, err = lookupSqlIntPinParentType(int64(*v))
	case *uint:
		if v == nil {
			return errPinParentTypeNilPtr
		}
		*x, err = lookupSqlIntPinParentType(int64(*v))
	case *uint64:
		if v == nil {
			return errPinParentTypeNilPtr
		}
		*x, err = lookupSqlIntPinParentType(int64(*v))
	case *string:
		if v == nil {
			return errPinParentTypeNilPtr
		}
		*x, err = ParsePinParentType(*v)
	default:
		return errors.New("invalid type for PinParentType")
	}

	return
}

// Value implements the driver Valuer interface.
func (x PinParentType) Value() (driver.Value, error) {
	val, ok := sqlIntPinParentTypeValue[x]
	if !ok {
		return nil, ErrInvalidPinParentType
	}
	return int64(val), nil
}
