package client

import (
	"context"
	"errors"
	"github.com/open-olive/loop-development-kit/ldk/go/v2/server"
	"github.com/open-olive/loop-development-kit/ldk/go/v2/service"
	"io"

	"github.com/open-olive/loop-development-kit/ldk/go/v2/proto"
)

type CursorClient struct {
	CursorClient proto.CursorClient
	Session      *server.Session
}

func (c *CursorClient) Position(ctx context.Context) (service.CursorPosition, error) {
	resp, err := c.CursorClient.CursorPosition(ctx, &proto.CursorPositionRequest{
		Session: c.Session.ToProto(),
	})

	if err != nil {
		return service.CursorPosition{}, err
	}
	return service.CursorPosition{
		X: int(resp.X),
		Y: int(resp.Y),
	}, nil
}

func (c *CursorClient) ListenPosition(ctx context.Context, handler service.ListenPositionHandler) error {
	cursorReadStreamClient, err := c.CursorClient.CursorPositionStream(ctx, &proto.CursorPositionStreamRequest{
		Session: c.Session.ToProto(),
	})
	if err != nil {
		return err
	}

	for {
		resp, err := cursorReadStreamClient.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			handler(service.CursorPosition{}, err)
		}
		if resp.GetError() != "" {
			err = errors.New(resp.GetError())
		}
		handler(service.CursorPosition{
			X: int(resp.X),
			Y: int(resp.Y),
		}, err)
	}

	return nil
}
