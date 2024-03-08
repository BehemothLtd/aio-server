//go:generate go-enum --marshal
package enums

/*
ENUM(
available = 1
using = 2
discharged = 3
fixing = 4
)
*/
type DeviceStateType string
