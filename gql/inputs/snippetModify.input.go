package inputs

// For create/update
type MsSnippetInput struct {
	Title       *string
	Content     *string
	SnippetType *int32 // we have to use int32 OR float32 for input struct because Graphql ASKED FOR IT
}

func (msi *MsSnippetInput) GetTitle() string {
	if msi == nil {
		return ""
	}

	if msi.Title == nil {
		return ""
	} else {
		return *msi.Title
	}
}

func (msi *MsSnippetInput) GetContent() string {
	if msi == nil {
		return ""
	}

	if msi.Content == nil {
		return ""
	} else {
		return *msi.Content
	}
}

func (msi *MsSnippetInput) GetSnippetType() int32 {
	if msi == nil {
		return 0
	}

	if msi.SnippetType == nil {
		return 0
	} else {
		return *msi.SnippetType
	}
}
