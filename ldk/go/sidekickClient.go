package ldk

import "github.com/open-olive/loop-development-kit/ldk/go/v2/service"

// SidekickClient is used by Loops to interact with the host system through Sidekick
type SidekickClient struct {
	ClipboardService  service.ClipboardService
	VaultService      service.VaultService
	WhisperService    service.WhisperService
	KeyboardService   service.KeyboardService
	ProcessService    service.ProcessService
	CursorService     service.CursorService
	FilesystemService service.FilesystemService
	WindowService     service.WindowService
	UiService         service.UIService
	NetworkService    service.NetworkService
}

// Clipboard is used interact with the ClipboardService
func (m *SidekickClient) Clipboard() service.ClipboardService {
	return m.ClipboardService
}

// The Vault is used to store and retrieve sensitive data
func (m *SidekickClient) Vault() service.VaultService {
	return m.VaultService
}

// Whisper is used to send whispers to Sidekick
func (m *SidekickClient) Whisper() service.WhisperService {
	return m.WhisperService
}

// Keyboard is used to listen for KeyboardService events like keystrokes and hot-keys
func (m *SidekickClient) Keyboard() service.KeyboardService {
	return m.KeyboardService
}

// Process is used to list processes and listen for new user processes
func (m *SidekickClient) Process() service.ProcessService {
	return m.ProcessService
}

// Cursor is used to listen for the CursorService position and related events
func (m *SidekickClient) Cursor() service.CursorService {
	return m.CursorService
}

// Filesystem is used to interact with the host system's FilesystemService
func (m *SidekickClient) Filesystem() service.FilesystemService {
	return m.FilesystemService
}

// UI is used to listen for user interface events occurring in Sidekick, like search bar entry
func (m *SidekickClient) UI() service.UIService {
	return m.UiService
}

// Network is used to send/receive HTTP requests
func (m *SidekickClient) Network() service.NetworkService {
	return m.NetworkService
}

// Window is used to listen for the active WindowService, and other WindowService related events
func (m *SidekickClient) Window() service.WindowService {
	return m.WindowService
}
