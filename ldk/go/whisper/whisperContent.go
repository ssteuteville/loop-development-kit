package whisper

// WhisperContentType defines the type of WhisperService content
type WhisperContentType string

const (
	// WhisperContentTypeConfirm is the content type used by a confirm WhisperService
	WhisperContentTypeConfirm WhisperContentType = "confirm"

	// WhisperContentTypeDisambiguation is the content type used by a disambiguation WhisperService
	WhisperContentTypeDisambiguation WhisperContentType = "disambiguation"

	// WhisperContentTypeForm is the content type used by a form WhisperService
	WhisperContentTypeForm WhisperContentType = "form"

	// WhisperContentTypeMarkdown is the content type used by a markdown WhisperService
	WhisperContentTypeMarkdown WhisperContentType = "markdown"

	// WhisperContentTypeList is the content type used by a list WhisperService
	WhisperContentTypeList WhisperContentType = "list"
)

// WhisperContent is an interface for the different types of WhisperService content
type WhisperContent interface {
	Type() WhisperContentType
}
