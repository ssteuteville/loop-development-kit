package ldktest

import (
	"context"
	"github.com/open-olive/loop-development-kit/ldk/go/v2/whisper"
)

type WhisperService struct {
	Confirmf        func(context.Context, *whisper.WhisperContentConfirm) (bool, error)
	Disambiguationf func(context.Context, *whisper.WhisperContentDisambiguation) (bool, error)
	Formf           func(context.Context, *whisper.WhisperContentForm) (bool, map[string]whisper.WhisperContentFormOutput, error)
	Listf           func(context.Context, *whisper.WhisperContentList) error
	Markdownf       func(context.Context, *whisper.WhisperContentMarkdown) error
}

func (w *WhisperService) Confirm(ctx context.Context, content *whisper.WhisperContentConfirm) (bool, error) {
	return w.Confirmf(ctx, content)
}

func (w *WhisperService) Disambiguation(ctx context.Context, content *whisper.WhisperContentDisambiguation) (bool, error) {
	return w.Disambiguationf(ctx, content)
}

func (w *WhisperService) Form(ctx context.Context, content *whisper.WhisperContentForm) (bool, map[string]whisper.WhisperContentFormOutput, error) {
	return w.Formf(ctx, content)
}

func (w *WhisperService) List(ctx context.Context, content *whisper.WhisperContentList) error {
	return w.Listf(ctx, content)
}

func (w *WhisperService) Markdown(ctx context.Context, content *whisper.WhisperContentMarkdown) error {
	return w.Markdownf(ctx, content)
}
