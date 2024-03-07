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
	// ProjectPriorityLow is a ProjectPriority of type Low.
	ProjectPriorityLow ProjectPriority = "low"
	// ProjectPriorityMedium is a ProjectPriority of type medium.
	ProjectPriorityMedium ProjectPriority = "medium"
	// ProjectPriorityHigh is a ProjectPriority of type high.
	ProjectPriorityHigh ProjectPriority = "high"
	
)

var ErrInvalidProjectPriority = errors.New("not a valid ProjectPriority")

// String implements the Stringer interface.
func (x ProjectPriority) String() string {
	return string(x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x ProjectPriority) IsValid() bool {
	_, err := ParseProjectPriority(string(x))
	return err == nil
}

var _ProjectPriorityValue = map[string]ProjectPriority{
	"low":  ProjectPriorityLow,
	"medium": ProjectPriorityMedium,
	"high": ProjectPriorityHigh,
}

// ParseProjectPriority attempts to convert a string to a ParseProjectPriority.
func ParseProjectPriority(name string) (ProjectPriority, error) {
	if x, ok := _ProjectPriorityValue[name]; ok {
		return x, nil
	}
	return ProjectPriority(""), fmt.Errorf("%s is %w", name, ErrInvalidProjectPriority)
}

// MarshalText implements the text marshaller method.
func (x ProjectPriority) MarshalText() ([]byte, error) {
	return []byte(string(x)), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *ProjectPriority) UnmarshalText(text []byte) error {
	tmp, err := ParseProjectPriority(string(text))
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

var errProjectPriorityNilPtr = errors.New("value pointer is nil") // one per type for package clashes

var sqlIntProjectPriorityMap = map[int64]ProjectPriority{
	1: ProjectPriorityLow,
	2: ProjectPriorityMedium,
	3: ProjectPriorityHigh,
}

var sqlIntProjectPriorityValue = map[ProjectPriority]int64{
	ProjectPriorityLow:  1,
	ProjectPriorityMedium: 2,
	ProjectPriorityHigh: 3,
}

func lookupSqlIntProjectPriority(val int64) (ProjectPriority, error) {
	x, ok := sqlIntProjectPriorityMap[val]
	if !ok {
		return x, fmt.Errorf("%v is not %w", val, ErrInvalidProjectPriority)
	}
	return x, nil
}

// Scan implements the Scanner interface.
func (x *ProjectPriority) Scan(value interface{}) (err error) {
	if value == nil {
		*x = ProjectPriority("")
		return
	}

	// A wider range of scannable types.
	// driver.Value values at the top of the list for expediency
	switch v := value.(type) {
	case int64:
		*x, err = lookupSqlIntProjectPriority(v)
	case string:
		*x, err = ParseProjectPriority(v)
	case []byte:
		if val, verr := strconv.ParseInt(string(v), 10, 64); verr == nil {
			*x, err = lookupSqlIntProjectPriority(val)
		} else {
			// try parsing the value as a string
			*x, err = ParseProjectPriority(string(v))
		}
	case ProjectPriority:
		*x = v
	case int:
		*x, err = lookupSqlIntProjectPriority(int64(v))
	case *ProjectPriority:
		if v == nil {
			return errProjectPriorityNilPtr
		}
		*x = *v
	case uint:
		*x, err = lookupSqlIntProjectPriority(int64(v))
	case uint64:
		*x, err = lookupSqlIntProjectPriority(int64(v))
	case *int:
		if v == nil {
			return errProjectPriorityNilPtr
		}
		*x, err = lookupSqlIntProjectPriority(int64(*v))
	case *int64:
		if v == nil {
			return errProjectPriorityNilPtr
		}
		*x, err = lookupSqlIntProjectPriority(int64(*v))
	case float64: // json marshals everything as a float64 if it's a number
		*x, err = lookupSqlIntProjectPriority(int64(v))
	case *float64: // json marshals everything as a float64 if it's a number
		if v == nil {
			return errProjectPriorityNilPtr
		}
		*x, err = lookupSqlIntProjectPriority(int64(*v))
	case *uint:
		if v == nil {
			return errProjectPriorityNilPtr
		}
		*x, err = lookupSqlIntProjectPriority(int64(*v))
	case *uint64:
		if v == nil {
			return errProjectPriorityNilPtr
		}
		*x, err = lookupSqlIntProjectPriority(int64(*v))
	case *string:
		if v == nil {
			return errProjectPriorityNilPtr
		}
		*x, err = ParseProjectPriority(*v)
	default:
		return errors.New("invalid type for ProjectPriority")
	}

	return
}

// Value implements the driver Valuer interface.
func (x ProjectPriority) Value() (driver.Value, error) {
	val, ok := sqlIntProjectPriorityValue[x]
	if !ok {
		return nil, ErrInvalidProjectPriority
	}
	return int64(val), nil
}
