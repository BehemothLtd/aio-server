//go:generate go-enum --marshal
package enums

/*
ENUM(
pending = 1
approved = 2
rejected = 3
personal_days_off = 4
)
*/
type RequestStateType string
