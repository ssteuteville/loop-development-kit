package ldktest

import (
	"github.com/open-olive/loop-development-kit/ldk/go/v2/service"
)

type Sidekick struct {
	ClipboardService  service.ClipboardService
	CursorService     service.CursorService
	FilesystemService service.FilesystemService
	KeyboardService   service.KeyboardService
	NetworkService    service.NetworkService
	ProcessService    service.ProcessService
	VaultService      service.VaultService
	UIService         service.UIService
	WhisperService    service.WhisperService
	WindowService     service.WindowService
}

func (s *Sidekick) Clipboard() service.ClipboardService   { return s.ClipboardService }
func (s *Sidekick) Cursor() service.CursorService         { return s.CursorService }
func (s *Sidekick) Filesystem() service.FilesystemService { return s.FilesystemService }
func (s *Sidekick) Keyboard() service.KeyboardService     { return s.KeyboardService }
func (s *Sidekick) Network() service.NetworkService       { return s.NetworkService }
func (s *Sidekick) Process() service.ProcessService { return s.ProcessService }
func (s *Sidekick) Vault() service.VaultService     { return s.VaultService }
func (s *Sidekick) UI() service.UIService           { return s.UIService }
func (s *Sidekick) Whisper() service.WhisperService { return s.WhisperService }
func (s *Sidekick) Window() service.WindowService         { return s.WindowService }
