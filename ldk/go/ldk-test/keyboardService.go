package ldktest

import (
	"context"
	"github.com/open-olive/loop-development-kit/ldk/go/v2/service"
)

type KeyboardService struct {
	ListenHotkeyf    func(context.Context, service.Hotkey, service.ListenHotkeyHandler) error
	ListenTextf      func(context.Context, service.ListenTextHandler) error
	ListenCharacterf func(context.Context, service.ListenCharacterHandler) error
}

func (k *KeyboardService) ListenHotkey(ctx context.Context, hotkey service.Hotkey, handler service.ListenHotkeyHandler) error {
	return k.ListenHotkeyf(ctx, hotkey, handler)
}

func (k *KeyboardService) ListenText(ctx context.Context, handler service.ListenTextHandler) error {
	return k.ListenTextf(ctx, handler)
}

func (k *KeyboardService) ListenCharacter(ctx context.Context, handler service.ListenCharacterHandler) error {
	return k.ListenCharacterf(ctx, handler)
}
