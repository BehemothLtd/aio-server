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

func StartAndEndOfWeek(t time.Time) (startOfWeek, endOfWeek time.Time) {
	// Start of week (Monday 00:00:00)
	// Assuming week starts on Monday, subtract the weekday number from the current date.
	// Go's time.Weekday starts from Sunday as 0.
	offset := int(t.Weekday()) - 1 // Monday as the first day of the week
	if offset < 0 {
		offset = 6 // If today is Sunday, make offset point to last Monday
	}
	startOfWeek = t.AddDate(0, 0, -offset)
	startOfWeek = time.Date(startOfWeek.Year(), startOfWeek.Month(), startOfWeek.Day(), 0, 0, 0, 0, startOfWeek.Location())

	// End of week (Sunday 23:59:59)
	endOfWeek = startOfWeek.AddDate(0, 0, 6)
	endOfWeek = time.Date(endOfWeek.Year(), endOfWeek.Month(), endOfWeek.Day(), 23, 59, 59, 0, endOfWeek.Location())

	return startOfWeek, endOfWeek
}
