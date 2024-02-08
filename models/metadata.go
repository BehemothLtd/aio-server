package models

type Metadata struct {
	Total   int64
	PerPage int
	Page    int
	Pages   int
	Count   int
	Next    int
	Prev    int
	From    int
	To      int
}
