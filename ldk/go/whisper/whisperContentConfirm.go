package whisper

type WhisperContentConfirm struct {
	Label string `json:"label"`

	Markdown string `json:"markdown"`

	RejectLabel  string `json:"rejectLabel"`
	ResolveLabel string `json:"resolveLabel"`
}

// Type returns the type of the WhisperService content
func (c *WhisperContentConfirm) Type() WhisperContentType {
	return WhisperContentTypeConfirm
}
