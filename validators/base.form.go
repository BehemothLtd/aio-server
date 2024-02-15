package validators

import (
	"aio-server/exceptions"
	"slices"
)

// Form represents a form with attributes and errors.
type Form struct {
	Attributes []FieldAttributeInterface
	Errors     []exceptions.ResourceModifyErrors
}

// AddAttributes adds attributes to the form.
func (form *Form) AddAttributes(attributes ...FieldAttributeInterface) {
	form.Attributes = append(form.Attributes, attributes...)
}

// FindAttrByCode finds an attribute by its code.
func (form *Form) FindAttrByCode(attributeCode string) FieldAttributeInterface {
	idx := slices.IndexFunc(form.Attributes, func(a FieldAttributeInterface) bool { return a.GetCode() == attributeCode })

	if idx != -1 {
		return form.Attributes[idx]
	} else {
		return nil
	}
}

// SummaryErrors summarizes errors in the form.
func (form *Form) SummaryErrors() {
	for _, attribute := range form.Attributes {
		attributeErr := attribute.GetErrors()
		if len(attributeErr.Errors) > 0 {
			form.Errors = append(form.Errors, attributeErr)
		}
	}
}
