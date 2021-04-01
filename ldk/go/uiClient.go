package ldk

import (
	"context"
	"errors"
	"github.com/open-olive/loop-development-kit/ldk/go/v2/server"
	"github.com/open-olive/loop-development-kit/ldk/go/v2/service"
	"io"

	"github.com/open-olive/loop-development-kit/ldk/go/v2/proto"
)

// UIClient is used to hook into UI events such as listening to search controls.
type UIClient struct {
	UiClient proto.UIClient
	Session  *server.Session
}

// ListenSearchbar allows you to listen to searchbar events.
func (u *UIClient) ListenSearchbar(ctx context.Context, handler service.ListenSearchHandler) error {
	msg := &proto.SearchbarStreamRequest{Session: u.Session.ToProto()}
	client, err := u.UiClient.SearchbarStream(ctx, msg)
	if err != nil {
		return err
	}

	go func() {
		for {
			resp, err := client.Recv()
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
			handler(resp.Text, err)
		}
	}()

	return nil

}

// ListenGlobalSearch allows you to listen to global search (omnibar) events.
func (u *UIClient) ListenGlobalSearch(ctx context.Context, handler service.ListenSearchHandler) error {
	msg := &proto.GlobalSearchStreamRequest{Session: u.Session.ToProto()}
	client, err := u.UiClient.GlobalSearchStream(ctx, msg)
	if err != nil {
		return err
	}

	go func() {
		for {
			resp, err := client.Recv()
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
			handler(resp.Text, err)
		}
	}()
	return nil
}
