package msInputs

// SnippetFormInput represents input for creating or updating a snippet.
type SnippetFormInput struct {
	Title       *string
	Content     *string
	SnippetType *string
	Passkey     *string
	LockVersion *int32
}
