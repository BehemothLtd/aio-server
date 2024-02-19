package inputs

// MsSnippetInput represents input for creating or updating a snippet.
type MsSnippetInput struct {
	Title       *string
	Content     *string
	SnippetType *int32 // Use int32 or float32 as required by Graphql
	Passkey     *string
}
