package validators

import (
	"aio-server/exceptions"
	"aio-server/pkg/specialTypes"
	"slices"
)

// Form represents a form with attributes and errors.
type Form struct {
	Attributes []FieldAttributeInterface
	Errors     exceptions.ResourceModificationError
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
func (form *Form) summaryErrors() {
	err := exceptions.NewUnprocessableContentError("", make(map[string]*specialTypes.FieldAttributeErrorType))

	for _, attribute := range form.Attributes {
		attributeErr := attribute.GetErrors()
		attributeCode := attribute.GetCode()

		if len(attributeErr.Base) > 0 || attributeErr.Items != nil {
			if err.Errors[attributeCode] == nil {
				err.Errors[attributeCode] = &specialTypes.FieldAttributeErrorType{}

				err.Errors[attributeCode].Base = attributeErr.Base
				err.Errors[attributeCode].Items = attributeErr.Items
			}
		}
	}

	form.Errors = err.Errors
}
