package utilities

import (
	"crypto/subtle"
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

func SecureCompare(s1, s2 []byte) bool {
	return subtle.ConstantTimeCompare(s1, s2) == 1
}
