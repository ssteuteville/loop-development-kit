package ldk

import (
	"context"
	"errors"
	"github.com/open-olive/loop-development-kit/ldk/go/v2/server"
	"github.com/open-olive/loop-development-kit/ldk/go/v2/service"
	"io"

	"github.com/open-olive/loop-development-kit/ldk/go/v2/proto"
)

// KeyboardClient is used by the controller plugin to facilitate plugin initiated communication with the host
type KeyboardClient struct {
	KeyboardClient proto.KeyboardClient
	Session        *server.Session
}

// ListenHotkey will listen for a specific set of hotkeys and call the handler for each press up or down
func (k *KeyboardClient) ListenHotkey(ctx context.Context, hotkey service.Hotkey, handler service.ListenHotkeyHandler) error {
	key := &proto.KeyboardHotkey{Key: string(hotkey.Key), Modifiers: int32(hotkey.Modifiers)}

	stream, err := k.KeyboardClient.KeyboardHotkeyStream(ctx, &proto.KeyboardHotkeyStreamRequest{
		Hotkey:  key,
		Session: k.Session.ToProto(),
	})
	if err != nil {
		return err
	}

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				handler(false, err)
				return
			}

			if resp.GetError() != "" {
				err = errors.New(resp.GetError())
			}
			handler(resp.GetScanned(), err)
		}
	}()

	return nil
}

// ListenText will call the handler when a chunk of text has been captured from the keyboard
func (k *KeyboardClient) ListenText(ctx context.Context, handler service.ListenTextHandler) error {
	stream, err := k.KeyboardClient.KeyboardTextStream(ctx, &proto.KeyboardTextStreamRequest{
		Session: k.Session.ToProto(),
	})
	if err != nil {
		return err
	}

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				handler("", err)
				return
			}

			if resp.GetError() != "" {
				err = errors.New(resp.GetError())
			}
			handler(resp.GetText(), err)
		}
	}()

	return nil
}

// ListenCharacter will call the handler for each character typed on the keyboard
func (k *KeyboardClient) ListenCharacter(ctx context.Context, handler service.ListenCharacterHandler) error {
	stream, err := k.KeyboardClient.KeyboardCharacterStream(ctx, &proto.KeyboardCharacterStreamRequest{
		Session: k.Session.ToProto(),
	})
	if err != nil {
		return err
	}

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				handler(0, err)
				return
			}

			if resp.GetError() != "" {
				err = errors.New(resp.GetError())
			}
			// get first rune...
			for _, v := range resp.GetText() {
				handler(v, err)
				break
			}
		}
	}()

	return nil
}
