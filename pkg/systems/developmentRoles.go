package systems

import "slices"

type DevelopmentRole struct {
	Id    int32
	Code  string
	Title string
}

var defaultDevelopmentRoles = []DevelopmentRole{
	{Id: 1, Code: "dev", Title: "Developer"},
	{Id: 1, Code: "tester", Title: "Developer"},
	{Id: 1, Code: "comtor", Title: "Developer"},
	{Id: 1, Code: "brs", Title: "BrSE"},
	{Id: 1, Code: "pm", Title: "PM"},
}

func GetDevelopmentRoles() []DevelopmentRole {
	return append([]DevelopmentRole(nil), defaultDevelopmentRoles...)
}

func FindDevelopmentRoleById(id int32) *DevelopmentRole {
	allDevelopmentRoles := GetDevelopmentRoles()

	if foundIdx := slices.IndexFunc(allDevelopmentRoles, func(p DevelopmentRole) bool { return p.Id == id }); foundIdx != -1 {
		return &allDevelopmentRoles[foundIdx]
	} else {
		return nil
	}
}
