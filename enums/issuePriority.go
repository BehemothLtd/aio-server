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

var SqlIntIssuePriorityValue = map[IssuePriority]int64{
	IssuePriorityLowest:  1,
	IssuePriorityLow:     2,
	IssuePriorityNormal:  3,
	IssuePriorityHigh:    4,
	IssuePriorityHighest: 5,
}
