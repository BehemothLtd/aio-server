package insightInputs

import "github.com/graph-gophers/graphql-go"

type ProjectUploadImagesInput struct {
	Id    graphql.ID
	Input ProjectUploadImagesFormInput
}

type ProjectUploadImagesFormInput struct {
	LogoKey  *string
	FileKeys *[]string
}
