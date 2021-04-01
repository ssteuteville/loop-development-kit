package service

import (
	"context"
	"github.com/open-olive/loop-development-kit/ldk/go/v2/whisper"
)

// WhisperService is an interface that defines what methods plugins can expect from the host
type WhisperService interface {
	// This function only returns once the WhisperService is closed, or when the context provided is cancelled.
	Confirm(context.Context, *whisper.WhisperContentConfirm) (bool, error)
	// This function only returns once the WhisperService is closed, or when the context provided is cancelled.
	Disambiguation(context.Context, *whisper.WhisperContentDisambiguation) (bool, error)
	// This function only returns once the WhisperService is closed, or when the context provided is cancelled.
	Form(context.Context, *whisper.WhisperContentForm) (bool, map[string]whisper.WhisperContentFormOutput, error)
	// This function only returns once the WhisperService is closed, or when the context provided is cancelled.
	Markdown(context.Context, *whisper.WhisperContentMarkdown) error
	// This function only returns once the WhisperService is closed, or when the context provided is cancelled.
	List(context.Context, *whisper.WhisperContentList) error
}
