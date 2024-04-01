package utilities

import (
	"fmt"
	"strings"
)

func IntToString(num int) string {
	return fmt.Sprintf("%d", num)
}

func BoolToString(bool bool) string {
	return fmt.Sprintf("%t", bool)
}

func SnakeCaseToHumanize(s string) string {
	words := strings.Split(s, "_")

	return strings.Join(words, " ")
}
