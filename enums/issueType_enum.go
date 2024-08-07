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
	// IssueTypeTask is a IssueType of type task.
	IssueTypeTask IssueType = "task"
	// IssueTypeBug is a IssueType of type bug.
	IssueTypeBug IssueType = "bug"
)

var ErrInvalidIssueType = errors.New("not a valid IssueType")

// String implements the Stringer interface.
func (x IssueType) String() string {
	return string(x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x IssueType) IsValid() bool {
	_, err := ParseIssueType(string(x))
	return err == nil
}

var _IssueTypeValue = map[string]IssueType{
	"task": IssueTypeTask,
	"bug":  IssueTypeBug,
}

// ParseIssueType attempts to convert a string to a IssueType.
func ParseIssueType(name string) (IssueType, error) {
	if x, ok := _IssueTypeValue[name]; ok {
		return x, nil
	}
	return IssueType(""), fmt.Errorf("%s is %w", name, ErrInvalidIssueType)
}

// MarshalText implements the text marshaller method.
func (x IssueType) MarshalText() ([]byte, error) {
	return []byte(string(x)), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *IssueType) UnmarshalText(text []byte) error {
	tmp, err := ParseIssueType(string(text))
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

var errIssueTypeNilPtr = errors.New("value pointer is nil") // one per type for package clashes

var sqlIntIssueTypeMap = map[int64]IssueType{
	1: IssueTypeTask,
	2: IssueTypeBug,
}

var sqlIntIssueTypeValue = map[IssueType]int64{
	IssueTypeTask: 1,
	IssueTypeBug:  2,
}

func lookupSqlIntIssueType(val int64) (IssueType, error) {
	x, ok := sqlIntIssueTypeMap[val]
	if !ok {
		return x, fmt.Errorf("%v is not %w", val, ErrInvalidIssueType)
	}
	return x, nil
}

// Scan implements the Scanner interface.
func (x *IssueType) Scan(value interface{}) (err error) {
	if value == nil {
		*x = IssueType("")
		return
	}

	// A wider range of scannable types.
	// driver.Value values at the top of the list for expediency
	switch v := value.(type) {
	case int64:
		*x, err = lookupSqlIntIssueType(v)
	case string:
		*x, err = ParseIssueType(v)
	case []byte:
		if val, verr := strconv.ParseInt(string(v), 10, 64); verr == nil {
			*x, err = lookupSqlIntIssueType(val)
		} else {
			// try parsing the value as a string
			*x, err = ParseIssueType(string(v))
		}
	case IssueType:
		*x = v
	case int:
		*x, err = lookupSqlIntIssueType(int64(v))
	case *IssueType:
		if v == nil {
			return errIssueTypeNilPtr
		}
		*x = *v
	case uint:
		*x, err = lookupSqlIntIssueType(int64(v))
	case uint64:
		*x, err = lookupSqlIntIssueType(int64(v))
	case *int:
		if v == nil {
			return errIssueTypeNilPtr
		}
		*x, err = lookupSqlIntIssueType(int64(*v))
	case *int64:
		if v == nil {
			return errIssueTypeNilPtr
		}
		*x, err = lookupSqlIntIssueType(int64(*v))
	case float64: // json marshals everything as a float64 if it's a number
		*x, err = lookupSqlIntIssueType(int64(v))
	case *float64: // json marshals everything as a float64 if it's a number
		if v == nil {
			return errIssueTypeNilPtr
		}
		*x, err = lookupSqlIntIssueType(int64(*v))
	case *uint:
		if v == nil {
			return errIssueTypeNilPtr
		}
		*x, err = lookupSqlIntIssueType(int64(*v))
	case *uint64:
		if v == nil {
			return errIssueTypeNilPtr
		}
		*x, err = lookupSqlIntIssueType(int64(*v))
	case *string:
		if v == nil {
			return errIssueTypeNilPtr
		}
		*x, err = ParseIssueType(*v)
	default:
		return errors.New("invalid type for IssueType")
	}

	return
}

// Value implements the driver Valuer interface.
func (x IssueType) Value() (driver.Value, error) {
	val, ok := sqlIntIssueTypeValue[x]
	if !ok {
		return nil, ErrInvalidIssueType
	}
	return int64(val), nil
}
