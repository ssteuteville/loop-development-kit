package loop

import (
	"context"
	"github.com/open-olive/loop-development-kit/ldk/go/v2/client"
	"github.com/open-olive/loop-development-kit/ldk/go/v2/utils"
	"github.com/open-olive/loop-development-kit/ldk/go/v2/whisper"

	ldk "github.com/open-olive/loop-development-kit/ldk/go/v2"
)

// Serve creates the new loop and tells the LDK to serve it
func Serve() error {
	logger := utils.NewLogger("example-search-searchbar")
	loop, err := NewLoop(logger)
	if err != nil {
		return err
	}
	ldk.ServeLoopPlugin(logger, loop)
	return nil
}

// Loop is a structure for generating SideKick whispers
type Loop struct {
	ctx    context.Context
	cancel context.CancelFunc

	sidekick client.Sidekick
	logger   *utils.Logger
}

// NewLoop returns a pointer to a loop
func NewLoop(logger *utils.Logger) (*Loop, error) {
	return &Loop{
		logger: logger,
	}, nil
}

// LoopStart is called by the host when the loop is started to provide access to the host process
func (l *Loop) LoopStart(sidekick client.Sidekick) error {
	l.logger.Info("starting example loop")
	l.ctx, l.cancel = context.WithCancel(context.Background())

	l.sidekick = sidekick

	return sidekick.UI().ListenSearchbar(l.ctx, func(text string, err error) {
		l.logger.Info("loop callback called")
		if err != nil {
			l.logger.Error("received error from callback", err)
			return
		}

		go func() {
			err = l.sidekick.Whisper().Markdown(l.ctx, &whisper.WhisperContentMarkdown{
				Label:    "Example Go Loop",
				Markdown: "Text from the searchbar: " + text,
			})
			if err != nil {
				l.logger.Error("failed to emit whisper", "error", err)
				return
			}
		}()
	})
}

// LoopStop is called by the host when the loop is stopped
func (l *Loop) LoopStop() error {
	l.logger.Info("LoopStop called")
	l.cancel()

	return nil
}
