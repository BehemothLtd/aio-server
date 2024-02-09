package models

type DataModificationError struct {
	Message string
	Errors  []*ResourceModifyErrors
}

type ResourceModifyErrors struct {
	Column string
	Errors []string
}
