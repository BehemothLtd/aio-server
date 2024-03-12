package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type UserTiming struct {
	ActiveAt   string `json:"active_at"`
	InactiveAt string `json:"inactive_at"`
}

func (t UserTiming) Value() (driver.Value, error) {
	if byteArray, err := json.Marshal(t); err != nil {
		return nil, err
	} else {
		return string(byteArray), nil
	}
}

func (t *UserTiming) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSON value:", value))
	}

	result := UserTiming{}
	if err := json.Unmarshal(bytes, &result); err != nil {
		return err
	}

	*t = result
	return nil
}
