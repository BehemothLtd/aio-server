//go:generate go-enum --marshal
package enums

/*
ENUM(
lowest = 1
low = 2
normal = 3
high = 4
highest = 5
)
*/
type IssuePriority string
