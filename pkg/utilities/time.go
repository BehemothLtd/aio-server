package utilities

import (
	"fmt"
	"time"
)

// Unit timestamp with millis second
func UnixTimestamp(t time.Time) string {
	second := t.Unix()
	nanos := t.UnixNano() % 1e9
	millis := nanos / 1e6

	timestamp := fmt.Sprintf("%d.%06d", second, millis)

	return timestamp
}

func UnixTimestampSecond(t time.Time) string {
	second := t.Unix()

	timestamp := fmt.Sprintf("%d", second)

	return timestamp
}
