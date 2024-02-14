package validators

import (
	"aio-server/exceptions"
	"aio-server/repository"
	"slices"
)

type Form struct {
	Attributes []FieldAttributeInterface
	Errors     []exceptions.ResourceModifyErrors
	Repo       *repository.Repository
}

func (form *Form) AddAttributes(attributes ...FieldAttributeInterface) {
	form.Attributes = append(form.Attributes, attributes...)
}

func (form *Form) FindAttrByCode(attributeCode string) FieldAttributeInterface {
	idx := slices.IndexFunc(form.Attributes, func(a FieldAttributeInterface) bool { return a.GetCode() == attributeCode })

	if idx != -1 {
		return form.Attributes[idx]
	} else {
		return nil
	}
}

func (form *Form) SummaryErrors() {
	for _, attribute := range form.Attributes {
		attributeErr := attribute.GetErrors()
		attributeErrors := attributeErr.Errors

		if len(attributeErrors) > 0 {
			form.Errors = append(form.Errors, attributeErr)
		}
	}
}
