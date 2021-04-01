package client

import (
	"context"
	"errors"
	"github.com/open-olive/loop-development-kit/ldk/go/v2/proto"
	"github.com/open-olive/loop-development-kit/ldk/go/v2/server"
	"github.com/open-olive/loop-development-kit/ldk/go/v2/service"
	"io"
)

// ClipboardClient is used by loops to facilitate communication with the clipboard service
type ClipboardClient struct {
	ClipboardClient proto.ClipboardClient
	Session         *server.Session
}

// Read is used by the loop to get the current clipboard text
func (c *ClipboardClient) Read(ctx context.Context) (string, error) {
	resp, err := c.ClipboardClient.ClipboardRead(ctx, &proto.ClipboardReadRequest{
		Session: c.Session.ToProto(),
	})
	if err != nil {
		return "", err
	}

	return resp.Text, nil
}

// Listen is used by the loop to establish a stream for handling clipboard changes
func (c *ClipboardClient) Listen(ctx context.Context, clipboardListenConfigurations service.ClipboardListenConfiguration) error {
	stream, err := c.ClipboardClient.ClipboardReadStream(ctx, &proto.ClipboardReadStreamRequest{
		Session: c.Session.ToProto(),
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
				clipboardListenConfigurations.Handler("", err)
				return
			}

			if resp.GetError() != "" {
				err = errors.New(resp.GetError())
			}
			clipboardListenConfigurations.Handler(resp.GetText(), err)
		}
	}()

	return nil
}

// Write is used by the loop to write text to the clipboard
func (c *ClipboardClient) Write(ctx context.Context, text string) error {
	_, err := c.ClipboardClient.ClipboardWrite(ctx, &proto.ClipboardWriteRequest{
		Session: c.Session.ToProto(),
		Text:    text,
	})
	if err != nil {
		return err
	}

	return nil
}
