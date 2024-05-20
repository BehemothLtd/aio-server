//go:generate go-enum --marshal
package enums

/*
ENUM(
task = 1
bug = 2
)
*/
type IssueType string

var SqlIntIssueTypeValue = map[IssueType]int64{
	IssueTypeTask: 1,
	IssueTypeBug:  2,
}
