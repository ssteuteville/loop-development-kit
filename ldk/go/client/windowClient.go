package client

import (
	"context"
	"errors"
	"github.com/open-olive/loop-development-kit/ldk/go/v2/server"
	"github.com/open-olive/loop-development-kit/ldk/go/v2/service"
	"github.com/open-olive/loop-development-kit/ldk/go/v2/utils"
	"io"

	"github.com/open-olive/loop-development-kit/ldk/go/v2/proto"
)

type WindowClient struct {
	WindowClient proto.WindowClient
	Session      *server.Session
}

func (w *WindowClient) ActiveWindow(ctx context.Context) (utils.WindowInfo, error) {
	resp, err := w.WindowClient.WindowActiveWindow(ctx, &proto.WindowActiveWindowRequest{
		Session: w.Session.ToProto(),
	})
	if err != nil {
		return utils.WindowInfo{}, err
	}

	return utils.WindowInfo{
		Title:  resp.Window.GetTitle(),
		Path:   resp.Window.GetPath(),
		PID:    int(resp.Window.GetPid()),
		X:      int(resp.Window.GetX()),
		Y:      int(resp.Window.GetY()),
		Width:  int(resp.Window.GetWidth()),
		Height: int(resp.Window.GetHeight()),
	}, nil
}

func (w *WindowClient) ListenActiveWindow(ctx context.Context, handler service.ListenActiveWindowHandler) error {
	stream, err := w.WindowClient.WindowActiveWindowStream(ctx, &proto.WindowActiveWindowStreamRequest{
		Session: w.Session.ToProto(),
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
				handler(utils.WindowInfo{}, err)
				return
			}

			if resp.GetError() != "" {
				err = errors.New(resp.GetError())
			}
			handler(utils.WindowInfo{
				Title:  resp.Window.GetTitle(),
				Path:   resp.Window.GetPath(),
				PID:    int(resp.Window.GetPid()),
				X:      int(resp.Window.GetX()),
				Y:      int(resp.Window.GetY()),
				Width:  int(resp.Window.GetWidth()),
				Height: int(resp.Window.GetHeight()),
			}, err)
		}
	}()

	return nil

}

func (w *WindowClient) State(ctx context.Context) ([]utils.WindowInfo, error) {
	resp, err := w.WindowClient.WindowState(ctx, &proto.WindowStateRequest{
		Session: w.Session.ToProto(),
	})

	if err != nil {
		return nil, err
	}

	infos := make([]utils.WindowInfo, len(resp.Window))
	for _, w := range resp.Window {
		infos = append(infos, utils.WindowInfo{
			Title:  w.GetTitle(),
			Path:   w.GetPath(),
			PID:    int(w.GetPid()),
			X:      int(w.GetX()),
			Y:      int(w.GetY()),
			Width:  int(w.GetWidth()),
			Height: int(w.GetHeight()),
		})
	}
	return infos, nil
}

func (w *WindowClient) ListenState(ctx context.Context, handler service.ListenWindowStateHandler) error {
	stream, err := w.WindowClient.WindowStateStream(ctx, &proto.WindowStateStreamRequest{
		Session: w.Session.ToProto(),
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
				handler(utils.WindowEvent{}, err)
				return
			}

			if resp.GetError() != "" {
				err = errors.New(resp.GetError())
			}
			we := utils.WindowEvent{
				Action: utils.ProtoWindowActionToAction(resp.GetAction()),
				Window: utils.WindowInfo{
					Title:  resp.Window.GetTitle(),
					Path:   resp.Window.GetPath(),
					PID:    int(resp.Window.GetPid()),
					X:      int(resp.Window.GetX()),
					Y:      int(resp.Window.GetY()),
					Width:  int(resp.Window.GetWidth()),
					Height: int(resp.Window.GetHeight()),
				},
			}

			handler(we, err)
		}
	}()

	return nil
}
