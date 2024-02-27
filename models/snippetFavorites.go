package models

type SnippetsFavorite struct {
	Id        int32
	UserId    int32
	SnippetId int32

	Snippet Snippet
	User    User
}
