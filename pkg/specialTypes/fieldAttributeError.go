package specialTypes

type FieldAttributeErrorType struct {
	Base  []string                    `json:"base"`
	Items map[int]map[string][]string `json:"items"`
}
