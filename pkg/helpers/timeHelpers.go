package helpers

import "time"

// Get time Beginning Of Day
func BeginningOfDay(t *time.Time) time.Time {
	if t == nil {
		now := time.Now()
		t = &now
	}

	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func EndOfDay(t *time.Time) time.Time {
	if t == nil {
		now := time.Now()
		t = &now
	}

	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, t.Location())
}
