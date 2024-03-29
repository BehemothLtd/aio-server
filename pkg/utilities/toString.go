package utilities

import "fmt"

func IntToString(num int) string {
	return fmt.Sprintf("%d", num)
}

func BoolToString(bool bool) string {
	return fmt.Sprintf("%t", bool)
}
