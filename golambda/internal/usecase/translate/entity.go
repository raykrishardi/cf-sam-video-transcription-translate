package translate

type TranslateDocumentInput struct {
	Content            []byte
	ContentType        string
	SourceLanguageCode *string
	TargetLanguageCode *string
}
