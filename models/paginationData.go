package models

type PaginationData struct {
	Metadata   Metadata
	Collection interface{}
}

type PagyInput struct {
	PerPage *int `json:"perPage,omitempty"`
	Page    *int `json:"page,omitempty"`
}
