package utilities

import (
	"fmt"
	"time"
)

func UnixTimestamp(t time.Time) string {
	second := t.Unix()
	nanos := t.UnixNano() % 1e9
	millis := nanos / 1e6

	timestamp := fmt.Sprintf("%d.%06d", second, millis)

	return timestamp
}
