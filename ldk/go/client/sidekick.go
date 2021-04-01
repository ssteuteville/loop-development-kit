package client

import "github.com/open-olive/loop-development-kit/ldk/go/v2/service"

// Sidekick is an interface that defines what methods plugins can expect from the host
type Sidekick interface {
	Clipboard() service.ClipboardService
	Vault() service.VaultService
	Whisper() service.WhisperService
	Keyboard() service.KeyboardService
	Process() service.ProcessService
	Cursor() service.CursorService
	Filesystem() service.FilesystemService
	Window() service.WindowService
	UI() service.UIService
	Network() service.NetworkService
}
