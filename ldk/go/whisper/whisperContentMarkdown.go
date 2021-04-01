package whisper

type WhisperContentMarkdown struct {
	Label string `json:"label"`

	Markdown string `json:"markdown"`
}

// Type returns the type of the WhisperService content
func (c *WhisperContentMarkdown) Type() WhisperContentType {
	return WhisperContentTypeMarkdown
}
