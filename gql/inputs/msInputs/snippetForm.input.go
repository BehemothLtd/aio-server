package msInputs

// SnippetFormInput represents input for creating or updating a snippet.
type SnippetFormInput struct {
	Title       *string
	Content     *string
	SnippetType *int32 // Use int32 or float32 as required by Graphql
	Passkey     *string
	LockVersion *int32
}
