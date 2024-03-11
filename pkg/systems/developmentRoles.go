package systems

type DevelopmentRole struct {
	Id    int
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
